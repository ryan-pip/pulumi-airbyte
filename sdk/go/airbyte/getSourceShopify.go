// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// SourceShopify DataSource
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
//			_, err := airbyte.LookupSourceShopify(ctx, &airbyte.LookupSourceShopifyArgs{
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
func LookupSourceShopify(ctx *pulumi.Context, args *LookupSourceShopifyArgs, opts ...pulumi.InvokeOption) (*LookupSourceShopifyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceShopifyResult
	err := ctx.Invoke("airbyte:index/getSourceShopify:getSourceShopify", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceShopify.
type LookupSourceShopifyArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceShopify.
type LookupSourceShopifyResult struct {
	Configuration GetSourceShopifyConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceShopifyOutput(ctx *pulumi.Context, args LookupSourceShopifyOutputArgs, opts ...pulumi.InvokeOption) LookupSourceShopifyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceShopifyResult, error) {
			args := v.(LookupSourceShopifyArgs)
			r, err := LookupSourceShopify(ctx, &args, opts...)
			var s LookupSourceShopifyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceShopifyResultOutput)
}

// A collection of arguments for invoking getSourceShopify.
type LookupSourceShopifyOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceShopifyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceShopifyArgs)(nil)).Elem()
}

// A collection of values returned by getSourceShopify.
type LookupSourceShopifyResultOutput struct{ *pulumi.OutputState }

func (LookupSourceShopifyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceShopifyResult)(nil)).Elem()
}

func (o LookupSourceShopifyResultOutput) ToLookupSourceShopifyResultOutput() LookupSourceShopifyResultOutput {
	return o
}

func (o LookupSourceShopifyResultOutput) ToLookupSourceShopifyResultOutputWithContext(ctx context.Context) LookupSourceShopifyResultOutput {
	return o
}

func (o LookupSourceShopifyResultOutput) Configuration() GetSourceShopifyConfigurationOutput {
	return o.ApplyT(func(v LookupSourceShopifyResult) GetSourceShopifyConfiguration { return v.Configuration }).(GetSourceShopifyConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceShopifyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceShopifyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceShopifyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceShopifyResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceShopifyResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceShopifyResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceShopifyResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceShopifyResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceShopifyResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceShopifyResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceShopifyResultOutput{})
}
