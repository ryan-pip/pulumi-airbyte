// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-airbyte/sdk/go/airbyte/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// DestinationVertica DataSource
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-airbyte/sdk/go/airbyte"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := airbyte.LookupDestinationVertica(ctx, &airbyte.LookupDestinationVerticaArgs{
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
func LookupDestinationVertica(ctx *pulumi.Context, args *LookupDestinationVerticaArgs, opts ...pulumi.InvokeOption) (*LookupDestinationVerticaResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDestinationVerticaResult
	err := ctx.Invoke("airbyte:index/getDestinationVertica:getDestinationVertica", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDestinationVertica.
type LookupDestinationVerticaArgs struct {
	DestinationId string `pulumi:"destinationId"`
}

// A collection of values returned by getDestinationVertica.
type LookupDestinationVerticaResult struct {
	Configuration GetDestinationVerticaConfiguration `pulumi:"configuration"`
	DestinationId string                             `pulumi:"destinationId"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupDestinationVerticaOutput(ctx *pulumi.Context, args LookupDestinationVerticaOutputArgs, opts ...pulumi.InvokeOption) LookupDestinationVerticaResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDestinationVerticaResult, error) {
			args := v.(LookupDestinationVerticaArgs)
			r, err := LookupDestinationVertica(ctx, &args, opts...)
			var s LookupDestinationVerticaResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDestinationVerticaResultOutput)
}

// A collection of arguments for invoking getDestinationVertica.
type LookupDestinationVerticaOutputArgs struct {
	DestinationId pulumi.StringInput `pulumi:"destinationId"`
}

func (LookupDestinationVerticaOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDestinationVerticaArgs)(nil)).Elem()
}

// A collection of values returned by getDestinationVertica.
type LookupDestinationVerticaResultOutput struct{ *pulumi.OutputState }

func (LookupDestinationVerticaResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDestinationVerticaResult)(nil)).Elem()
}

func (o LookupDestinationVerticaResultOutput) ToLookupDestinationVerticaResultOutput() LookupDestinationVerticaResultOutput {
	return o
}

func (o LookupDestinationVerticaResultOutput) ToLookupDestinationVerticaResultOutputWithContext(ctx context.Context) LookupDestinationVerticaResultOutput {
	return o
}

func (o LookupDestinationVerticaResultOutput) Configuration() GetDestinationVerticaConfigurationOutput {
	return o.ApplyT(func(v LookupDestinationVerticaResult) GetDestinationVerticaConfiguration { return v.Configuration }).(GetDestinationVerticaConfigurationOutput)
}

func (o LookupDestinationVerticaResultOutput) DestinationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationVerticaResult) string { return v.DestinationId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDestinationVerticaResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationVerticaResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDestinationVerticaResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationVerticaResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDestinationVerticaResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationVerticaResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDestinationVerticaResultOutput{})
}
