// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSourceJira(ctx *pulumi.Context, args *LookupSourceJiraArgs, opts ...pulumi.InvokeOption) (*LookupSourceJiraResult, error) {
	var rv LookupSourceJiraResult
	err := ctx.Invoke("airbyte:index/getSourceJira:getSourceJira", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceJira.
type LookupSourceJiraArgs struct {
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceJira.
type LookupSourceJiraResult struct {
	Configuration GetSourceJiraConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	Name        string  `pulumi:"name"`
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceJiraOutput(ctx *pulumi.Context, args LookupSourceJiraOutputArgs, opts ...pulumi.InvokeOption) LookupSourceJiraResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceJiraResult, error) {
			args := v.(LookupSourceJiraArgs)
			r, err := LookupSourceJira(ctx, &args, opts...)
			var s LookupSourceJiraResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceJiraResultOutput)
}

// A collection of arguments for invoking getSourceJira.
type LookupSourceJiraOutputArgs struct {
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceJiraOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceJiraArgs)(nil)).Elem()
}

// A collection of values returned by getSourceJira.
type LookupSourceJiraResultOutput struct{ *pulumi.OutputState }

func (LookupSourceJiraResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceJiraResult)(nil)).Elem()
}

func (o LookupSourceJiraResultOutput) ToLookupSourceJiraResultOutput() LookupSourceJiraResultOutput {
	return o
}

func (o LookupSourceJiraResultOutput) ToLookupSourceJiraResultOutputWithContext(ctx context.Context) LookupSourceJiraResultOutput {
	return o
}

func (o LookupSourceJiraResultOutput) Configuration() GetSourceJiraConfigurationOutput {
	return o.ApplyT(func(v LookupSourceJiraResult) GetSourceJiraConfiguration { return v.Configuration }).(GetSourceJiraConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceJiraResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceJiraResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceJiraResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceJiraResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceJiraResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceJiraResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceJiraResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceJiraResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceJiraResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceJiraResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceJiraResultOutput{})
}