// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-airbyte/sdk/go/airbyte/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SourceClickupAPI DataSource
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-airbyte/sdk/go/airbyte"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := airbyte.LookupSourceClickupApi(ctx, &airbyte.LookupSourceClickupApiArgs{
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
func LookupSourceClickupApi(ctx *pulumi.Context, args *LookupSourceClickupApiArgs, opts ...pulumi.InvokeOption) (*LookupSourceClickupApiResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceClickupApiResult
	err := ctx.Invoke("airbyte:index/getSourceClickupApi:getSourceClickupApi", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceClickupApi.
type LookupSourceClickupApiArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceClickupApi.
type LookupSourceClickupApiResult struct {
	Configuration GetSourceClickupApiConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceClickupApiOutput(ctx *pulumi.Context, args LookupSourceClickupApiOutputArgs, opts ...pulumi.InvokeOption) LookupSourceClickupApiResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceClickupApiResult, error) {
			args := v.(LookupSourceClickupApiArgs)
			r, err := LookupSourceClickupApi(ctx, &args, opts...)
			var s LookupSourceClickupApiResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceClickupApiResultOutput)
}

// A collection of arguments for invoking getSourceClickupApi.
type LookupSourceClickupApiOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceClickupApiOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceClickupApiArgs)(nil)).Elem()
}

// A collection of values returned by getSourceClickupApi.
type LookupSourceClickupApiResultOutput struct{ *pulumi.OutputState }

func (LookupSourceClickupApiResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceClickupApiResult)(nil)).Elem()
}

func (o LookupSourceClickupApiResultOutput) ToLookupSourceClickupApiResultOutput() LookupSourceClickupApiResultOutput {
	return o
}

func (o LookupSourceClickupApiResultOutput) ToLookupSourceClickupApiResultOutputWithContext(ctx context.Context) LookupSourceClickupApiResultOutput {
	return o
}

func (o LookupSourceClickupApiResultOutput) Configuration() GetSourceClickupApiConfigurationOutput {
	return o.ApplyT(func(v LookupSourceClickupApiResult) GetSourceClickupApiConfiguration { return v.Configuration }).(GetSourceClickupApiConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceClickupApiResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceClickupApiResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceClickupApiResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceClickupApiResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceClickupApiResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceClickupApiResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceClickupApiResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceClickupApiResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceClickupApiResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceClickupApiResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceClickupApiResultOutput{})
}
