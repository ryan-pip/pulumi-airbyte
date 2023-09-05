// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// SourceWoocommerce DataSource
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
//			_, err := airbyte.LookupSourceWoocommerce(ctx, &airbyte.LookupSourceWoocommerceArgs{
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
func LookupSourceWoocommerce(ctx *pulumi.Context, args *LookupSourceWoocommerceArgs, opts ...pulumi.InvokeOption) (*LookupSourceWoocommerceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceWoocommerceResult
	err := ctx.Invoke("airbyte:index/getSourceWoocommerce:getSourceWoocommerce", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceWoocommerce.
type LookupSourceWoocommerceArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceWoocommerce.
type LookupSourceWoocommerceResult struct {
	Configuration GetSourceWoocommerceConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceWoocommerceOutput(ctx *pulumi.Context, args LookupSourceWoocommerceOutputArgs, opts ...pulumi.InvokeOption) LookupSourceWoocommerceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceWoocommerceResult, error) {
			args := v.(LookupSourceWoocommerceArgs)
			r, err := LookupSourceWoocommerce(ctx, &args, opts...)
			var s LookupSourceWoocommerceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceWoocommerceResultOutput)
}

// A collection of arguments for invoking getSourceWoocommerce.
type LookupSourceWoocommerceOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceWoocommerceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceWoocommerceArgs)(nil)).Elem()
}

// A collection of values returned by getSourceWoocommerce.
type LookupSourceWoocommerceResultOutput struct{ *pulumi.OutputState }

func (LookupSourceWoocommerceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceWoocommerceResult)(nil)).Elem()
}

func (o LookupSourceWoocommerceResultOutput) ToLookupSourceWoocommerceResultOutput() LookupSourceWoocommerceResultOutput {
	return o
}

func (o LookupSourceWoocommerceResultOutput) ToLookupSourceWoocommerceResultOutputWithContext(ctx context.Context) LookupSourceWoocommerceResultOutput {
	return o
}

func (o LookupSourceWoocommerceResultOutput) Configuration() GetSourceWoocommerceConfigurationOutput {
	return o.ApplyT(func(v LookupSourceWoocommerceResult) GetSourceWoocommerceConfiguration { return v.Configuration }).(GetSourceWoocommerceConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceWoocommerceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceWoocommerceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceWoocommerceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceWoocommerceResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceWoocommerceResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceWoocommerceResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceWoocommerceResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceWoocommerceResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceWoocommerceResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceWoocommerceResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceWoocommerceResultOutput{})
}
