// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSourcePexelsApi(ctx *pulumi.Context, args *LookupSourcePexelsApiArgs, opts ...pulumi.InvokeOption) (*LookupSourcePexelsApiResult, error) {
	var rv LookupSourcePexelsApiResult
	err := ctx.Invoke("airbyte:index/getSourcePexelsApi:getSourcePexelsApi", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourcePexelsApi.
type LookupSourcePexelsApiArgs struct {
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourcePexelsApi.
type LookupSourcePexelsApiResult struct {
	Configuration GetSourcePexelsApiConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	Name        string  `pulumi:"name"`
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourcePexelsApiOutput(ctx *pulumi.Context, args LookupSourcePexelsApiOutputArgs, opts ...pulumi.InvokeOption) LookupSourcePexelsApiResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourcePexelsApiResult, error) {
			args := v.(LookupSourcePexelsApiArgs)
			r, err := LookupSourcePexelsApi(ctx, &args, opts...)
			var s LookupSourcePexelsApiResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourcePexelsApiResultOutput)
}

// A collection of arguments for invoking getSourcePexelsApi.
type LookupSourcePexelsApiOutputArgs struct {
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourcePexelsApiOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourcePexelsApiArgs)(nil)).Elem()
}

// A collection of values returned by getSourcePexelsApi.
type LookupSourcePexelsApiResultOutput struct{ *pulumi.OutputState }

func (LookupSourcePexelsApiResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourcePexelsApiResult)(nil)).Elem()
}

func (o LookupSourcePexelsApiResultOutput) ToLookupSourcePexelsApiResultOutput() LookupSourcePexelsApiResultOutput {
	return o
}

func (o LookupSourcePexelsApiResultOutput) ToLookupSourcePexelsApiResultOutputWithContext(ctx context.Context) LookupSourcePexelsApiResultOutput {
	return o
}

func (o LookupSourcePexelsApiResultOutput) Configuration() GetSourcePexelsApiConfigurationOutput {
	return o.ApplyT(func(v LookupSourcePexelsApiResult) GetSourcePexelsApiConfiguration { return v.Configuration }).(GetSourcePexelsApiConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourcePexelsApiResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourcePexelsApiResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourcePexelsApiResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourcePexelsApiResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourcePexelsApiResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourcePexelsApiResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourcePexelsApiResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourcePexelsApiResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourcePexelsApiResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourcePexelsApiResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourcePexelsApiResultOutput{})
}