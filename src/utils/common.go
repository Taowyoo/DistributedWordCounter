package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

type InstanceInfo struct {
	Name      string
	Id        string
	PublicIP  string
	PrivateIP string
}

func SubmitJob(client *sqs.Client, queueName string, instance InstanceInfo, fileKey string, s3bucket string) bool {
	// Get URL of queue
	queue := &queueName
	gQInput := &sqs.GetQueueUrlInput{
		QueueName: queue,
	}

	result, err := GetQueueURL(context.TODO(), client, gQInput)
	if err != nil {
		fmt.Println("Got an error getting the queue URL:")
		fmt.Println(err)
		return false
	}

	queueURL := result.QueueUrl

	sMInput := &sqs.SendMessageInput{
		DelaySeconds: 10,
		MessageAttributes: map[string]types.MessageAttributeValue{
			"JobId": {
				DataType:    aws.String("Number"),
				StringValue: aws.String(fmt.Sprint(time.Now().UnixNano() / int64(time.Millisecond))),
			},
			"WorkerId": {
				DataType:    aws.String("String"),
				StringValue: aws.String(instance.Id),
			},
			"WorkerPubIP": {
				DataType:    aws.String("String"),
				StringValue: aws.String(instance.PublicIP),
			},
			"WorkerPriIP": {
				DataType:    aws.String("String"),
				StringValue: aws.String(instance.PrivateIP),
			},
		},
		MessageBody: aws.String(fileKey + " " + s3bucket),
		QueueUrl:    queueURL,
	}

	resp, err := SendMsg(context.TODO(), client, sMInput)
	if err != nil {
		fmt.Println("Got an error sending the message:")
		fmt.Println(err)
		return false
	}

	fmt.Printf("Sent job msg with ID '%s' for instance '%s':'%s' to queue:'%s'\n", *resp.MessageId, instance.Id, instance.PublicIP, *queueURL)
	return true
}

func ListEC2Instances(client *ec2.Client) []InstanceInfo {
	input := &ec2.DescribeInstancesInput{}

	result, err := GetInstances(context.TODO(), client, input)
	if err != nil {
		fmt.Println("Got an error retrieving information about your Amazon EC2 instances:")
		fmt.Println(err)
		return nil
	}
	ret := make([]InstanceInfo, 0, 3)
	for _, r := range result.Reservations {
		// fmt.Println("Reservation ID: " + *r.ReservationId)
		// fmt.Println("Running Instance IDs:")
		for _, i := range r.Instances {
			if i.State.Code == 16 {
				// fmt.Println("   ", *i.InstanceId, *i.PublicIpAddress, *i.PrivateIpAddress)
				tags := i.Tags
				var name string
				for _, val := range tags {
					if *val.Key == "Name" {
						name = *val.Value
					}
				}
				ret = append(ret, InstanceInfo{
					Name:      name,
					Id:        *i.InstanceId,
					PublicIP:  *i.PublicIpAddress,
					PrivateIP: *i.PrivateIpAddress,
				})
			}

		}
		// fmt.Println("")
	}
	fmt.Println("Running Instance:")
	fmt.Printf("   %10s %20s %15s %15s\n", "Name", "Id", "PublicIP", "PrivateIP")
	for _, val := range ret {
		fmt.Printf("   %10s %20s %15s %15s\n", val.Name, val.Id, val.PublicIP, val.PrivateIP)
	}
	return ret
}

func GetLPMessagesByURL(client *sqs.Client, queueURL string, msgNum int, waitTime int) (*sqs.ReceiveMessageOutput, error) {
	mInput := &sqs.ReceiveMessageInput{
		QueueUrl: &queueURL,
		AttributeNames: []types.QueueAttributeName{
			"SentTimestamp",
		},
		MaxNumberOfMessages: 1,
		MessageAttributeNames: []string{
			"All",
		},
		WaitTimeSeconds: int32(waitTime),
	}

	return GetLPMessages(context.TODO(), client, mInput)
}

func GetQueueURLSimple(client *sqs.Client, queueName string) string {
	qInput := &sqs.GetQueueUrlInput{
		QueueName: &queueName,
	}

	result, err := GetQueueURL(context.TODO(), client, qInput)
	if err != nil {
		fmt.Println("Got an error getting the queue URL:")
		fmt.Println(err)
		return ""
	}

	return *result.QueueUrl
}

func RemoveMessageSimple(client *sqs.Client, queueName string, handle string) {
	dMInput := &sqs.DeleteMessageInput{
		QueueUrl:      &queueName,
		ReceiptHandle: &handle,
	}

	_, err := RemoveMessage(context.TODO(), client, dMInput)
	if err != nil {
		fmt.Println("Got an error deleting the message:")
		fmt.Println(err)
		return
	}
	fmt.Println("Deleted message from queue with URL " + queueName)
}

func DeleteObjectSimple(client *s3.Client, objectKey string, bucket string) {
	input := &s3.DeleteObjectInput{
		Bucket: &bucket,
		Key:    &objectKey,
	}

	_, err := DeleteItem(context.TODO(), client, input)
	if err != nil {
		fmt.Println("Got an error deleting item:")
		fmt.Println(err)
		return
	}

	fmt.Println("Deleted " + objectKey + " from " + bucket)
}
