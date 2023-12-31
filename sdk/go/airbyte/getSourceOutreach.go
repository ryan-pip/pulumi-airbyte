// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// SourceOutreach DataSource
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
//			_, err := airbyte.LookupSourceOutreach(ctx, &airbyte.LookupSourceOutreachArgs{
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
func LookupSourceOutreach(ctx *pulumi.Context, args *LookupSourceOutreachArgs, opts ...pulumi.InvokeOption) (*LookupSourceOutreachResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceOutreachResult
	err := ctx.Invoke("airbyte:index/getSourceOutreach:getSourceOutreach", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceOutreach.
type LookupSourceOutreachArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceOutreach.
type LookupSourceOutreachResult struct {
	Configuration GetSourceOutreachConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceOutreachOutput(ctx *pulumi.Context, args LookupSourceOutreachOutputArgs, opts ...pulumi.InvokeOption) LookupSourceOutreachResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceOutreachResult, error) {
			args := v.(LookupSourceOutreachArgs)
			r, err := LookupSourceOutreach(ctx, &args, opts...)
			var s LookupSourceOutreachResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceOutreachResultOutput)
}

// A collection of arguments for invoking getSourceOutreach.
type LookupSourceOutreachOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceOutreachOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceOutreachArgs)(nil)).Elem()
}

// A collection of values returned by getSourceOutreach.
type LookupSourceOutreachResultOutput struct{ *pulumi.OutputState }

func (LookupSourceOutreachResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceOutreachResult)(nil)).Elem()
}

func (o LookupSourceOutreachResultOutput) ToLookupSourceOutreachResultOutput() LookupSourceOutreachResultOutput {
	return o
}

func (o LookupSourceOutreachResultOutput) ToLookupSourceOutreachResultOutputWithContext(ctx context.Context) LookupSourceOutreachResultOutput {
	return o
}

func (o LookupSourceOutreachResultOutput) Configuration() GetSourceOutreachConfigurationOutput {
	return o.ApplyT(func(v LookupSourceOutreachResult) GetSourceOutreachConfiguration { return v.Configuration }).(GetSourceOutreachConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceOutreachResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceOutreachResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceOutreachResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceOutreachResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceOutreachResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceOutreachResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceOutreachResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceOutreachResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceOutreachResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceOutreachResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceOutreachResultOutput{})
}
