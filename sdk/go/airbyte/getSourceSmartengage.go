// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"internal"
)

// SourceSmartengage DataSource
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
// _, err := airbyte.LookupSourceSmartengage(ctx, %!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference), nil);
// if err != nil {
// return err
// }
// return nil
// })
// }
// ```
func LookupSourceSmartengage(ctx *pulumi.Context, args *LookupSourceSmartengageArgs, opts ...pulumi.InvokeOption) (*LookupSourceSmartengageResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceSmartengageResult
	err := ctx.Invoke("airbyte:index/getSourceSmartengage:getSourceSmartengage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceSmartengage.
type LookupSourceSmartengageArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceSmartengage.
type LookupSourceSmartengageResult struct {
	Configuration GetSourceSmartengageConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceSmartengageOutput(ctx *pulumi.Context, args LookupSourceSmartengageOutputArgs, opts ...pulumi.InvokeOption) LookupSourceSmartengageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceSmartengageResult, error) {
			args := v.(LookupSourceSmartengageArgs)
			r, err := LookupSourceSmartengage(ctx, &args, opts...)
			var s LookupSourceSmartengageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceSmartengageResultOutput)
}

// A collection of arguments for invoking getSourceSmartengage.
type LookupSourceSmartengageOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceSmartengageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceSmartengageArgs)(nil)).Elem()
}

// A collection of values returned by getSourceSmartengage.
type LookupSourceSmartengageResultOutput struct{ *pulumi.OutputState }

func (LookupSourceSmartengageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceSmartengageResult)(nil)).Elem()
}

func (o LookupSourceSmartengageResultOutput) ToLookupSourceSmartengageResultOutput() LookupSourceSmartengageResultOutput {
	return o
}

func (o LookupSourceSmartengageResultOutput) ToLookupSourceSmartengageResultOutputWithContext(ctx context.Context) LookupSourceSmartengageResultOutput {
	return o
}

func (o LookupSourceSmartengageResultOutput) Configuration() GetSourceSmartengageConfigurationOutput {
	return o.ApplyT(func(v LookupSourceSmartengageResult) GetSourceSmartengageConfiguration { return v.Configuration }).(GetSourceSmartengageConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceSmartengageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSmartengageResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceSmartengageResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSmartengageResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceSmartengageResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceSmartengageResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceSmartengageResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSmartengageResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceSmartengageResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSmartengageResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceSmartengageResultOutput{})
}
