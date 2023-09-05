// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"internal"
)

// SourceBambooHr DataSource
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
// _, err := airbyte.LookupSourceBambooHr(ctx, %!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference), nil);
// if err != nil {
// return err
// }
// return nil
// })
// }
// ```
func LookupSourceBambooHr(ctx *pulumi.Context, args *LookupSourceBambooHrArgs, opts ...pulumi.InvokeOption) (*LookupSourceBambooHrResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceBambooHrResult
	err := ctx.Invoke("airbyte:index/getSourceBambooHr:getSourceBambooHr", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceBambooHr.
type LookupSourceBambooHrArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceBambooHr.
type LookupSourceBambooHrResult struct {
	Configuration GetSourceBambooHrConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceBambooHrOutput(ctx *pulumi.Context, args LookupSourceBambooHrOutputArgs, opts ...pulumi.InvokeOption) LookupSourceBambooHrResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceBambooHrResult, error) {
			args := v.(LookupSourceBambooHrArgs)
			r, err := LookupSourceBambooHr(ctx, &args, opts...)
			var s LookupSourceBambooHrResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceBambooHrResultOutput)
}

// A collection of arguments for invoking getSourceBambooHr.
type LookupSourceBambooHrOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceBambooHrOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceBambooHrArgs)(nil)).Elem()
}

// A collection of values returned by getSourceBambooHr.
type LookupSourceBambooHrResultOutput struct{ *pulumi.OutputState }

func (LookupSourceBambooHrResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceBambooHrResult)(nil)).Elem()
}

func (o LookupSourceBambooHrResultOutput) ToLookupSourceBambooHrResultOutput() LookupSourceBambooHrResultOutput {
	return o
}

func (o LookupSourceBambooHrResultOutput) ToLookupSourceBambooHrResultOutputWithContext(ctx context.Context) LookupSourceBambooHrResultOutput {
	return o
}

func (o LookupSourceBambooHrResultOutput) Configuration() GetSourceBambooHrConfigurationOutput {
	return o.ApplyT(func(v LookupSourceBambooHrResult) GetSourceBambooHrConfiguration { return v.Configuration }).(GetSourceBambooHrConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceBambooHrResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceBambooHrResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceBambooHrResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceBambooHrResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceBambooHrResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceBambooHrResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceBambooHrResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceBambooHrResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceBambooHrResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceBambooHrResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceBambooHrResultOutput{})
}
