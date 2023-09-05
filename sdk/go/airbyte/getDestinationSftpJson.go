// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// DestinationSftpJSON DataSource
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
//			_, err := airbyte.LookupDestinationSftpJson(ctx, &airbyte.LookupDestinationSftpJsonArgs{
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
func LookupDestinationSftpJson(ctx *pulumi.Context, args *LookupDestinationSftpJsonArgs, opts ...pulumi.InvokeOption) (*LookupDestinationSftpJsonResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDestinationSftpJsonResult
	err := ctx.Invoke("airbyte:index/getDestinationSftpJson:getDestinationSftpJson", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDestinationSftpJson.
type LookupDestinationSftpJsonArgs struct {
	DestinationId string `pulumi:"destinationId"`
}

// A collection of values returned by getDestinationSftpJson.
type LookupDestinationSftpJsonResult struct {
	Configuration GetDestinationSftpJsonConfiguration `pulumi:"configuration"`
	DestinationId string                              `pulumi:"destinationId"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupDestinationSftpJsonOutput(ctx *pulumi.Context, args LookupDestinationSftpJsonOutputArgs, opts ...pulumi.InvokeOption) LookupDestinationSftpJsonResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDestinationSftpJsonResult, error) {
			args := v.(LookupDestinationSftpJsonArgs)
			r, err := LookupDestinationSftpJson(ctx, &args, opts...)
			var s LookupDestinationSftpJsonResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDestinationSftpJsonResultOutput)
}

// A collection of arguments for invoking getDestinationSftpJson.
type LookupDestinationSftpJsonOutputArgs struct {
	DestinationId pulumi.StringInput `pulumi:"destinationId"`
}

func (LookupDestinationSftpJsonOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDestinationSftpJsonArgs)(nil)).Elem()
}

// A collection of values returned by getDestinationSftpJson.
type LookupDestinationSftpJsonResultOutput struct{ *pulumi.OutputState }

func (LookupDestinationSftpJsonResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDestinationSftpJsonResult)(nil)).Elem()
}

func (o LookupDestinationSftpJsonResultOutput) ToLookupDestinationSftpJsonResultOutput() LookupDestinationSftpJsonResultOutput {
	return o
}

func (o LookupDestinationSftpJsonResultOutput) ToLookupDestinationSftpJsonResultOutputWithContext(ctx context.Context) LookupDestinationSftpJsonResultOutput {
	return o
}

func (o LookupDestinationSftpJsonResultOutput) Configuration() GetDestinationSftpJsonConfigurationOutput {
	return o.ApplyT(func(v LookupDestinationSftpJsonResult) GetDestinationSftpJsonConfiguration { return v.Configuration }).(GetDestinationSftpJsonConfigurationOutput)
}

func (o LookupDestinationSftpJsonResultOutput) DestinationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationSftpJsonResult) string { return v.DestinationId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDestinationSftpJsonResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationSftpJsonResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDestinationSftpJsonResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationSftpJsonResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDestinationSftpJsonResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationSftpJsonResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDestinationSftpJsonResultOutput{})
}
