// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// SourceOnesignal DataSource
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
//			_, err := airbyte.LookupSourceOnesignal(ctx, &airbyte.LookupSourceOnesignalArgs{
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
func LookupSourceOnesignal(ctx *pulumi.Context, args *LookupSourceOnesignalArgs, opts ...pulumi.InvokeOption) (*LookupSourceOnesignalResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceOnesignalResult
	err := ctx.Invoke("airbyte:index/getSourceOnesignal:getSourceOnesignal", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceOnesignal.
type LookupSourceOnesignalArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceOnesignal.
type LookupSourceOnesignalResult struct {
	Configuration GetSourceOnesignalConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceOnesignalOutput(ctx *pulumi.Context, args LookupSourceOnesignalOutputArgs, opts ...pulumi.InvokeOption) LookupSourceOnesignalResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceOnesignalResult, error) {
			args := v.(LookupSourceOnesignalArgs)
			r, err := LookupSourceOnesignal(ctx, &args, opts...)
			var s LookupSourceOnesignalResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceOnesignalResultOutput)
}

// A collection of arguments for invoking getSourceOnesignal.
type LookupSourceOnesignalOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceOnesignalOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceOnesignalArgs)(nil)).Elem()
}

// A collection of values returned by getSourceOnesignal.
type LookupSourceOnesignalResultOutput struct{ *pulumi.OutputState }

func (LookupSourceOnesignalResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceOnesignalResult)(nil)).Elem()
}

func (o LookupSourceOnesignalResultOutput) ToLookupSourceOnesignalResultOutput() LookupSourceOnesignalResultOutput {
	return o
}

func (o LookupSourceOnesignalResultOutput) ToLookupSourceOnesignalResultOutputWithContext(ctx context.Context) LookupSourceOnesignalResultOutput {
	return o
}

func (o LookupSourceOnesignalResultOutput) Configuration() GetSourceOnesignalConfigurationOutput {
	return o.ApplyT(func(v LookupSourceOnesignalResult) GetSourceOnesignalConfiguration { return v.Configuration }).(GetSourceOnesignalConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceOnesignalResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceOnesignalResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceOnesignalResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceOnesignalResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceOnesignalResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceOnesignalResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceOnesignalResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceOnesignalResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceOnesignalResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceOnesignalResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceOnesignalResultOutput{})
}
