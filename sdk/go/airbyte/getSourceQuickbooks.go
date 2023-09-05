// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"internal"
)

// SourceQuickbooks DataSource
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
// _, err := airbyte.LookupSourceQuickbooks(ctx, %!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference), nil);
// if err != nil {
// return err
// }
// return nil
// })
// }
// ```
func LookupSourceQuickbooks(ctx *pulumi.Context, args *LookupSourceQuickbooksArgs, opts ...pulumi.InvokeOption) (*LookupSourceQuickbooksResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceQuickbooksResult
	err := ctx.Invoke("airbyte:index/getSourceQuickbooks:getSourceQuickbooks", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceQuickbooks.
type LookupSourceQuickbooksArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceQuickbooks.
type LookupSourceQuickbooksResult struct {
	Configuration GetSourceQuickbooksConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceQuickbooksOutput(ctx *pulumi.Context, args LookupSourceQuickbooksOutputArgs, opts ...pulumi.InvokeOption) LookupSourceQuickbooksResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceQuickbooksResult, error) {
			args := v.(LookupSourceQuickbooksArgs)
			r, err := LookupSourceQuickbooks(ctx, &args, opts...)
			var s LookupSourceQuickbooksResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceQuickbooksResultOutput)
}

// A collection of arguments for invoking getSourceQuickbooks.
type LookupSourceQuickbooksOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceQuickbooksOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceQuickbooksArgs)(nil)).Elem()
}

// A collection of values returned by getSourceQuickbooks.
type LookupSourceQuickbooksResultOutput struct{ *pulumi.OutputState }

func (LookupSourceQuickbooksResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceQuickbooksResult)(nil)).Elem()
}

func (o LookupSourceQuickbooksResultOutput) ToLookupSourceQuickbooksResultOutput() LookupSourceQuickbooksResultOutput {
	return o
}

func (o LookupSourceQuickbooksResultOutput) ToLookupSourceQuickbooksResultOutputWithContext(ctx context.Context) LookupSourceQuickbooksResultOutput {
	return o
}

func (o LookupSourceQuickbooksResultOutput) Configuration() GetSourceQuickbooksConfigurationOutput {
	return o.ApplyT(func(v LookupSourceQuickbooksResult) GetSourceQuickbooksConfiguration { return v.Configuration }).(GetSourceQuickbooksConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceQuickbooksResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceQuickbooksResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceQuickbooksResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceQuickbooksResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceQuickbooksResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceQuickbooksResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceQuickbooksResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceQuickbooksResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceQuickbooksResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceQuickbooksResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceQuickbooksResultOutput{})
}
