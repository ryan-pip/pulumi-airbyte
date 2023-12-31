// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// DestinationClickhouse DataSource
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
//			_, err := airbyte.LookupDestinationClickhouse(ctx, &airbyte.LookupDestinationClickhouseArgs{
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
func LookupDestinationClickhouse(ctx *pulumi.Context, args *LookupDestinationClickhouseArgs, opts ...pulumi.InvokeOption) (*LookupDestinationClickhouseResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDestinationClickhouseResult
	err := ctx.Invoke("airbyte:index/getDestinationClickhouse:getDestinationClickhouse", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDestinationClickhouse.
type LookupDestinationClickhouseArgs struct {
	DestinationId string `pulumi:"destinationId"`
}

// A collection of values returned by getDestinationClickhouse.
type LookupDestinationClickhouseResult struct {
	Configuration GetDestinationClickhouseConfiguration `pulumi:"configuration"`
	DestinationId string                                `pulumi:"destinationId"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupDestinationClickhouseOutput(ctx *pulumi.Context, args LookupDestinationClickhouseOutputArgs, opts ...pulumi.InvokeOption) LookupDestinationClickhouseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDestinationClickhouseResult, error) {
			args := v.(LookupDestinationClickhouseArgs)
			r, err := LookupDestinationClickhouse(ctx, &args, opts...)
			var s LookupDestinationClickhouseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDestinationClickhouseResultOutput)
}

// A collection of arguments for invoking getDestinationClickhouse.
type LookupDestinationClickhouseOutputArgs struct {
	DestinationId pulumi.StringInput `pulumi:"destinationId"`
}

func (LookupDestinationClickhouseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDestinationClickhouseArgs)(nil)).Elem()
}

// A collection of values returned by getDestinationClickhouse.
type LookupDestinationClickhouseResultOutput struct{ *pulumi.OutputState }

func (LookupDestinationClickhouseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDestinationClickhouseResult)(nil)).Elem()
}

func (o LookupDestinationClickhouseResultOutput) ToLookupDestinationClickhouseResultOutput() LookupDestinationClickhouseResultOutput {
	return o
}

func (o LookupDestinationClickhouseResultOutput) ToLookupDestinationClickhouseResultOutputWithContext(ctx context.Context) LookupDestinationClickhouseResultOutput {
	return o
}

func (o LookupDestinationClickhouseResultOutput) Configuration() GetDestinationClickhouseConfigurationOutput {
	return o.ApplyT(func(v LookupDestinationClickhouseResult) GetDestinationClickhouseConfiguration {
		return v.Configuration
	}).(GetDestinationClickhouseConfigurationOutput)
}

func (o LookupDestinationClickhouseResultOutput) DestinationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationClickhouseResult) string { return v.DestinationId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDestinationClickhouseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationClickhouseResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDestinationClickhouseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationClickhouseResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDestinationClickhouseResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationClickhouseResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDestinationClickhouseResultOutput{})
}
