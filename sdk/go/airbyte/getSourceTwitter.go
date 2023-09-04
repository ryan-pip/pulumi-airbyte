// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSourceTwitter(ctx *pulumi.Context, args *LookupSourceTwitterArgs, opts ...pulumi.InvokeOption) (*LookupSourceTwitterResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSourceTwitterResult
	err := ctx.Invoke("airbyte:index/getSourceTwitter:getSourceTwitter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceTwitter.
type LookupSourceTwitterArgs struct {
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceTwitter.
type LookupSourceTwitterResult struct {
	Configuration GetSourceTwitterConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	Name        string  `pulumi:"name"`
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceTwitterOutput(ctx *pulumi.Context, args LookupSourceTwitterOutputArgs, opts ...pulumi.InvokeOption) LookupSourceTwitterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceTwitterResult, error) {
			args := v.(LookupSourceTwitterArgs)
			r, err := LookupSourceTwitter(ctx, &args, opts...)
			var s LookupSourceTwitterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceTwitterResultOutput)
}

// A collection of arguments for invoking getSourceTwitter.
type LookupSourceTwitterOutputArgs struct {
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceTwitterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceTwitterArgs)(nil)).Elem()
}

// A collection of values returned by getSourceTwitter.
type LookupSourceTwitterResultOutput struct{ *pulumi.OutputState }

func (LookupSourceTwitterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceTwitterResult)(nil)).Elem()
}

func (o LookupSourceTwitterResultOutput) ToLookupSourceTwitterResultOutput() LookupSourceTwitterResultOutput {
	return o
}

func (o LookupSourceTwitterResultOutput) ToLookupSourceTwitterResultOutputWithContext(ctx context.Context) LookupSourceTwitterResultOutput {
	return o
}

func (o LookupSourceTwitterResultOutput) Configuration() GetSourceTwitterConfigurationOutput {
	return o.ApplyT(func(v LookupSourceTwitterResult) GetSourceTwitterConfiguration { return v.Configuration }).(GetSourceTwitterConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceTwitterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceTwitterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceTwitterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceTwitterResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceTwitterResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceTwitterResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceTwitterResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceTwitterResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceTwitterResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceTwitterResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceTwitterResultOutput{})
}
