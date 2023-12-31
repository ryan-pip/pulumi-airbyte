// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// SourceCoinAPI DataSource
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
//			_, err := airbyte.LookupSourceCoinApi(ctx, &airbyte.LookupSourceCoinApiArgs{
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
func LookupSourceCoinApi(ctx *pulumi.Context, args *LookupSourceCoinApiArgs, opts ...pulumi.InvokeOption) (*LookupSourceCoinApiResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceCoinApiResult
	err := ctx.Invoke("airbyte:index/getSourceCoinApi:getSourceCoinApi", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceCoinApi.
type LookupSourceCoinApiArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceCoinApi.
type LookupSourceCoinApiResult struct {
	Configuration GetSourceCoinApiConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceCoinApiOutput(ctx *pulumi.Context, args LookupSourceCoinApiOutputArgs, opts ...pulumi.InvokeOption) LookupSourceCoinApiResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceCoinApiResult, error) {
			args := v.(LookupSourceCoinApiArgs)
			r, err := LookupSourceCoinApi(ctx, &args, opts...)
			var s LookupSourceCoinApiResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceCoinApiResultOutput)
}

// A collection of arguments for invoking getSourceCoinApi.
type LookupSourceCoinApiOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceCoinApiOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceCoinApiArgs)(nil)).Elem()
}

// A collection of values returned by getSourceCoinApi.
type LookupSourceCoinApiResultOutput struct{ *pulumi.OutputState }

func (LookupSourceCoinApiResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceCoinApiResult)(nil)).Elem()
}

func (o LookupSourceCoinApiResultOutput) ToLookupSourceCoinApiResultOutput() LookupSourceCoinApiResultOutput {
	return o
}

func (o LookupSourceCoinApiResultOutput) ToLookupSourceCoinApiResultOutputWithContext(ctx context.Context) LookupSourceCoinApiResultOutput {
	return o
}

func (o LookupSourceCoinApiResultOutput) Configuration() GetSourceCoinApiConfigurationOutput {
	return o.ApplyT(func(v LookupSourceCoinApiResult) GetSourceCoinApiConfiguration { return v.Configuration }).(GetSourceCoinApiConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceCoinApiResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceCoinApiResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceCoinApiResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceCoinApiResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceCoinApiResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceCoinApiResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceCoinApiResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceCoinApiResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceCoinApiResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceCoinApiResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceCoinApiResultOutput{})
}
