// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// SourcePrestashop DataSource
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
//			_, err := airbyte.LookupSourcePrestashop(ctx, &airbyte.LookupSourcePrestashopArgs{
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
func LookupSourcePrestashop(ctx *pulumi.Context, args *LookupSourcePrestashopArgs, opts ...pulumi.InvokeOption) (*LookupSourcePrestashopResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourcePrestashopResult
	err := ctx.Invoke("airbyte:index/getSourcePrestashop:getSourcePrestashop", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourcePrestashop.
type LookupSourcePrestashopArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourcePrestashop.
type LookupSourcePrestashopResult struct {
	Configuration GetSourcePrestashopConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourcePrestashopOutput(ctx *pulumi.Context, args LookupSourcePrestashopOutputArgs, opts ...pulumi.InvokeOption) LookupSourcePrestashopResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourcePrestashopResult, error) {
			args := v.(LookupSourcePrestashopArgs)
			r, err := LookupSourcePrestashop(ctx, &args, opts...)
			var s LookupSourcePrestashopResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourcePrestashopResultOutput)
}

// A collection of arguments for invoking getSourcePrestashop.
type LookupSourcePrestashopOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourcePrestashopOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourcePrestashopArgs)(nil)).Elem()
}

// A collection of values returned by getSourcePrestashop.
type LookupSourcePrestashopResultOutput struct{ *pulumi.OutputState }

func (LookupSourcePrestashopResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourcePrestashopResult)(nil)).Elem()
}

func (o LookupSourcePrestashopResultOutput) ToLookupSourcePrestashopResultOutput() LookupSourcePrestashopResultOutput {
	return o
}

func (o LookupSourcePrestashopResultOutput) ToLookupSourcePrestashopResultOutputWithContext(ctx context.Context) LookupSourcePrestashopResultOutput {
	return o
}

func (o LookupSourcePrestashopResultOutput) Configuration() GetSourcePrestashopConfigurationOutput {
	return o.ApplyT(func(v LookupSourcePrestashopResult) GetSourcePrestashopConfiguration { return v.Configuration }).(GetSourcePrestashopConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourcePrestashopResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourcePrestashopResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourcePrestashopResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourcePrestashopResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourcePrestashopResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourcePrestashopResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourcePrestashopResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourcePrestashopResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourcePrestashopResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourcePrestashopResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourcePrestashopResultOutput{})
}
