// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSourceTiktokMarketing(ctx *pulumi.Context, args *LookupSourceTiktokMarketingArgs, opts ...pulumi.InvokeOption) (*LookupSourceTiktokMarketingResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSourceTiktokMarketingResult
	err := ctx.Invoke("airbyte:index/getSourceTiktokMarketing:getSourceTiktokMarketing", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceTiktokMarketing.
type LookupSourceTiktokMarketingArgs struct {
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceTiktokMarketing.
type LookupSourceTiktokMarketingResult struct {
	Configuration GetSourceTiktokMarketingConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	Name        string  `pulumi:"name"`
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceTiktokMarketingOutput(ctx *pulumi.Context, args LookupSourceTiktokMarketingOutputArgs, opts ...pulumi.InvokeOption) LookupSourceTiktokMarketingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceTiktokMarketingResult, error) {
			args := v.(LookupSourceTiktokMarketingArgs)
			r, err := LookupSourceTiktokMarketing(ctx, &args, opts...)
			var s LookupSourceTiktokMarketingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceTiktokMarketingResultOutput)
}

// A collection of arguments for invoking getSourceTiktokMarketing.
type LookupSourceTiktokMarketingOutputArgs struct {
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceTiktokMarketingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceTiktokMarketingArgs)(nil)).Elem()
}

// A collection of values returned by getSourceTiktokMarketing.
type LookupSourceTiktokMarketingResultOutput struct{ *pulumi.OutputState }

func (LookupSourceTiktokMarketingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceTiktokMarketingResult)(nil)).Elem()
}

func (o LookupSourceTiktokMarketingResultOutput) ToLookupSourceTiktokMarketingResultOutput() LookupSourceTiktokMarketingResultOutput {
	return o
}

func (o LookupSourceTiktokMarketingResultOutput) ToLookupSourceTiktokMarketingResultOutputWithContext(ctx context.Context) LookupSourceTiktokMarketingResultOutput {
	return o
}

func (o LookupSourceTiktokMarketingResultOutput) Configuration() GetSourceTiktokMarketingConfigurationOutput {
	return o.ApplyT(func(v LookupSourceTiktokMarketingResult) GetSourceTiktokMarketingConfiguration {
		return v.Configuration
	}).(GetSourceTiktokMarketingConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceTiktokMarketingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceTiktokMarketingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceTiktokMarketingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceTiktokMarketingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceTiktokMarketingResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceTiktokMarketingResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceTiktokMarketingResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceTiktokMarketingResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceTiktokMarketingResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceTiktokMarketingResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceTiktokMarketingResultOutput{})
}
