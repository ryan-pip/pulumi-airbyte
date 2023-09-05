// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"internal"
)

// SourceMailjetSms DataSource
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
// _, err := airbyte.LookupSourceMailjetSms(ctx, %!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference), nil);
// if err != nil {
// return err
// }
// return nil
// })
// }
// ```
func LookupSourceMailjetSms(ctx *pulumi.Context, args *LookupSourceMailjetSmsArgs, opts ...pulumi.InvokeOption) (*LookupSourceMailjetSmsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceMailjetSmsResult
	err := ctx.Invoke("airbyte:index/getSourceMailjetSms:getSourceMailjetSms", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceMailjetSms.
type LookupSourceMailjetSmsArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceMailjetSms.
type LookupSourceMailjetSmsResult struct {
	Configuration GetSourceMailjetSmsConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceMailjetSmsOutput(ctx *pulumi.Context, args LookupSourceMailjetSmsOutputArgs, opts ...pulumi.InvokeOption) LookupSourceMailjetSmsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceMailjetSmsResult, error) {
			args := v.(LookupSourceMailjetSmsArgs)
			r, err := LookupSourceMailjetSms(ctx, &args, opts...)
			var s LookupSourceMailjetSmsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceMailjetSmsResultOutput)
}

// A collection of arguments for invoking getSourceMailjetSms.
type LookupSourceMailjetSmsOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceMailjetSmsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceMailjetSmsArgs)(nil)).Elem()
}

// A collection of values returned by getSourceMailjetSms.
type LookupSourceMailjetSmsResultOutput struct{ *pulumi.OutputState }

func (LookupSourceMailjetSmsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceMailjetSmsResult)(nil)).Elem()
}

func (o LookupSourceMailjetSmsResultOutput) ToLookupSourceMailjetSmsResultOutput() LookupSourceMailjetSmsResultOutput {
	return o
}

func (o LookupSourceMailjetSmsResultOutput) ToLookupSourceMailjetSmsResultOutputWithContext(ctx context.Context) LookupSourceMailjetSmsResultOutput {
	return o
}

func (o LookupSourceMailjetSmsResultOutput) Configuration() GetSourceMailjetSmsConfigurationOutput {
	return o.ApplyT(func(v LookupSourceMailjetSmsResult) GetSourceMailjetSmsConfiguration { return v.Configuration }).(GetSourceMailjetSmsConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceMailjetSmsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceMailjetSmsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceMailjetSmsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceMailjetSmsResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceMailjetSmsResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceMailjetSmsResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceMailjetSmsResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceMailjetSmsResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceMailjetSmsResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceMailjetSmsResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceMailjetSmsResultOutput{})
}
