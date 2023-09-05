// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// SourceDixa DataSource
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
//			_, err := airbyte.LookupSourceDixa(ctx, &airbyte.LookupSourceDixaArgs{
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
func LookupSourceDixa(ctx *pulumi.Context, args *LookupSourceDixaArgs, opts ...pulumi.InvokeOption) (*LookupSourceDixaResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceDixaResult
	err := ctx.Invoke("airbyte:index/getSourceDixa:getSourceDixa", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceDixa.
type LookupSourceDixaArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceDixa.
type LookupSourceDixaResult struct {
	Configuration GetSourceDixaConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceDixaOutput(ctx *pulumi.Context, args LookupSourceDixaOutputArgs, opts ...pulumi.InvokeOption) LookupSourceDixaResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceDixaResult, error) {
			args := v.(LookupSourceDixaArgs)
			r, err := LookupSourceDixa(ctx, &args, opts...)
			var s LookupSourceDixaResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceDixaResultOutput)
}

// A collection of arguments for invoking getSourceDixa.
type LookupSourceDixaOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceDixaOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceDixaArgs)(nil)).Elem()
}

// A collection of values returned by getSourceDixa.
type LookupSourceDixaResultOutput struct{ *pulumi.OutputState }

func (LookupSourceDixaResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceDixaResult)(nil)).Elem()
}

func (o LookupSourceDixaResultOutput) ToLookupSourceDixaResultOutput() LookupSourceDixaResultOutput {
	return o
}

func (o LookupSourceDixaResultOutput) ToLookupSourceDixaResultOutputWithContext(ctx context.Context) LookupSourceDixaResultOutput {
	return o
}

func (o LookupSourceDixaResultOutput) Configuration() GetSourceDixaConfigurationOutput {
	return o.ApplyT(func(v LookupSourceDixaResult) GetSourceDixaConfiguration { return v.Configuration }).(GetSourceDixaConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceDixaResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceDixaResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceDixaResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceDixaResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceDixaResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceDixaResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceDixaResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceDixaResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceDixaResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceDixaResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceDixaResultOutput{})
}
