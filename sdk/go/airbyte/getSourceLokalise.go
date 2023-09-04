// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSourceLokalise(ctx *pulumi.Context, args *LookupSourceLokaliseArgs, opts ...pulumi.InvokeOption) (*LookupSourceLokaliseResult, error) {
	var rv LookupSourceLokaliseResult
	err := ctx.Invoke("airbyte:index/getSourceLokalise:getSourceLokalise", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceLokalise.
type LookupSourceLokaliseArgs struct {
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceLokalise.
type LookupSourceLokaliseResult struct {
	Configuration GetSourceLokaliseConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	Name        string  `pulumi:"name"`
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceLokaliseOutput(ctx *pulumi.Context, args LookupSourceLokaliseOutputArgs, opts ...pulumi.InvokeOption) LookupSourceLokaliseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceLokaliseResult, error) {
			args := v.(LookupSourceLokaliseArgs)
			r, err := LookupSourceLokalise(ctx, &args, opts...)
			var s LookupSourceLokaliseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceLokaliseResultOutput)
}

// A collection of arguments for invoking getSourceLokalise.
type LookupSourceLokaliseOutputArgs struct {
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceLokaliseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceLokaliseArgs)(nil)).Elem()
}

// A collection of values returned by getSourceLokalise.
type LookupSourceLokaliseResultOutput struct{ *pulumi.OutputState }

func (LookupSourceLokaliseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceLokaliseResult)(nil)).Elem()
}

func (o LookupSourceLokaliseResultOutput) ToLookupSourceLokaliseResultOutput() LookupSourceLokaliseResultOutput {
	return o
}

func (o LookupSourceLokaliseResultOutput) ToLookupSourceLokaliseResultOutputWithContext(ctx context.Context) LookupSourceLokaliseResultOutput {
	return o
}

func (o LookupSourceLokaliseResultOutput) Configuration() GetSourceLokaliseConfigurationOutput {
	return o.ApplyT(func(v LookupSourceLokaliseResult) GetSourceLokaliseConfiguration { return v.Configuration }).(GetSourceLokaliseConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceLokaliseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceLokaliseResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceLokaliseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceLokaliseResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceLokaliseResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceLokaliseResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceLokaliseResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceLokaliseResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceLokaliseResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceLokaliseResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceLokaliseResultOutput{})
}