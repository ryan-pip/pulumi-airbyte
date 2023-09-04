// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSourcePolygonStockApi(ctx *pulumi.Context, args *LookupSourcePolygonStockApiArgs, opts ...pulumi.InvokeOption) (*LookupSourcePolygonStockApiResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSourcePolygonStockApiResult
	err := ctx.Invoke("airbyte:index/getSourcePolygonStockApi:getSourcePolygonStockApi", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourcePolygonStockApi.
type LookupSourcePolygonStockApiArgs struct {
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourcePolygonStockApi.
type LookupSourcePolygonStockApiResult struct {
	Configuration GetSourcePolygonStockApiConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	Name        string  `pulumi:"name"`
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourcePolygonStockApiOutput(ctx *pulumi.Context, args LookupSourcePolygonStockApiOutputArgs, opts ...pulumi.InvokeOption) LookupSourcePolygonStockApiResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourcePolygonStockApiResult, error) {
			args := v.(LookupSourcePolygonStockApiArgs)
			r, err := LookupSourcePolygonStockApi(ctx, &args, opts...)
			var s LookupSourcePolygonStockApiResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourcePolygonStockApiResultOutput)
}

// A collection of arguments for invoking getSourcePolygonStockApi.
type LookupSourcePolygonStockApiOutputArgs struct {
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourcePolygonStockApiOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourcePolygonStockApiArgs)(nil)).Elem()
}

// A collection of values returned by getSourcePolygonStockApi.
type LookupSourcePolygonStockApiResultOutput struct{ *pulumi.OutputState }

func (LookupSourcePolygonStockApiResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourcePolygonStockApiResult)(nil)).Elem()
}

func (o LookupSourcePolygonStockApiResultOutput) ToLookupSourcePolygonStockApiResultOutput() LookupSourcePolygonStockApiResultOutput {
	return o
}

func (o LookupSourcePolygonStockApiResultOutput) ToLookupSourcePolygonStockApiResultOutputWithContext(ctx context.Context) LookupSourcePolygonStockApiResultOutput {
	return o
}

func (o LookupSourcePolygonStockApiResultOutput) Configuration() GetSourcePolygonStockApiConfigurationOutput {
	return o.ApplyT(func(v LookupSourcePolygonStockApiResult) GetSourcePolygonStockApiConfiguration {
		return v.Configuration
	}).(GetSourcePolygonStockApiConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourcePolygonStockApiResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourcePolygonStockApiResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourcePolygonStockApiResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourcePolygonStockApiResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourcePolygonStockApiResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourcePolygonStockApiResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourcePolygonStockApiResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourcePolygonStockApiResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourcePolygonStockApiResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourcePolygonStockApiResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourcePolygonStockApiResultOutput{})
}
