// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSourceLinkedinPages(ctx *pulumi.Context, args *LookupSourceLinkedinPagesArgs, opts ...pulumi.InvokeOption) (*LookupSourceLinkedinPagesResult, error) {
	var rv LookupSourceLinkedinPagesResult
	err := ctx.Invoke("airbyte:index/getSourceLinkedinPages:getSourceLinkedinPages", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceLinkedinPages.
type LookupSourceLinkedinPagesArgs struct {
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceLinkedinPages.
type LookupSourceLinkedinPagesResult struct {
	Configuration GetSourceLinkedinPagesConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	Name        string  `pulumi:"name"`
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceLinkedinPagesOutput(ctx *pulumi.Context, args LookupSourceLinkedinPagesOutputArgs, opts ...pulumi.InvokeOption) LookupSourceLinkedinPagesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceLinkedinPagesResult, error) {
			args := v.(LookupSourceLinkedinPagesArgs)
			r, err := LookupSourceLinkedinPages(ctx, &args, opts...)
			var s LookupSourceLinkedinPagesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceLinkedinPagesResultOutput)
}

// A collection of arguments for invoking getSourceLinkedinPages.
type LookupSourceLinkedinPagesOutputArgs struct {
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceLinkedinPagesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceLinkedinPagesArgs)(nil)).Elem()
}

// A collection of values returned by getSourceLinkedinPages.
type LookupSourceLinkedinPagesResultOutput struct{ *pulumi.OutputState }

func (LookupSourceLinkedinPagesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceLinkedinPagesResult)(nil)).Elem()
}

func (o LookupSourceLinkedinPagesResultOutput) ToLookupSourceLinkedinPagesResultOutput() LookupSourceLinkedinPagesResultOutput {
	return o
}

func (o LookupSourceLinkedinPagesResultOutput) ToLookupSourceLinkedinPagesResultOutputWithContext(ctx context.Context) LookupSourceLinkedinPagesResultOutput {
	return o
}

func (o LookupSourceLinkedinPagesResultOutput) Configuration() GetSourceLinkedinPagesConfigurationOutput {
	return o.ApplyT(func(v LookupSourceLinkedinPagesResult) GetSourceLinkedinPagesConfiguration { return v.Configuration }).(GetSourceLinkedinPagesConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceLinkedinPagesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceLinkedinPagesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceLinkedinPagesResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceLinkedinPagesResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceLinkedinPagesResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceLinkedinPagesResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceLinkedinPagesResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceLinkedinPagesResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceLinkedinPagesResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceLinkedinPagesResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceLinkedinPagesResultOutput{})
}