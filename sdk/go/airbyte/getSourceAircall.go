// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSourceAircall(ctx *pulumi.Context, args *LookupSourceAircallArgs, opts ...pulumi.InvokeOption) (*LookupSourceAircallResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSourceAircallResult
	err := ctx.Invoke("airbyte:index/getSourceAircall:getSourceAircall", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceAircall.
type LookupSourceAircallArgs struct {
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceAircall.
type LookupSourceAircallResult struct {
	Configuration GetSourceAircallConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	Name        string  `pulumi:"name"`
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceAircallOutput(ctx *pulumi.Context, args LookupSourceAircallOutputArgs, opts ...pulumi.InvokeOption) LookupSourceAircallResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceAircallResult, error) {
			args := v.(LookupSourceAircallArgs)
			r, err := LookupSourceAircall(ctx, &args, opts...)
			var s LookupSourceAircallResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceAircallResultOutput)
}

// A collection of arguments for invoking getSourceAircall.
type LookupSourceAircallOutputArgs struct {
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceAircallOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceAircallArgs)(nil)).Elem()
}

// A collection of values returned by getSourceAircall.
type LookupSourceAircallResultOutput struct{ *pulumi.OutputState }

func (LookupSourceAircallResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceAircallResult)(nil)).Elem()
}

func (o LookupSourceAircallResultOutput) ToLookupSourceAircallResultOutput() LookupSourceAircallResultOutput {
	return o
}

func (o LookupSourceAircallResultOutput) ToLookupSourceAircallResultOutputWithContext(ctx context.Context) LookupSourceAircallResultOutput {
	return o
}

func (o LookupSourceAircallResultOutput) Configuration() GetSourceAircallConfigurationOutput {
	return o.ApplyT(func(v LookupSourceAircallResult) GetSourceAircallConfiguration { return v.Configuration }).(GetSourceAircallConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceAircallResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceAircallResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceAircallResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceAircallResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceAircallResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceAircallResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceAircallResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceAircallResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceAircallResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceAircallResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceAircallResultOutput{})
}
