// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"internal"
)

// SourcePaystack DataSource
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
// _, err := airbyte.LookupSourcePaystack(ctx, %!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference), nil);
// if err != nil {
// return err
// }
// return nil
// })
// }
// ```
func LookupSourcePaystack(ctx *pulumi.Context, args *LookupSourcePaystackArgs, opts ...pulumi.InvokeOption) (*LookupSourcePaystackResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourcePaystackResult
	err := ctx.Invoke("airbyte:index/getSourcePaystack:getSourcePaystack", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourcePaystack.
type LookupSourcePaystackArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourcePaystack.
type LookupSourcePaystackResult struct {
	Configuration GetSourcePaystackConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourcePaystackOutput(ctx *pulumi.Context, args LookupSourcePaystackOutputArgs, opts ...pulumi.InvokeOption) LookupSourcePaystackResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourcePaystackResult, error) {
			args := v.(LookupSourcePaystackArgs)
			r, err := LookupSourcePaystack(ctx, &args, opts...)
			var s LookupSourcePaystackResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourcePaystackResultOutput)
}

// A collection of arguments for invoking getSourcePaystack.
type LookupSourcePaystackOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourcePaystackOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourcePaystackArgs)(nil)).Elem()
}

// A collection of values returned by getSourcePaystack.
type LookupSourcePaystackResultOutput struct{ *pulumi.OutputState }

func (LookupSourcePaystackResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourcePaystackResult)(nil)).Elem()
}

func (o LookupSourcePaystackResultOutput) ToLookupSourcePaystackResultOutput() LookupSourcePaystackResultOutput {
	return o
}

func (o LookupSourcePaystackResultOutput) ToLookupSourcePaystackResultOutputWithContext(ctx context.Context) LookupSourcePaystackResultOutput {
	return o
}

func (o LookupSourcePaystackResultOutput) Configuration() GetSourcePaystackConfigurationOutput {
	return o.ApplyT(func(v LookupSourcePaystackResult) GetSourcePaystackConfiguration { return v.Configuration }).(GetSourcePaystackConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourcePaystackResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourcePaystackResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourcePaystackResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourcePaystackResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourcePaystackResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourcePaystackResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourcePaystackResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourcePaystackResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourcePaystackResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourcePaystackResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourcePaystackResultOutput{})
}
