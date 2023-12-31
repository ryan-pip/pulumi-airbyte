// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// DestinationFirebolt DataSource
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
//			_, err := airbyte.LookupDestinationFirebolt(ctx, &airbyte.LookupDestinationFireboltArgs{
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
func LookupDestinationFirebolt(ctx *pulumi.Context, args *LookupDestinationFireboltArgs, opts ...pulumi.InvokeOption) (*LookupDestinationFireboltResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDestinationFireboltResult
	err := ctx.Invoke("airbyte:index/getDestinationFirebolt:getDestinationFirebolt", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDestinationFirebolt.
type LookupDestinationFireboltArgs struct {
	DestinationId string `pulumi:"destinationId"`
}

// A collection of values returned by getDestinationFirebolt.
type LookupDestinationFireboltResult struct {
	Configuration GetDestinationFireboltConfiguration `pulumi:"configuration"`
	DestinationId string                              `pulumi:"destinationId"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupDestinationFireboltOutput(ctx *pulumi.Context, args LookupDestinationFireboltOutputArgs, opts ...pulumi.InvokeOption) LookupDestinationFireboltResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDestinationFireboltResult, error) {
			args := v.(LookupDestinationFireboltArgs)
			r, err := LookupDestinationFirebolt(ctx, &args, opts...)
			var s LookupDestinationFireboltResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDestinationFireboltResultOutput)
}

// A collection of arguments for invoking getDestinationFirebolt.
type LookupDestinationFireboltOutputArgs struct {
	DestinationId pulumi.StringInput `pulumi:"destinationId"`
}

func (LookupDestinationFireboltOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDestinationFireboltArgs)(nil)).Elem()
}

// A collection of values returned by getDestinationFirebolt.
type LookupDestinationFireboltResultOutput struct{ *pulumi.OutputState }

func (LookupDestinationFireboltResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDestinationFireboltResult)(nil)).Elem()
}

func (o LookupDestinationFireboltResultOutput) ToLookupDestinationFireboltResultOutput() LookupDestinationFireboltResultOutput {
	return o
}

func (o LookupDestinationFireboltResultOutput) ToLookupDestinationFireboltResultOutputWithContext(ctx context.Context) LookupDestinationFireboltResultOutput {
	return o
}

func (o LookupDestinationFireboltResultOutput) Configuration() GetDestinationFireboltConfigurationOutput {
	return o.ApplyT(func(v LookupDestinationFireboltResult) GetDestinationFireboltConfiguration { return v.Configuration }).(GetDestinationFireboltConfigurationOutput)
}

func (o LookupDestinationFireboltResultOutput) DestinationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationFireboltResult) string { return v.DestinationId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDestinationFireboltResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationFireboltResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDestinationFireboltResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationFireboltResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDestinationFireboltResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationFireboltResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDestinationFireboltResultOutput{})
}
