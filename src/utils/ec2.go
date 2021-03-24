package utils

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/smithy-go"
)

// EC2DescribeInstancesAPI defines the interface for the DescribeInstances function.
// We use this interface to test the function using a mocked service.
type EC2DescribeInstancesAPI interface {
	DescribeInstances(ctx context.Context,
		params *ec2.DescribeInstancesInput,
		optFns ...func(*ec2.Options)) (*ec2.DescribeInstancesOutput, error)
}

// GetInstances retrieves information about your Amazon Elastic Compute Cloud (Amazon EC2) instances.
// Inputs:
//     c is the context of the method call, which includes the AWS Region.
//     api is the interface that defines the method call.
//     input defines the input arguments to the service call.
// Output:
//     If success, a DescribeInstancesOutput object containing the result of the service call and nil.
//     Otherwise, nil and an error from the call to DescribeInstances.
func GetInstances(c context.Context, api EC2DescribeInstancesAPI, input *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
	return api.DescribeInstances(c, input)
}

// EC2MonitorInstancesAPI defines the interface for the MonitorInstances and UnmonitorInstances functions.
// We use this interface to test the function using a mocked service.
type EC2MonitorInstancesAPI interface {
	MonitorInstances(ctx context.Context,
		params *ec2.MonitorInstancesInput,
		optFns ...func(*ec2.Options)) (*ec2.MonitorInstancesOutput, error)

	UnmonitorInstances(ctx context.Context,
		params *ec2.UnmonitorInstancesInput,
		optFns ...func(*ec2.Options)) (*ec2.UnmonitorInstancesOutput, error)
}

// EnableMonitoring enables monitoring for an Amazon EC2 instance.
// Inputs:
//     c is the context of the method call, which includes the AWS Region.
//     api is the interface that defines the method call.
//     input defines the input arguments to the service call.
// Output:
//     If success, a MonitorInstancesOutput object containing the result of the service call and nil.
//     Otherwise, nil and an error from the call to MonitorInstances.
func EnableMonitoring(c context.Context, api EC2MonitorInstancesAPI, input *ec2.MonitorInstancesInput) (*ec2.MonitorInstancesOutput, error) {
	resp, err := api.MonitorInstances(c, input)

	// Do we have a DryRunOperation error?
	var apiErr smithy.APIError
	if errors.As(err, &apiErr) && apiErr.ErrorCode() == "DryRunOperation" {
		fmt.Println("User has permission to enable monitoring.")
		input.DryRun = false
		return api.MonitorInstances(c, input)
	}

	return resp, err
}

// DisableMonitoring disables monitoring for an Amazon EC2 instance.
// Inputs:
//     c is the context of the method call, which includes the AWS Region.
//     api is the interface that defines the method call.
//     input defines the input arguments to the service call.
// Output:
//     If success, a UnmonitorInstancesOutput object containing the result of the service call and nil.
//     Otherwise, nil and an error from the call to UnmonitorInstances.
func DisableMonitoring(c context.Context, api EC2MonitorInstancesAPI, input *ec2.UnmonitorInstancesInput) (*ec2.UnmonitorInstancesOutput, error) {
	resp, err := api.UnmonitorInstances(c, input)

	// Do we have a DryRunOperation error?
	var apiErr smithy.APIError
	if errors.As(err, &apiErr) && apiErr.ErrorCode() == "DryRunOperation" {
		fmt.Println("User has permission to disable monitoring.")
		input.DryRun = false
		return api.UnmonitorInstances(c, input)
	}

	return resp, err
}

// EC2CreateInstanceAPI defines the interface for the RunInstances and CreateTags functions.
// We use this interface to test the functions using a mocked service.
type EC2CreateInstanceAPI interface {
	RunInstances(ctx context.Context,
		params *ec2.RunInstancesInput,
		optFns ...func(*ec2.Options)) (*ec2.RunInstancesOutput, error)

	CreateTags(ctx context.Context,
		params *ec2.CreateTagsInput,
		optFns ...func(*ec2.Options)) (*ec2.CreateTagsOutput, error)
}

// MakeInstance creates an Amazon Elastic Compute Cloud (Amazon EC2) instance.
// Inputs:
//     c is the context of the method call, which includes the AWS Region.
//     api is the interface that defines the method call.
//     input defines the input arguments to the service call.
// Output:
//     If success, a RunInstancesOutput object containing the result of the service call and nil.
//     Otherwise, nil and an error from the call to RunInstances.
func MakeInstance(c context.Context, api EC2CreateInstanceAPI, input *ec2.RunInstancesInput) (*ec2.RunInstancesOutput, error) {
	return api.RunInstances(c, input)
}

