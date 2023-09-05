// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"internal"
)

// DestinationTypesense DataSource
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
// func main() {
// pulumi.Run(func(ctx *pulumi.Context) error {
// _, err := airbyte.LookupDestinationTypesense(ctx, %!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference), nil);
// if err != nil {
// return err
// }
// return nil
// })
// }
// ```
func LookupDestinationTypesense(ctx *pulumi.Context, args *LookupDestinationTypesenseArgs, opts ...pulumi.InvokeOption) (*LookupDestinationTypesenseResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDestinationTypesenseResult
	err := ctx.Invoke("airbyte:index/getDestinationTypesense:getDestinationTypesense", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDestinationTypesense.
type LookupDestinationTypesenseArgs struct {
	DestinationId string `pulumi:"destinationId"`
}

// A collection of values returned by getDestinationTypesense.
type LookupDestinationTypesenseResult struct {
	Configuration GetDestinationTypesenseConfiguration `pulumi:"configuration"`
	DestinationId string                               `pulumi:"destinationId"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupDestinationTypesenseOutput(ctx *pulumi.Context, args LookupDestinationTypesenseOutputArgs, opts ...pulumi.InvokeOption) LookupDestinationTypesenseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDestinationTypesenseResult, error) {
			args := v.(LookupDestinationTypesenseArgs)
			r, err := LookupDestinationTypesense(ctx, &args, opts...)
			var s LookupDestinationTypesenseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDestinationTypesenseResultOutput)
}

// A collection of arguments for invoking getDestinationTypesense.
type LookupDestinationTypesenseOutputArgs struct {
	DestinationId pulumi.StringInput `pulumi:"destinationId"`
}

func (LookupDestinationTypesenseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDestinationTypesenseArgs)(nil)).Elem()
}

// A collection of values returned by getDestinationTypesense.
type LookupDestinationTypesenseResultOutput struct{ *pulumi.OutputState }

func (LookupDestinationTypesenseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDestinationTypesenseResult)(nil)).Elem()
}

func (o LookupDestinationTypesenseResultOutput) ToLookupDestinationTypesenseResultOutput() LookupDestinationTypesenseResultOutput {
	return o
}

func (o LookupDestinationTypesenseResultOutput) ToLookupDestinationTypesenseResultOutputWithContext(ctx context.Context) LookupDestinationTypesenseResultOutput {
	return o
}

func (o LookupDestinationTypesenseResultOutput) Configuration() GetDestinationTypesenseConfigurationOutput {
	return o.ApplyT(func(v LookupDestinationTypesenseResult) GetDestinationTypesenseConfiguration { return v.Configuration }).(GetDestinationTypesenseConfigurationOutput)
}

func (o LookupDestinationTypesenseResultOutput) DestinationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationTypesenseResult) string { return v.DestinationId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDestinationTypesenseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationTypesenseResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDestinationTypesenseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationTypesenseResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDestinationTypesenseResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationTypesenseResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDestinationTypesenseResultOutput{})
}
