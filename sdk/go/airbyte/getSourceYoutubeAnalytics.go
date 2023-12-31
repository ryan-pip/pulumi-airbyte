// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// SourceYoutubeAnalytics DataSource
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
//			_, err := airbyte.LookupSourceYoutubeAnalytics(ctx, &airbyte.LookupSourceYoutubeAnalyticsArgs{
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
func LookupSourceYoutubeAnalytics(ctx *pulumi.Context, args *LookupSourceYoutubeAnalyticsArgs, opts ...pulumi.InvokeOption) (*LookupSourceYoutubeAnalyticsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceYoutubeAnalyticsResult
	err := ctx.Invoke("airbyte:index/getSourceYoutubeAnalytics:getSourceYoutubeAnalytics", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceYoutubeAnalytics.
type LookupSourceYoutubeAnalyticsArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceYoutubeAnalytics.
type LookupSourceYoutubeAnalyticsResult struct {
	Configuration GetSourceYoutubeAnalyticsConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceYoutubeAnalyticsOutput(ctx *pulumi.Context, args LookupSourceYoutubeAnalyticsOutputArgs, opts ...pulumi.InvokeOption) LookupSourceYoutubeAnalyticsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceYoutubeAnalyticsResult, error) {
			args := v.(LookupSourceYoutubeAnalyticsArgs)
			r, err := LookupSourceYoutubeAnalytics(ctx, &args, opts...)
			var s LookupSourceYoutubeAnalyticsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceYoutubeAnalyticsResultOutput)
}

// A collection of arguments for invoking getSourceYoutubeAnalytics.
type LookupSourceYoutubeAnalyticsOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceYoutubeAnalyticsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceYoutubeAnalyticsArgs)(nil)).Elem()
}

// A collection of values returned by getSourceYoutubeAnalytics.
type LookupSourceYoutubeAnalyticsResultOutput struct{ *pulumi.OutputState }

func (LookupSourceYoutubeAnalyticsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceYoutubeAnalyticsResult)(nil)).Elem()
}

func (o LookupSourceYoutubeAnalyticsResultOutput) ToLookupSourceYoutubeAnalyticsResultOutput() LookupSourceYoutubeAnalyticsResultOutput {
	return o
}

func (o LookupSourceYoutubeAnalyticsResultOutput) ToLookupSourceYoutubeAnalyticsResultOutputWithContext(ctx context.Context) LookupSourceYoutubeAnalyticsResultOutput {
	return o
}

func (o LookupSourceYoutubeAnalyticsResultOutput) Configuration() GetSourceYoutubeAnalyticsConfigurationOutput {
	return o.ApplyT(func(v LookupSourceYoutubeAnalyticsResult) GetSourceYoutubeAnalyticsConfiguration {
		return v.Configuration
	}).(GetSourceYoutubeAnalyticsConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceYoutubeAnalyticsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceYoutubeAnalyticsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceYoutubeAnalyticsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceYoutubeAnalyticsResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceYoutubeAnalyticsResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceYoutubeAnalyticsResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceYoutubeAnalyticsResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceYoutubeAnalyticsResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceYoutubeAnalyticsResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceYoutubeAnalyticsResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceYoutubeAnalyticsResultOutput{})
}
