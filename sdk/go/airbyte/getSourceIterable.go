// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSourceIterable(ctx *pulumi.Context, args *LookupSourceIterableArgs, opts ...pulumi.InvokeOption) (*LookupSourceIterableResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSourceIterableResult
	err := ctx.Invoke("airbyte:index/getSourceIterable:getSourceIterable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceIterable.
type LookupSourceIterableArgs struct {
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceIterable.
type LookupSourceIterableResult struct {
	Configuration GetSourceIterableConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	Name        string  `pulumi:"name"`
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceIterableOutput(ctx *pulumi.Context, args LookupSourceIterableOutputArgs, opts ...pulumi.InvokeOption) LookupSourceIterableResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceIterableResult, error) {
			args := v.(LookupSourceIterableArgs)
			r, err := LookupSourceIterable(ctx, &args, opts...)
			var s LookupSourceIterableResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceIterableResultOutput)
}

// A collection of arguments for invoking getSourceIterable.
type LookupSourceIterableOutputArgs struct {
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceIterableOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceIterableArgs)(nil)).Elem()
}

// A collection of values returned by getSourceIterable.
type LookupSourceIterableResultOutput struct{ *pulumi.OutputState }

func (LookupSourceIterableResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceIterableResult)(nil)).Elem()
}

func (o LookupSourceIterableResultOutput) ToLookupSourceIterableResultOutput() LookupSourceIterableResultOutput {
	return o
}

func (o LookupSourceIterableResultOutput) ToLookupSourceIterableResultOutputWithContext(ctx context.Context) LookupSourceIterableResultOutput {
	return o
}

func (o LookupSourceIterableResultOutput) Configuration() GetSourceIterableConfigurationOutput {
	return o.ApplyT(func(v LookupSourceIterableResult) GetSourceIterableConfiguration { return v.Configuration }).(GetSourceIterableConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceIterableResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceIterableResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceIterableResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceIterableResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceIterableResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceIterableResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceIterableResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceIterableResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceIterableResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceIterableResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceIterableResultOutput{})
}
