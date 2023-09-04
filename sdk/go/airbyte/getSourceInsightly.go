// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSourceInsightly(ctx *pulumi.Context, args *LookupSourceInsightlyArgs, opts ...pulumi.InvokeOption) (*LookupSourceInsightlyResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSourceInsightlyResult
	err := ctx.Invoke("airbyte:index/getSourceInsightly:getSourceInsightly", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceInsightly.
type LookupSourceInsightlyArgs struct {
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceInsightly.
type LookupSourceInsightlyResult struct {
	Configuration GetSourceInsightlyConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	Name        string  `pulumi:"name"`
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceInsightlyOutput(ctx *pulumi.Context, args LookupSourceInsightlyOutputArgs, opts ...pulumi.InvokeOption) LookupSourceInsightlyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceInsightlyResult, error) {
			args := v.(LookupSourceInsightlyArgs)
			r, err := LookupSourceInsightly(ctx, &args, opts...)
			var s LookupSourceInsightlyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceInsightlyResultOutput)
}

// A collection of arguments for invoking getSourceInsightly.
type LookupSourceInsightlyOutputArgs struct {
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceInsightlyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceInsightlyArgs)(nil)).Elem()
}

// A collection of values returned by getSourceInsightly.
type LookupSourceInsightlyResultOutput struct{ *pulumi.OutputState }

func (LookupSourceInsightlyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceInsightlyResult)(nil)).Elem()
}

func (o LookupSourceInsightlyResultOutput) ToLookupSourceInsightlyResultOutput() LookupSourceInsightlyResultOutput {
	return o
}

func (o LookupSourceInsightlyResultOutput) ToLookupSourceInsightlyResultOutputWithContext(ctx context.Context) LookupSourceInsightlyResultOutput {
	return o
}

func (o LookupSourceInsightlyResultOutput) Configuration() GetSourceInsightlyConfigurationOutput {
	return o.ApplyT(func(v LookupSourceInsightlyResult) GetSourceInsightlyConfiguration { return v.Configuration }).(GetSourceInsightlyConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceInsightlyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceInsightlyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceInsightlyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceInsightlyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceInsightlyResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceInsightlyResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceInsightlyResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceInsightlyResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceInsightlyResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceInsightlyResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceInsightlyResultOutput{})
}
