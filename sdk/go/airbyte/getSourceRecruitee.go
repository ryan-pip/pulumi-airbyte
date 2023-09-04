// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSourceRecruitee(ctx *pulumi.Context, args *LookupSourceRecruiteeArgs, opts ...pulumi.InvokeOption) (*LookupSourceRecruiteeResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSourceRecruiteeResult
	err := ctx.Invoke("airbyte:index/getSourceRecruitee:getSourceRecruitee", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceRecruitee.
type LookupSourceRecruiteeArgs struct {
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceRecruitee.
type LookupSourceRecruiteeResult struct {
	Configuration GetSourceRecruiteeConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	Name        string  `pulumi:"name"`
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceRecruiteeOutput(ctx *pulumi.Context, args LookupSourceRecruiteeOutputArgs, opts ...pulumi.InvokeOption) LookupSourceRecruiteeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceRecruiteeResult, error) {
			args := v.(LookupSourceRecruiteeArgs)
			r, err := LookupSourceRecruitee(ctx, &args, opts...)
			var s LookupSourceRecruiteeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceRecruiteeResultOutput)
}

// A collection of arguments for invoking getSourceRecruitee.
type LookupSourceRecruiteeOutputArgs struct {
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceRecruiteeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceRecruiteeArgs)(nil)).Elem()
}

// A collection of values returned by getSourceRecruitee.
type LookupSourceRecruiteeResultOutput struct{ *pulumi.OutputState }

func (LookupSourceRecruiteeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceRecruiteeResult)(nil)).Elem()
}

func (o LookupSourceRecruiteeResultOutput) ToLookupSourceRecruiteeResultOutput() LookupSourceRecruiteeResultOutput {
	return o
}

func (o LookupSourceRecruiteeResultOutput) ToLookupSourceRecruiteeResultOutputWithContext(ctx context.Context) LookupSourceRecruiteeResultOutput {
	return o
}

func (o LookupSourceRecruiteeResultOutput) Configuration() GetSourceRecruiteeConfigurationOutput {
	return o.ApplyT(func(v LookupSourceRecruiteeResult) GetSourceRecruiteeConfiguration { return v.Configuration }).(GetSourceRecruiteeConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceRecruiteeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceRecruiteeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceRecruiteeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceRecruiteeResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceRecruiteeResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceRecruiteeResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceRecruiteeResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceRecruiteeResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceRecruiteeResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceRecruiteeResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceRecruiteeResultOutput{})
}
