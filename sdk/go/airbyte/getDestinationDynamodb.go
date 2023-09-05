// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// DestinationDynamodb DataSource
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := airbyte.LookupDestinationDynamodb(ctx, &airbyte.LookupDestinationDynamodbArgs{
//				DestinationId: "...my_destination_id...",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupDestinationDynamodb(ctx *pulumi.Context, args *LookupDestinationDynamodbArgs, opts ...pulumi.InvokeOption) (*LookupDestinationDynamodbResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDestinationDynamodbResult
	err := ctx.Invoke("airbyte:index/getDestinationDynamodb:getDestinationDynamodb", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDestinationDynamodb.
type LookupDestinationDynamodbArgs struct {
	DestinationId string `pulumi:"destinationId"`
}

// A collection of values returned by getDestinationDynamodb.
type LookupDestinationDynamodbResult struct {
	Configuration GetDestinationDynamodbConfiguration `pulumi:"configuration"`
	DestinationId string                              `pulumi:"destinationId"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupDestinationDynamodbOutput(ctx *pulumi.Context, args LookupDestinationDynamodbOutputArgs, opts ...pulumi.InvokeOption) LookupDestinationDynamodbResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDestinationDynamodbResult, error) {
			args := v.(LookupDestinationDynamodbArgs)
			r, err := LookupDestinationDynamodb(ctx, &args, opts...)
			var s LookupDestinationDynamodbResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDestinationDynamodbResultOutput)
}

// A collection of arguments for invoking getDestinationDynamodb.
type LookupDestinationDynamodbOutputArgs struct {
	DestinationId pulumi.StringInput `pulumi:"destinationId"`
}

func (LookupDestinationDynamodbOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDestinationDynamodbArgs)(nil)).Elem()
}

// A collection of values returned by getDestinationDynamodb.
type LookupDestinationDynamodbResultOutput struct{ *pulumi.OutputState }

func (LookupDestinationDynamodbResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDestinationDynamodbResult)(nil)).Elem()
}

func (o LookupDestinationDynamodbResultOutput) ToLookupDestinationDynamodbResultOutput() LookupDestinationDynamodbResultOutput {
	return o
}

func (o LookupDestinationDynamodbResultOutput) ToLookupDestinationDynamodbResultOutputWithContext(ctx context.Context) LookupDestinationDynamodbResultOutput {
	return o
}

func (o LookupDestinationDynamodbResultOutput) Configuration() GetDestinationDynamodbConfigurationOutput {
	return o.ApplyT(func(v LookupDestinationDynamodbResult) GetDestinationDynamodbConfiguration { return v.Configuration }).(GetDestinationDynamodbConfigurationOutput)
}

func (o LookupDestinationDynamodbResultOutput) DestinationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationDynamodbResult) string { return v.DestinationId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDestinationDynamodbResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationDynamodbResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDestinationDynamodbResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationDynamodbResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDestinationDynamodbResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationDynamodbResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDestinationDynamodbResultOutput{})
}
