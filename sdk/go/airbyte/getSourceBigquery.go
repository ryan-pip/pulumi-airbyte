// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSourceBigquery(ctx *pulumi.Context, args *LookupSourceBigqueryArgs, opts ...pulumi.InvokeOption) (*LookupSourceBigqueryResult, error) {
	var rv LookupSourceBigqueryResult
	err := ctx.Invoke("airbyte:index/getSourceBigquery:getSourceBigquery", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceBigquery.
type LookupSourceBigqueryArgs struct {
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceBigquery.
type LookupSourceBigqueryResult struct {
	Configuration GetSourceBigqueryConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	Name        string  `pulumi:"name"`
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceBigqueryOutput(ctx *pulumi.Context, args LookupSourceBigqueryOutputArgs, opts ...pulumi.InvokeOption) LookupSourceBigqueryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceBigqueryResult, error) {
			args := v.(LookupSourceBigqueryArgs)
			r, err := LookupSourceBigquery(ctx, &args, opts...)
			var s LookupSourceBigqueryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceBigqueryResultOutput)
}

// A collection of arguments for invoking getSourceBigquery.
type LookupSourceBigqueryOutputArgs struct {
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceBigqueryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceBigqueryArgs)(nil)).Elem()
}

// A collection of values returned by getSourceBigquery.
type LookupSourceBigqueryResultOutput struct{ *pulumi.OutputState }

func (LookupSourceBigqueryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceBigqueryResult)(nil)).Elem()
}

func (o LookupSourceBigqueryResultOutput) ToLookupSourceBigqueryResultOutput() LookupSourceBigqueryResultOutput {
	return o
}

func (o LookupSourceBigqueryResultOutput) ToLookupSourceBigqueryResultOutputWithContext(ctx context.Context) LookupSourceBigqueryResultOutput {
	return o
}

func (o LookupSourceBigqueryResultOutput) Configuration() GetSourceBigqueryConfigurationOutput {
	return o.ApplyT(func(v LookupSourceBigqueryResult) GetSourceBigqueryConfiguration { return v.Configuration }).(GetSourceBigqueryConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceBigqueryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceBigqueryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceBigqueryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceBigqueryResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceBigqueryResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceBigqueryResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceBigqueryResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceBigqueryResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceBigqueryResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceBigqueryResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceBigqueryResultOutput{})
}