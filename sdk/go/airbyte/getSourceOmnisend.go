// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"internal"
)

// SourceOmnisend DataSource
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
// _, err := airbyte.LookupSourceOmnisend(ctx, %!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference), nil);
// if err != nil {
// return err
// }
// return nil
// })
// }
// ```
func LookupSourceOmnisend(ctx *pulumi.Context, args *LookupSourceOmnisendArgs, opts ...pulumi.InvokeOption) (*LookupSourceOmnisendResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceOmnisendResult
	err := ctx.Invoke("airbyte:index/getSourceOmnisend:getSourceOmnisend", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceOmnisend.
type LookupSourceOmnisendArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceOmnisend.
type LookupSourceOmnisendResult struct {
	Configuration GetSourceOmnisendConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceOmnisendOutput(ctx *pulumi.Context, args LookupSourceOmnisendOutputArgs, opts ...pulumi.InvokeOption) LookupSourceOmnisendResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceOmnisendResult, error) {
			args := v.(LookupSourceOmnisendArgs)
			r, err := LookupSourceOmnisend(ctx, &args, opts...)
			var s LookupSourceOmnisendResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceOmnisendResultOutput)
}

// A collection of arguments for invoking getSourceOmnisend.
type LookupSourceOmnisendOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceOmnisendOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceOmnisendArgs)(nil)).Elem()
}

// A collection of values returned by getSourceOmnisend.
type LookupSourceOmnisendResultOutput struct{ *pulumi.OutputState }

func (LookupSourceOmnisendResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceOmnisendResult)(nil)).Elem()
}

func (o LookupSourceOmnisendResultOutput) ToLookupSourceOmnisendResultOutput() LookupSourceOmnisendResultOutput {
	return o
}

func (o LookupSourceOmnisendResultOutput) ToLookupSourceOmnisendResultOutputWithContext(ctx context.Context) LookupSourceOmnisendResultOutput {
	return o
}

func (o LookupSourceOmnisendResultOutput) Configuration() GetSourceOmnisendConfigurationOutput {
	return o.ApplyT(func(v LookupSourceOmnisendResult) GetSourceOmnisendConfiguration { return v.Configuration }).(GetSourceOmnisendConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceOmnisendResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceOmnisendResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceOmnisendResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceOmnisendResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceOmnisendResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceOmnisendResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceOmnisendResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceOmnisendResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceOmnisendResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceOmnisendResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceOmnisendResultOutput{})
}
