// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// SourceSpacexAPI DataSource
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
//			_, err := airbyte.LookupSourceSpaceXApi(ctx, &airbyte.LookupSourceSpaceXApiArgs{
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
func LookupSourceSpaceXApi(ctx *pulumi.Context, args *LookupSourceSpaceXApiArgs, opts ...pulumi.InvokeOption) (*LookupSourceSpaceXApiResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceSpaceXApiResult
	err := ctx.Invoke("airbyte:index/getSourceSpaceXApi:getSourceSpaceXApi", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceSpaceXApi.
type LookupSourceSpaceXApiArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceSpaceXApi.
type LookupSourceSpaceXApiResult struct {
	Configuration GetSourceSpaceXApiConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceSpaceXApiOutput(ctx *pulumi.Context, args LookupSourceSpaceXApiOutputArgs, opts ...pulumi.InvokeOption) LookupSourceSpaceXApiResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceSpaceXApiResult, error) {
			args := v.(LookupSourceSpaceXApiArgs)
			r, err := LookupSourceSpaceXApi(ctx, &args, opts...)
			var s LookupSourceSpaceXApiResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceSpaceXApiResultOutput)
}

// A collection of arguments for invoking getSourceSpaceXApi.
type LookupSourceSpaceXApiOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceSpaceXApiOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceSpaceXApiArgs)(nil)).Elem()
}

// A collection of values returned by getSourceSpaceXApi.
type LookupSourceSpaceXApiResultOutput struct{ *pulumi.OutputState }

func (LookupSourceSpaceXApiResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceSpaceXApiResult)(nil)).Elem()
}

func (o LookupSourceSpaceXApiResultOutput) ToLookupSourceSpaceXApiResultOutput() LookupSourceSpaceXApiResultOutput {
	return o
}

func (o LookupSourceSpaceXApiResultOutput) ToLookupSourceSpaceXApiResultOutputWithContext(ctx context.Context) LookupSourceSpaceXApiResultOutput {
	return o
}

func (o LookupSourceSpaceXApiResultOutput) Configuration() GetSourceSpaceXApiConfigurationOutput {
	return o.ApplyT(func(v LookupSourceSpaceXApiResult) GetSourceSpaceXApiConfiguration { return v.Configuration }).(GetSourceSpaceXApiConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceSpaceXApiResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSpaceXApiResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceSpaceXApiResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSpaceXApiResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceSpaceXApiResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceSpaceXApiResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceSpaceXApiResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSpaceXApiResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceSpaceXApiResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSpaceXApiResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceSpaceXApiResultOutput{})
}
