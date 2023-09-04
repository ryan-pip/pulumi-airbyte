// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSourceHarvest(ctx *pulumi.Context, args *LookupSourceHarvestArgs, opts ...pulumi.InvokeOption) (*LookupSourceHarvestResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSourceHarvestResult
	err := ctx.Invoke("airbyte:index/getSourceHarvest:getSourceHarvest", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceHarvest.
type LookupSourceHarvestArgs struct {
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceHarvest.
type LookupSourceHarvestResult struct {
	Configuration GetSourceHarvestConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	Name        string  `pulumi:"name"`
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceHarvestOutput(ctx *pulumi.Context, args LookupSourceHarvestOutputArgs, opts ...pulumi.InvokeOption) LookupSourceHarvestResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceHarvestResult, error) {
			args := v.(LookupSourceHarvestArgs)
			r, err := LookupSourceHarvest(ctx, &args, opts...)
			var s LookupSourceHarvestResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceHarvestResultOutput)
}

// A collection of arguments for invoking getSourceHarvest.
type LookupSourceHarvestOutputArgs struct {
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceHarvestOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceHarvestArgs)(nil)).Elem()
}

// A collection of values returned by getSourceHarvest.
type LookupSourceHarvestResultOutput struct{ *pulumi.OutputState }

func (LookupSourceHarvestResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceHarvestResult)(nil)).Elem()
}

func (o LookupSourceHarvestResultOutput) ToLookupSourceHarvestResultOutput() LookupSourceHarvestResultOutput {
	return o
}

func (o LookupSourceHarvestResultOutput) ToLookupSourceHarvestResultOutputWithContext(ctx context.Context) LookupSourceHarvestResultOutput {
	return o
}

func (o LookupSourceHarvestResultOutput) Configuration() GetSourceHarvestConfigurationOutput {
	return o.ApplyT(func(v LookupSourceHarvestResult) GetSourceHarvestConfiguration { return v.Configuration }).(GetSourceHarvestConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceHarvestResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceHarvestResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceHarvestResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceHarvestResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceHarvestResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceHarvestResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceHarvestResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceHarvestResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceHarvestResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceHarvestResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceHarvestResultOutput{})
}
