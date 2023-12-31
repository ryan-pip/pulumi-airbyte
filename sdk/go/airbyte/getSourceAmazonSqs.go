// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// SourceAmazonSqs DataSource
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
//			_, err := airbyte.LookupSourceAmazonSqs(ctx, &airbyte.LookupSourceAmazonSqsArgs{
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
func LookupSourceAmazonSqs(ctx *pulumi.Context, args *LookupSourceAmazonSqsArgs, opts ...pulumi.InvokeOption) (*LookupSourceAmazonSqsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceAmazonSqsResult
	err := ctx.Invoke("airbyte:index/getSourceAmazonSqs:getSourceAmazonSqs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceAmazonSqs.
type LookupSourceAmazonSqsArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceAmazonSqs.
type LookupSourceAmazonSqsResult struct {
	Configuration GetSourceAmazonSqsConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceAmazonSqsOutput(ctx *pulumi.Context, args LookupSourceAmazonSqsOutputArgs, opts ...pulumi.InvokeOption) LookupSourceAmazonSqsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceAmazonSqsResult, error) {
			args := v.(LookupSourceAmazonSqsArgs)
			r, err := LookupSourceAmazonSqs(ctx, &args, opts...)
			var s LookupSourceAmazonSqsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceAmazonSqsResultOutput)
}

// A collection of arguments for invoking getSourceAmazonSqs.
type LookupSourceAmazonSqsOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceAmazonSqsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceAmazonSqsArgs)(nil)).Elem()
}

// A collection of values returned by getSourceAmazonSqs.
type LookupSourceAmazonSqsResultOutput struct{ *pulumi.OutputState }

func (LookupSourceAmazonSqsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceAmazonSqsResult)(nil)).Elem()
}

func (o LookupSourceAmazonSqsResultOutput) ToLookupSourceAmazonSqsResultOutput() LookupSourceAmazonSqsResultOutput {
	return o
}

func (o LookupSourceAmazonSqsResultOutput) ToLookupSourceAmazonSqsResultOutputWithContext(ctx context.Context) LookupSourceAmazonSqsResultOutput {
	return o
}

func (o LookupSourceAmazonSqsResultOutput) Configuration() GetSourceAmazonSqsConfigurationOutput {
	return o.ApplyT(func(v LookupSourceAmazonSqsResult) GetSourceAmazonSqsConfiguration { return v.Configuration }).(GetSourceAmazonSqsConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceAmazonSqsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceAmazonSqsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceAmazonSqsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceAmazonSqsResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceAmazonSqsResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceAmazonSqsResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceAmazonSqsResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceAmazonSqsResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceAmazonSqsResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceAmazonSqsResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceAmazonSqsResultOutput{})
}
