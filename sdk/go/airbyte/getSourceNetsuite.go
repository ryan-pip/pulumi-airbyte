// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSourceNetsuite(ctx *pulumi.Context, args *LookupSourceNetsuiteArgs, opts ...pulumi.InvokeOption) (*LookupSourceNetsuiteResult, error) {
	var rv LookupSourceNetsuiteResult
	err := ctx.Invoke("airbyte:index/getSourceNetsuite:getSourceNetsuite", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceNetsuite.
type LookupSourceNetsuiteArgs struct {
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceNetsuite.
type LookupSourceNetsuiteResult struct {
	Configuration GetSourceNetsuiteConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	Name        string  `pulumi:"name"`
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceNetsuiteOutput(ctx *pulumi.Context, args LookupSourceNetsuiteOutputArgs, opts ...pulumi.InvokeOption) LookupSourceNetsuiteResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceNetsuiteResult, error) {
			args := v.(LookupSourceNetsuiteArgs)
			r, err := LookupSourceNetsuite(ctx, &args, opts...)
			var s LookupSourceNetsuiteResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceNetsuiteResultOutput)
}

// A collection of arguments for invoking getSourceNetsuite.
type LookupSourceNetsuiteOutputArgs struct {
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceNetsuiteOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceNetsuiteArgs)(nil)).Elem()
}

// A collection of values returned by getSourceNetsuite.
type LookupSourceNetsuiteResultOutput struct{ *pulumi.OutputState }

func (LookupSourceNetsuiteResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceNetsuiteResult)(nil)).Elem()
}

func (o LookupSourceNetsuiteResultOutput) ToLookupSourceNetsuiteResultOutput() LookupSourceNetsuiteResultOutput {
	return o
}

func (o LookupSourceNetsuiteResultOutput) ToLookupSourceNetsuiteResultOutputWithContext(ctx context.Context) LookupSourceNetsuiteResultOutput {
	return o
}

func (o LookupSourceNetsuiteResultOutput) Configuration() GetSourceNetsuiteConfigurationOutput {
	return o.ApplyT(func(v LookupSourceNetsuiteResult) GetSourceNetsuiteConfiguration { return v.Configuration }).(GetSourceNetsuiteConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceNetsuiteResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceNetsuiteResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceNetsuiteResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceNetsuiteResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceNetsuiteResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceNetsuiteResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceNetsuiteResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceNetsuiteResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceNetsuiteResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceNetsuiteResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceNetsuiteResultOutput{})
}