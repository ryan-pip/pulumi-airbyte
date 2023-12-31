// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// SourceRss DataSource
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
//			_, err := airbyte.LookupSourceRss(ctx, &airbyte.LookupSourceRssArgs{
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
func LookupSourceRss(ctx *pulumi.Context, args *LookupSourceRssArgs, opts ...pulumi.InvokeOption) (*LookupSourceRssResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceRssResult
	err := ctx.Invoke("airbyte:index/getSourceRss:getSourceRss", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceRss.
type LookupSourceRssArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceRss.
type LookupSourceRssResult struct {
	Configuration GetSourceRssConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceRssOutput(ctx *pulumi.Context, args LookupSourceRssOutputArgs, opts ...pulumi.InvokeOption) LookupSourceRssResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceRssResult, error) {
			args := v.(LookupSourceRssArgs)
			r, err := LookupSourceRss(ctx, &args, opts...)
			var s LookupSourceRssResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceRssResultOutput)
}

// A collection of arguments for invoking getSourceRss.
type LookupSourceRssOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceRssOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceRssArgs)(nil)).Elem()
}

// A collection of values returned by getSourceRss.
type LookupSourceRssResultOutput struct{ *pulumi.OutputState }

func (LookupSourceRssResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceRssResult)(nil)).Elem()
}

func (o LookupSourceRssResultOutput) ToLookupSourceRssResultOutput() LookupSourceRssResultOutput {
	return o
}

func (o LookupSourceRssResultOutput) ToLookupSourceRssResultOutputWithContext(ctx context.Context) LookupSourceRssResultOutput {
	return o
}

func (o LookupSourceRssResultOutput) Configuration() GetSourceRssConfigurationOutput {
	return o.ApplyT(func(v LookupSourceRssResult) GetSourceRssConfiguration { return v.Configuration }).(GetSourceRssConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceRssResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceRssResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceRssResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceRssResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceRssResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceRssResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceRssResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceRssResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceRssResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceRssResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceRssResultOutput{})
}
