// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSourceRedshift(ctx *pulumi.Context, args *LookupSourceRedshiftArgs, opts ...pulumi.InvokeOption) (*LookupSourceRedshiftResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSourceRedshiftResult
	err := ctx.Invoke("airbyte:index/getSourceRedshift:getSourceRedshift", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceRedshift.
type LookupSourceRedshiftArgs struct {
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceRedshift.
type LookupSourceRedshiftResult struct {
	Configuration GetSourceRedshiftConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	Name        string  `pulumi:"name"`
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceRedshiftOutput(ctx *pulumi.Context, args LookupSourceRedshiftOutputArgs, opts ...pulumi.InvokeOption) LookupSourceRedshiftResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceRedshiftResult, error) {
			args := v.(LookupSourceRedshiftArgs)
			r, err := LookupSourceRedshift(ctx, &args, opts...)
			var s LookupSourceRedshiftResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceRedshiftResultOutput)
}

// A collection of arguments for invoking getSourceRedshift.
type LookupSourceRedshiftOutputArgs struct {
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceRedshiftOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceRedshiftArgs)(nil)).Elem()
}

// A collection of values returned by getSourceRedshift.
type LookupSourceRedshiftResultOutput struct{ *pulumi.OutputState }

func (LookupSourceRedshiftResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceRedshiftResult)(nil)).Elem()
}

func (o LookupSourceRedshiftResultOutput) ToLookupSourceRedshiftResultOutput() LookupSourceRedshiftResultOutput {
	return o
}

func (o LookupSourceRedshiftResultOutput) ToLookupSourceRedshiftResultOutputWithContext(ctx context.Context) LookupSourceRedshiftResultOutput {
	return o
}

func (o LookupSourceRedshiftResultOutput) Configuration() GetSourceRedshiftConfigurationOutput {
	return o.ApplyT(func(v LookupSourceRedshiftResult) GetSourceRedshiftConfiguration { return v.Configuration }).(GetSourceRedshiftConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceRedshiftResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceRedshiftResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceRedshiftResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceRedshiftResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceRedshiftResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceRedshiftResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceRedshiftResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceRedshiftResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceRedshiftResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceRedshiftResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceRedshiftResultOutput{})
}
