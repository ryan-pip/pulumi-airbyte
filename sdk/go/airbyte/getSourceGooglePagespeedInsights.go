// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSourceGooglePagespeedInsights(ctx *pulumi.Context, args *LookupSourceGooglePagespeedInsightsArgs, opts ...pulumi.InvokeOption) (*LookupSourceGooglePagespeedInsightsResult, error) {
	var rv LookupSourceGooglePagespeedInsightsResult
	err := ctx.Invoke("airbyte:index/getSourceGooglePagespeedInsights:getSourceGooglePagespeedInsights", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceGooglePagespeedInsights.
type LookupSourceGooglePagespeedInsightsArgs struct {
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceGooglePagespeedInsights.
type LookupSourceGooglePagespeedInsightsResult struct {
	Configuration GetSourceGooglePagespeedInsightsConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	Name        string  `pulumi:"name"`
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceGooglePagespeedInsightsOutput(ctx *pulumi.Context, args LookupSourceGooglePagespeedInsightsOutputArgs, opts ...pulumi.InvokeOption) LookupSourceGooglePagespeedInsightsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceGooglePagespeedInsightsResult, error) {
			args := v.(LookupSourceGooglePagespeedInsightsArgs)
			r, err := LookupSourceGooglePagespeedInsights(ctx, &args, opts...)
			var s LookupSourceGooglePagespeedInsightsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceGooglePagespeedInsightsResultOutput)
}

// A collection of arguments for invoking getSourceGooglePagespeedInsights.
type LookupSourceGooglePagespeedInsightsOutputArgs struct {
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceGooglePagespeedInsightsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceGooglePagespeedInsightsArgs)(nil)).Elem()
}

// A collection of values returned by getSourceGooglePagespeedInsights.
type LookupSourceGooglePagespeedInsightsResultOutput struct{ *pulumi.OutputState }

func (LookupSourceGooglePagespeedInsightsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceGooglePagespeedInsightsResult)(nil)).Elem()
}

func (o LookupSourceGooglePagespeedInsightsResultOutput) ToLookupSourceGooglePagespeedInsightsResultOutput() LookupSourceGooglePagespeedInsightsResultOutput {
	return o
}

func (o LookupSourceGooglePagespeedInsightsResultOutput) ToLookupSourceGooglePagespeedInsightsResultOutputWithContext(ctx context.Context) LookupSourceGooglePagespeedInsightsResultOutput {
	return o
}

func (o LookupSourceGooglePagespeedInsightsResultOutput) Configuration() GetSourceGooglePagespeedInsightsConfigurationOutput {
	return o.ApplyT(func(v LookupSourceGooglePagespeedInsightsResult) GetSourceGooglePagespeedInsightsConfiguration {
		return v.Configuration
	}).(GetSourceGooglePagespeedInsightsConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceGooglePagespeedInsightsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceGooglePagespeedInsightsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceGooglePagespeedInsightsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceGooglePagespeedInsightsResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceGooglePagespeedInsightsResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceGooglePagespeedInsightsResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceGooglePagespeedInsightsResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceGooglePagespeedInsightsResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceGooglePagespeedInsightsResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceGooglePagespeedInsightsResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceGooglePagespeedInsightsResultOutput{})
}