// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSourceZendeskTalk(ctx *pulumi.Context, args *LookupSourceZendeskTalkArgs, opts ...pulumi.InvokeOption) (*LookupSourceZendeskTalkResult, error) {
	var rv LookupSourceZendeskTalkResult
	err := ctx.Invoke("airbyte:index/getSourceZendeskTalk:getSourceZendeskTalk", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceZendeskTalk.
type LookupSourceZendeskTalkArgs struct {
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceZendeskTalk.
type LookupSourceZendeskTalkResult struct {
	Configuration GetSourceZendeskTalkConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	Name        string  `pulumi:"name"`
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceZendeskTalkOutput(ctx *pulumi.Context, args LookupSourceZendeskTalkOutputArgs, opts ...pulumi.InvokeOption) LookupSourceZendeskTalkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceZendeskTalkResult, error) {
			args := v.(LookupSourceZendeskTalkArgs)
			r, err := LookupSourceZendeskTalk(ctx, &args, opts...)
			var s LookupSourceZendeskTalkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceZendeskTalkResultOutput)
}

// A collection of arguments for invoking getSourceZendeskTalk.
type LookupSourceZendeskTalkOutputArgs struct {
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceZendeskTalkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceZendeskTalkArgs)(nil)).Elem()
}

// A collection of values returned by getSourceZendeskTalk.
type LookupSourceZendeskTalkResultOutput struct{ *pulumi.OutputState }

func (LookupSourceZendeskTalkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceZendeskTalkResult)(nil)).Elem()
}

func (o LookupSourceZendeskTalkResultOutput) ToLookupSourceZendeskTalkResultOutput() LookupSourceZendeskTalkResultOutput {
	return o
}

func (o LookupSourceZendeskTalkResultOutput) ToLookupSourceZendeskTalkResultOutputWithContext(ctx context.Context) LookupSourceZendeskTalkResultOutput {
	return o
}

func (o LookupSourceZendeskTalkResultOutput) Configuration() GetSourceZendeskTalkConfigurationOutput {
	return o.ApplyT(func(v LookupSourceZendeskTalkResult) GetSourceZendeskTalkConfiguration { return v.Configuration }).(GetSourceZendeskTalkConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceZendeskTalkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceZendeskTalkResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceZendeskTalkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceZendeskTalkResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceZendeskTalkResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceZendeskTalkResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceZendeskTalkResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceZendeskTalkResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceZendeskTalkResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceZendeskTalkResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceZendeskTalkResultOutput{})
}