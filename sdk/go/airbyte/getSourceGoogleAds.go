// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSourceGoogleAds(ctx *pulumi.Context, args *LookupSourceGoogleAdsArgs, opts ...pulumi.InvokeOption) (*LookupSourceGoogleAdsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSourceGoogleAdsResult
	err := ctx.Invoke("airbyte:index/getSourceGoogleAds:getSourceGoogleAds", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceGoogleAds.
type LookupSourceGoogleAdsArgs struct {
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceGoogleAds.
type LookupSourceGoogleAdsResult struct {
	Configuration GetSourceGoogleAdsConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	Name        string  `pulumi:"name"`
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceGoogleAdsOutput(ctx *pulumi.Context, args LookupSourceGoogleAdsOutputArgs, opts ...pulumi.InvokeOption) LookupSourceGoogleAdsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceGoogleAdsResult, error) {
			args := v.(LookupSourceGoogleAdsArgs)
			r, err := LookupSourceGoogleAds(ctx, &args, opts...)
			var s LookupSourceGoogleAdsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceGoogleAdsResultOutput)
}

// A collection of arguments for invoking getSourceGoogleAds.
type LookupSourceGoogleAdsOutputArgs struct {
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceGoogleAdsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceGoogleAdsArgs)(nil)).Elem()
}

// A collection of values returned by getSourceGoogleAds.
type LookupSourceGoogleAdsResultOutput struct{ *pulumi.OutputState }

func (LookupSourceGoogleAdsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceGoogleAdsResult)(nil)).Elem()
}

func (o LookupSourceGoogleAdsResultOutput) ToLookupSourceGoogleAdsResultOutput() LookupSourceGoogleAdsResultOutput {
	return o
}

func (o LookupSourceGoogleAdsResultOutput) ToLookupSourceGoogleAdsResultOutputWithContext(ctx context.Context) LookupSourceGoogleAdsResultOutput {
	return o
}

func (o LookupSourceGoogleAdsResultOutput) Configuration() GetSourceGoogleAdsConfigurationOutput {
	return o.ApplyT(func(v LookupSourceGoogleAdsResult) GetSourceGoogleAdsConfiguration { return v.Configuration }).(GetSourceGoogleAdsConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceGoogleAdsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceGoogleAdsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceGoogleAdsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceGoogleAdsResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceGoogleAdsResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceGoogleAdsResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceGoogleAdsResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceGoogleAdsResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceGoogleAdsResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceGoogleAdsResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceGoogleAdsResultOutput{})
}
