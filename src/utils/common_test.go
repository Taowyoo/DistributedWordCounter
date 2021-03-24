package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type jsCfg struct {
	AccessKeyID        string
	SecretAccessKey    string
	Region             string
	DataBucketName     string
	ResultBucketName   string
	JobQueueName       string
	ResultQueueName    string
	SubJobQueueName    string
	SubResultQueueName string
}

var myCfg jsCfg

func Test_listEC2Instances(t *testing.T) {
	// Read cfg from jsonFile
	testPath := "../../config/config.json"
	data, err := ioutil.ReadFile(testPath)
	if err != nil {
		fmt.Printf("Failed to read config file '%s':%v", testPath, err)
	}
	err = json.Unmarshal(data, &myCfg)
	if err != nil {
		fmt.Printf("Failed to prase config file '%s':%v", testPath, err)
	}

	// Set cred env variables
	// fmt.Println("AWS_REGION:", myCfg.Region)
	os.Setenv("AWS_REGION", myCfg.Region)
	os.Setenv("AWS_ACCESS_KEY_ID", myCfg.AccessKeyID)
	os.Setenv("AWS_SECRET_ACCESS_KEY", myCfg.SecretAccessKey)
	os.Setenv("AWS_SESSION_TOKEN", "")
	// Load config from env variables
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	ec2client := ec2.NewFromConfig(cfg)
	type args struct {
		client *ec2.Client
	}
	tests := []struct {
		name string
		args args
		want []InstanceInfo
	}{
		{
			name: "TestList",
			args: args{
				client: ec2client,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListEC2Instances(tt.args.client); got == nil {
				t.Errorf("listEC2Instances() = %v, want not nil", got)
			}
		})
	}
}

func Test_submitJob(t *testing.T) {
	// Read cfg from jsonFile
	testPath := "../../config/config.json"
	data, err := ioutil.ReadFile(testPath)
	if err != nil {
		fmt.Printf("Failed to read config file '%s':%v", testPath, err)
	}
	err = json.Unmarshal(data, &myCfg)
	if err != nil {
		fmt.Printf("Failed to prase config file '%s':%v", testPath, err)
	}

	// Set cred env variables
	// fmt.Println("AWS_REGION:", myCfg.Region)
	os.Setenv("AWS_REGION", myCfg.Region)
	os.Setenv("AWS_ACCESS_KEY_ID", myCfg.AccessKeyID)
	os.Setenv("AWS_SECRET_ACCESS_KEY", myCfg.SecretAccessKey)
	os.Setenv("AWS_SESSION_TOKEN", "")
	// Load config from env variables
	cfg, err := config.LoadDefaultConfig(context.TODO())
	sqsclient := sqs.NewFromConfig(cfg)

	type args struct {
		client    *sqs.Client
		queueName string
		instance  InstanceInfo
		fileKey   string
		s3bucket  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestSubmitJobToSqs",
			args: args{
				client:    sqsclient,
				queueName: "jobs",
				instance: InstanceInfo{
					Name:      "workerTest",
					Id:        "testId0",
					PublicIP:  "1.2.3.4",
					PrivateIP: "192.168.0.1",
				},
				fileKey:  "fileKeyValue",
				s3bucket: "s3bucketName",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubmitJob(tt.args.client, tt.args.queueName, tt.args.instance, tt.args.fileKey, tt.args.s3bucket); got != tt.want {
				t.Errorf("submitJob() = %v, want %v", got, tt.want)
			}
		})
	}
}