// MakeTags creates tags for an Amazon Elastic Compute Cloud (Amazon EC2) instance.
// Inputs:
//     c is the context of the method call, which includes the AWS Region.
//     api is the interface that defines the method call.
//     input defines the input arguments to the service call.
// Output:
//     If success, a CreateTagsOutput object containing the result of the service call and nil.
//     Otherwise, nil and an error from the call to CreateTags.
func MakeTags(c context.Context, api EC2CreateInstanceAPI, input *ec2.CreateTagsInput) (*ec2.CreateTagsOutput, error) {
	return api.CreateTags(c, input)
}

// EC2RebootInstancesAPI defines the interface for the RebootInstances function.
// We use this interface to test the function using a mocked service.
type EC2RebootInstancesAPI interface {
	RebootInstances(ctx context.Context,
		params *ec2.RebootInstancesInput,
		optFns ...func(*ec2.Options)) (*ec2.RebootInstancesOutput, error)
}

// RebootInstance reboots an Amazon Elastic Compute Cloud (Amazon EC2) instance.
// Inputs:
//     c is the context of the method call, which includes the AWS Region.
//     api is the interface that defines the method call.
//     input defines the input arguments to the service call.
// Output:
//     If success, a RebootInstancesOutput object containing the result of the service call and nil.
//     Otherwise, nil and an error from the call to RebootInstances.
func RebootInstance(c context.Context, api EC2RebootInstancesAPI, input *ec2.RebootInstancesInput) (*ec2.RebootInstancesOutput, error) {
	resp, err := api.RebootInstances(c, input)

	var apiErr smithy.APIError
	if errors.As(err, &apiErr) && apiErr.ErrorCode() == "DryRunOperation" {
		fmt.Println("User has permission to enable monitoring.")
		input.DryRun = false
		return api.RebootInstances(c, input)
	}

	return resp, err
}

// EC2StartInstancesAPI defines the interface for the StartInstances function.
// We use this interface to test the function using a mocked service.
type EC2StartInstancesAPI interface {
	StartInstances(ctx context.Context,
		params *ec2.StartInstancesInput,
		optFns ...func(*ec2.Options)) (*ec2.StartInstancesOutput, error)
}

// StartInstance starts an Amazon Elastic Compute Cloud (Amazon EC2) instance.
// Inputs:
//     c is the context of the method call, which includes the AWS Region.
//     api is the interface that defines the method call.
//     input defines the input arguments to the service call.
// Output:
//     If success, a StartInstancesOutput object containing the result of the service call and nil.
//     Otherwise, nil and an error from the call to StartInstances.
func StartInstance(c context.Context, api EC2StartInstancesAPI, input *ec2.StartInstancesInput) (*ec2.StartInstancesOutput, error) {
	resp, err := api.StartInstances(c, input)

	var apiErr smithy.APIError
	if errors.As(err, &apiErr) && apiErr.ErrorCode() == "DryRunOperation" {
		fmt.Println("User has permission to start an instance.")
		input.DryRun = false
		return api.StartInstances(c, input)
	}

	return resp, err
}

// EC2StopInstancesAPI defines the interface for the StopInstances function.
// We use this interface to test the function using a mocked service.
type EC2StopInstancesAPI interface {
	StopInstances(ctx context.Context,
		params *ec2.StopInstancesInput,
		optFns ...func(*ec2.Options)) (*ec2.StopInstancesOutput, error)
}

// StopInstance stops an Amazon Elastic Compute Cloud (Amazon EC2) instance.
// Inputs:
//     c is the context of the method call, which includes the AWS Region.
//     api is the interface that defines the method call.
//     input defines the input arguments to the service call.
// Output:
//     If success, a StopInstancesOutput object containing the result of the service call and nil.
//     Otherwise, nil and an error from the call to StopInstances.
func StopInstance(c context.Context, api EC2StopInstancesAPI, input *ec2.StopInstancesInput) (*ec2.StopInstancesOutput, error) {
	resp, err := api.StopInstances(c, input)

	var apiErr smithy.APIError
	if errors.As(err, &apiErr) && apiErr.ErrorCode() == "DryRunOperation" {
		fmt.Println("User has permission to stop instances.")
		input.DryRun = false
		return api.StopInstances(c, input)
	}

	return resp, err
}
