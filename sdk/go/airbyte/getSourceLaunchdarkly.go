// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// SourceLaunchdarkly DataSource
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
//			_, err := airbyte.LookupSourceLaunchdarkly(ctx, &airbyte.LookupSourceLaunchdarklyArgs{
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
func LookupSourceLaunchdarkly(ctx *pulumi.Context, args *LookupSourceLaunchdarklyArgs, opts ...pulumi.InvokeOption) (*LookupSourceLaunchdarklyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceLaunchdarklyResult
	err := ctx.Invoke("airbyte:index/getSourceLaunchdarkly:getSourceLaunchdarkly", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceLaunchdarkly.
type LookupSourceLaunchdarklyArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceLaunchdarkly.
type LookupSourceLaunchdarklyResult struct {
	Configuration GetSourceLaunchdarklyConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceLaunchdarklyOutput(ctx *pulumi.Context, args LookupSourceLaunchdarklyOutputArgs, opts ...pulumi.InvokeOption) LookupSourceLaunchdarklyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceLaunchdarklyResult, error) {
			args := v.(LookupSourceLaunchdarklyArgs)
			r, err := LookupSourceLaunchdarkly(ctx, &args, opts...)
			var s LookupSourceLaunchdarklyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceLaunchdarklyResultOutput)
}

// A collection of arguments for invoking getSourceLaunchdarkly.
type LookupSourceLaunchdarklyOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceLaunchdarklyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceLaunchdarklyArgs)(nil)).Elem()
}

// A collection of values returned by getSourceLaunchdarkly.
type LookupSourceLaunchdarklyResultOutput struct{ *pulumi.OutputState }

func (LookupSourceLaunchdarklyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceLaunchdarklyResult)(nil)).Elem()
}

func (o LookupSourceLaunchdarklyResultOutput) ToLookupSourceLaunchdarklyResultOutput() LookupSourceLaunchdarklyResultOutput {
	return o
}

func (o LookupSourceLaunchdarklyResultOutput) ToLookupSourceLaunchdarklyResultOutputWithContext(ctx context.Context) LookupSourceLaunchdarklyResultOutput {
	return o
}

func (o LookupSourceLaunchdarklyResultOutput) Configuration() GetSourceLaunchdarklyConfigurationOutput {
	return o.ApplyT(func(v LookupSourceLaunchdarklyResult) GetSourceLaunchdarklyConfiguration { return v.Configuration }).(GetSourceLaunchdarklyConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceLaunchdarklyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceLaunchdarklyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceLaunchdarklyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceLaunchdarklyResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceLaunchdarklyResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceLaunchdarklyResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceLaunchdarklyResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceLaunchdarklyResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceLaunchdarklyResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceLaunchdarklyResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceLaunchdarklyResultOutput{})
}
