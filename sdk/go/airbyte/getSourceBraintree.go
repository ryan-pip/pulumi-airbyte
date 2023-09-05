// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-airbyte/sdk/go/airbyte/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SourceBraintree DataSource
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
//			_, err := airbyte.LookupSourceBraintree(ctx, &airbyte.LookupSourceBraintreeArgs{
//				SecretId: pulumi.StringRef("...my_secret_id..."),
//				SourceId: "...my_source_id...",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupSourceBraintree(ctx *pulumi.Context, args *LookupSourceBraintreeArgs, opts ...pulumi.InvokeOption) (*LookupSourceBraintreeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceBraintreeResult
	err := ctx.Invoke("airbyte:index/getSourceBraintree:getSourceBraintree", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceBraintree.
type LookupSourceBraintreeArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceBraintree.
type LookupSourceBraintreeResult struct {
	Configuration GetSourceBraintreeConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceBraintreeOutput(ctx *pulumi.Context, args LookupSourceBraintreeOutputArgs, opts ...pulumi.InvokeOption) LookupSourceBraintreeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceBraintreeResult, error) {
			args := v.(LookupSourceBraintreeArgs)
			r, err := LookupSourceBraintree(ctx, &args, opts...)
			var s LookupSourceBraintreeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceBraintreeResultOutput)
}

// A collection of arguments for invoking getSourceBraintree.
type LookupSourceBraintreeOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceBraintreeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceBraintreeArgs)(nil)).Elem()
}

// A collection of values returned by getSourceBraintree.
type LookupSourceBraintreeResultOutput struct{ *pulumi.OutputState }

func (LookupSourceBraintreeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceBraintreeResult)(nil)).Elem()
}

func (o LookupSourceBraintreeResultOutput) ToLookupSourceBraintreeResultOutput() LookupSourceBraintreeResultOutput {
	return o
}

func (o LookupSourceBraintreeResultOutput) ToLookupSourceBraintreeResultOutputWithContext(ctx context.Context) LookupSourceBraintreeResultOutput {
	return o
}

func (o LookupSourceBraintreeResultOutput) Configuration() GetSourceBraintreeConfigurationOutput {
	return o.ApplyT(func(v LookupSourceBraintreeResult) GetSourceBraintreeConfiguration { return v.Configuration }).(GetSourceBraintreeConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceBraintreeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceBraintreeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceBraintreeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceBraintreeResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceBraintreeResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceBraintreeResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceBraintreeResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceBraintreeResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceBraintreeResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceBraintreeResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceBraintreeResultOutput{})
}
