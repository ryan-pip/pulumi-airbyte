// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSourceSurveySparrow(ctx *pulumi.Context, args *LookupSourceSurveySparrowArgs, opts ...pulumi.InvokeOption) (*LookupSourceSurveySparrowResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSourceSurveySparrowResult
	err := ctx.Invoke("airbyte:index/getSourceSurveySparrow:getSourceSurveySparrow", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceSurveySparrow.
type LookupSourceSurveySparrowArgs struct {
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceSurveySparrow.
type LookupSourceSurveySparrowResult struct {
	Configuration GetSourceSurveySparrowConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	Name        string  `pulumi:"name"`
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceSurveySparrowOutput(ctx *pulumi.Context, args LookupSourceSurveySparrowOutputArgs, opts ...pulumi.InvokeOption) LookupSourceSurveySparrowResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceSurveySparrowResult, error) {
			args := v.(LookupSourceSurveySparrowArgs)
			r, err := LookupSourceSurveySparrow(ctx, &args, opts...)
			var s LookupSourceSurveySparrowResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceSurveySparrowResultOutput)
}

// A collection of arguments for invoking getSourceSurveySparrow.
type LookupSourceSurveySparrowOutputArgs struct {
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceSurveySparrowOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceSurveySparrowArgs)(nil)).Elem()
}

// A collection of values returned by getSourceSurveySparrow.
type LookupSourceSurveySparrowResultOutput struct{ *pulumi.OutputState }

func (LookupSourceSurveySparrowResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceSurveySparrowResult)(nil)).Elem()
}

func (o LookupSourceSurveySparrowResultOutput) ToLookupSourceSurveySparrowResultOutput() LookupSourceSurveySparrowResultOutput {
	return o
}

func (o LookupSourceSurveySparrowResultOutput) ToLookupSourceSurveySparrowResultOutputWithContext(ctx context.Context) LookupSourceSurveySparrowResultOutput {
	return o
}

func (o LookupSourceSurveySparrowResultOutput) Configuration() GetSourceSurveySparrowConfigurationOutput {
	return o.ApplyT(func(v LookupSourceSurveySparrowResult) GetSourceSurveySparrowConfiguration { return v.Configuration }).(GetSourceSurveySparrowConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceSurveySparrowResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSurveySparrowResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceSurveySparrowResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSurveySparrowResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceSurveySparrowResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceSurveySparrowResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceSurveySparrowResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSurveySparrowResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceSurveySparrowResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSurveySparrowResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceSurveySparrowResultOutput{})
}
