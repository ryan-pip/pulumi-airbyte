// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// DestinationS3 DataSource
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
//			_, err := airbyte.LookupDestinationS3(ctx, &airbyte.LookupDestinationS3Args{
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
func LookupDestinationS3(ctx *pulumi.Context, args *LookupDestinationS3Args, opts ...pulumi.InvokeOption) (*LookupDestinationS3Result, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDestinationS3Result
	err := ctx.Invoke("airbyte:index/getDestinationS3:getDestinationS3", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDestinationS3.
type LookupDestinationS3Args struct {
	DestinationId string `pulumi:"destinationId"`
}

// A collection of values returned by getDestinationS3.
type LookupDestinationS3Result struct {
	Configuration GetDestinationS3Configuration `pulumi:"configuration"`
	DestinationId string                        `pulumi:"destinationId"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupDestinationS3Output(ctx *pulumi.Context, args LookupDestinationS3OutputArgs, opts ...pulumi.InvokeOption) LookupDestinationS3ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDestinationS3Result, error) {
			args := v.(LookupDestinationS3Args)
			r, err := LookupDestinationS3(ctx, &args, opts...)
			var s LookupDestinationS3Result
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDestinationS3ResultOutput)
}

// A collection of arguments for invoking getDestinationS3.
type LookupDestinationS3OutputArgs struct {
	DestinationId pulumi.StringInput `pulumi:"destinationId"`
}

func (LookupDestinationS3OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDestinationS3Args)(nil)).Elem()
}

// A collection of values returned by getDestinationS3.
type LookupDestinationS3ResultOutput struct{ *pulumi.OutputState }

func (LookupDestinationS3ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDestinationS3Result)(nil)).Elem()
}

func (o LookupDestinationS3ResultOutput) ToLookupDestinationS3ResultOutput() LookupDestinationS3ResultOutput {
	return o
}

func (o LookupDestinationS3ResultOutput) ToLookupDestinationS3ResultOutputWithContext(ctx context.Context) LookupDestinationS3ResultOutput {
	return o
}

func (o LookupDestinationS3ResultOutput) Configuration() GetDestinationS3ConfigurationOutput {
	return o.ApplyT(func(v LookupDestinationS3Result) GetDestinationS3Configuration { return v.Configuration }).(GetDestinationS3ConfigurationOutput)
}

func (o LookupDestinationS3ResultOutput) DestinationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationS3Result) string { return v.DestinationId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDestinationS3ResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationS3Result) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDestinationS3ResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationS3Result) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDestinationS3ResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationS3Result) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDestinationS3ResultOutput{})
}
