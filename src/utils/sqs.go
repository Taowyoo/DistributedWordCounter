package utils

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

// SQSDeleteMessageAPI defines the interface for the GetQueueUrl and DeleteMessage functions.
// We use this interface to test the functions using a mocked service.
type SQSAPI interface {
	GetQueueUrl(ctx context.Context,
		params *sqs.GetQueueUrlInput,
		optFns ...func(*sqs.Options)) (*sqs.GetQueueUrlOutput, error)

	DeleteMessage(ctx context.Context,
		params *sqs.DeleteMessageInput,
		optFns ...func(*sqs.Options)) (*sqs.DeleteMessageOutput, error)
	ReceiveMessage(ctx context.Context,
		params *sqs.ReceiveMessageInput,
		optFns ...func(*sqs.Options)) (*sqs.ReceiveMessageOutput, error)

	SendMessage(ctx context.Context,
		params *sqs.SendMessageInput,
		optFns ...func(*sqs.Options)) (*sqs.SendMessageOutput, error)
}

// GetQueueURL gets the URL of an Amazon SQS queue.
// Inputs:
//     c is the context of the method call, which includes the AWS Region.
//     api is the interface that defines the method call.
//     input defines the input arguments to the service call.
// Output:
//     If success, a GetQueueUrlOutput object containing the result of the service call and nil.
//     Otherwise, nil and an error from the call to GetQueueUrl.
func GetQueueURL(c context.Context, api SQSAPI, input *sqs.GetQueueUrlInput) (*sqs.GetQueueUrlOutput, error) {
	return api.GetQueueUrl(c, input)
}

// RemoveMessage deletes a message from an Amazon SQS queue.
// Inputs:
//     c is the context of the method call, which includes the AWS Region.
//     api is the interface that defines the method call.
//     input defines the input arguments to the service call.
// Output:
//     If success, a DeleteMessageOutput object containing the result of the service call and nil.
//     Otherwise, nil and an error from the call to DeleteMessage.
func RemoveMessage(c context.Context, api SQSAPI, input *sqs.DeleteMessageInput) (*sqs.DeleteMessageOutput, error) {
	return api.DeleteMessage(c, input)
}

// GetMessages gets the most recent message from an Amazon SQS queue.
// Inputs:
//     c is the context of the method call, which includes the AWS Region.
//     api is the interface that defines the method call.
//     input defines the input arguments to the service call.
// Output:
//     If success, a ReceiveMessageOutput object containing the result of the service call and nil.
//     Otherwise, nil and an error from the call to ReceiveMessage.
func GetMessages(c context.Context, api SQSAPI, input *sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error) {
	return api.ReceiveMessage(c, input)
}

// GetLPMessages gets the messages from an Amazon SQS long polling queue.
// Inputs:
//     c is the context of the method call, which includes the AWS Region.
//     api is the interface that defines the method call.
//     input defines the input arguments to the service call.
// Output:
//     If success, a ReceiveMessageOutput object containing the result of the service call and nil.
//     Otherwise, nil and an error from the call to ReceiveMessage.
func GetLPMessages(c context.Context, api SQSAPI, input *sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error) {
	return api.ReceiveMessage(c, input)
}

// SendMsg sends a message to an Amazon SQS queue.
// Inputs:
//     c is the context of the method call, which includes the AWS Region.
//     api is the interface that defines the method call.
//     input defines the input arguments to the service call.
// Output:
//     If success, a SendMessageOutput object containing the result of the service call and nil.
//     Otherwise, nil and an error from the call to SendMessage.
func SendMsg(c context.Context, api SQSAPI, input *sqs.SendMessageInput) (*sqs.SendMessageOutput, error) {
	return api.SendMessage(c, input)
}

// SQSListQueuesAPI defines the interface for the ListQueues function.
// We use this interface to test the function using a mocked service.
type SQSListQueuesAPI interface {
	ListQueues(ctx context.Context,
		params *sqs.ListQueuesInput,
		optFns ...func(*sqs.Options)) (*sqs.ListQueuesOutput, error)
}

// GetQueues retrieves a list of your Amazon Simple Queue Service (Amazon SQS) queues.
// Inputs:
//     c is the context of the method call, which includes the AWS Region.
//     api is the interface that defines the method call.
//     input defines the input arguments to the service call.
// Output:
//     If success, a ListQueuesOutput object containing the result of the service call and nil.
//     Otherwise, nil and an error from the call to ListQueues.
func GetQueues(c context.Context, api SQSListQueuesAPI, input *sqs.ListQueuesInput) (*sqs.ListQueuesOutput, error) {
	return api.ListQueues(c, input)
}
