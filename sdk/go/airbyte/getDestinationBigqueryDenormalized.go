// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// DestinationBigqueryDenormalized DataSource
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
//			_, err := airbyte.LookupDestinationBigqueryDenormalized(ctx, &airbyte.LookupDestinationBigqueryDenormalizedArgs{
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
func LookupDestinationBigqueryDenormalized(ctx *pulumi.Context, args *LookupDestinationBigqueryDenormalizedArgs, opts ...pulumi.InvokeOption) (*LookupDestinationBigqueryDenormalizedResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDestinationBigqueryDenormalizedResult
	err := ctx.Invoke("airbyte:index/getDestinationBigqueryDenormalized:getDestinationBigqueryDenormalized", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDestinationBigqueryDenormalized.
type LookupDestinationBigqueryDenormalizedArgs struct {
	DestinationId string `pulumi:"destinationId"`
}

// A collection of values returned by getDestinationBigqueryDenormalized.
type LookupDestinationBigqueryDenormalizedResult struct {
	Configuration GetDestinationBigqueryDenormalizedConfiguration `pulumi:"configuration"`
	DestinationId string                                          `pulumi:"destinationId"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupDestinationBigqueryDenormalizedOutput(ctx *pulumi.Context, args LookupDestinationBigqueryDenormalizedOutputArgs, opts ...pulumi.InvokeOption) LookupDestinationBigqueryDenormalizedResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDestinationBigqueryDenormalizedResult, error) {
			args := v.(LookupDestinationBigqueryDenormalizedArgs)
			r, err := LookupDestinationBigqueryDenormalized(ctx, &args, opts...)
			var s LookupDestinationBigqueryDenormalizedResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDestinationBigqueryDenormalizedResultOutput)
}

// A collection of arguments for invoking getDestinationBigqueryDenormalized.
type LookupDestinationBigqueryDenormalizedOutputArgs struct {
	DestinationId pulumi.StringInput `pulumi:"destinationId"`
}

func (LookupDestinationBigqueryDenormalizedOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDestinationBigqueryDenormalizedArgs)(nil)).Elem()
}

// A collection of values returned by getDestinationBigqueryDenormalized.
type LookupDestinationBigqueryDenormalizedResultOutput struct{ *pulumi.OutputState }

func (LookupDestinationBigqueryDenormalizedResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDestinationBigqueryDenormalizedResult)(nil)).Elem()
}

func (o LookupDestinationBigqueryDenormalizedResultOutput) ToLookupDestinationBigqueryDenormalizedResultOutput() LookupDestinationBigqueryDenormalizedResultOutput {
	return o
}

func (o LookupDestinationBigqueryDenormalizedResultOutput) ToLookupDestinationBigqueryDenormalizedResultOutputWithContext(ctx context.Context) LookupDestinationBigqueryDenormalizedResultOutput {
	return o
}

func (o LookupDestinationBigqueryDenormalizedResultOutput) Configuration() GetDestinationBigqueryDenormalizedConfigurationOutput {
	return o.ApplyT(func(v LookupDestinationBigqueryDenormalizedResult) GetDestinationBigqueryDenormalizedConfiguration {
		return v.Configuration
	}).(GetDestinationBigqueryDenormalizedConfigurationOutput)
}

func (o LookupDestinationBigqueryDenormalizedResultOutput) DestinationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationBigqueryDenormalizedResult) string { return v.DestinationId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDestinationBigqueryDenormalizedResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationBigqueryDenormalizedResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDestinationBigqueryDenormalizedResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationBigqueryDenormalizedResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDestinationBigqueryDenormalizedResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationBigqueryDenormalizedResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDestinationBigqueryDenormalizedResultOutput{})
}
