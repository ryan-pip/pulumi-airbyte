// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// SourceZendeskSupport DataSource
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
//			_, err := airbyte.LookupSourceZendeskSupport(ctx, &airbyte.LookupSourceZendeskSupportArgs{
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
func LookupSourceZendeskSupport(ctx *pulumi.Context, args *LookupSourceZendeskSupportArgs, opts ...pulumi.InvokeOption) (*LookupSourceZendeskSupportResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceZendeskSupportResult
	err := ctx.Invoke("airbyte:index/getSourceZendeskSupport:getSourceZendeskSupport", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceZendeskSupport.
type LookupSourceZendeskSupportArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceZendeskSupport.
type LookupSourceZendeskSupportResult struct {
	Configuration GetSourceZendeskSupportConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceZendeskSupportOutput(ctx *pulumi.Context, args LookupSourceZendeskSupportOutputArgs, opts ...pulumi.InvokeOption) LookupSourceZendeskSupportResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceZendeskSupportResult, error) {
			args := v.(LookupSourceZendeskSupportArgs)
			r, err := LookupSourceZendeskSupport(ctx, &args, opts...)
			var s LookupSourceZendeskSupportResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceZendeskSupportResultOutput)
}

// A collection of arguments for invoking getSourceZendeskSupport.
type LookupSourceZendeskSupportOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceZendeskSupportOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceZendeskSupportArgs)(nil)).Elem()
}

// A collection of values returned by getSourceZendeskSupport.
type LookupSourceZendeskSupportResultOutput struct{ *pulumi.OutputState }

func (LookupSourceZendeskSupportResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceZendeskSupportResult)(nil)).Elem()
}

func (o LookupSourceZendeskSupportResultOutput) ToLookupSourceZendeskSupportResultOutput() LookupSourceZendeskSupportResultOutput {
	return o
}

func (o LookupSourceZendeskSupportResultOutput) ToLookupSourceZendeskSupportResultOutputWithContext(ctx context.Context) LookupSourceZendeskSupportResultOutput {
	return o
}

func (o LookupSourceZendeskSupportResultOutput) Configuration() GetSourceZendeskSupportConfigurationOutput {
	return o.ApplyT(func(v LookupSourceZendeskSupportResult) GetSourceZendeskSupportConfiguration { return v.Configuration }).(GetSourceZendeskSupportConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceZendeskSupportResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceZendeskSupportResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceZendeskSupportResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceZendeskSupportResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceZendeskSupportResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceZendeskSupportResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceZendeskSupportResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceZendeskSupportResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceZendeskSupportResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceZendeskSupportResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceZendeskSupportResultOutput{})
}
