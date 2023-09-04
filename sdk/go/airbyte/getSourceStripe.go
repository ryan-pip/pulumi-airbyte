// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSourceStripe(ctx *pulumi.Context, args *LookupSourceStripeArgs, opts ...pulumi.InvokeOption) (*LookupSourceStripeResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSourceStripeResult
	err := ctx.Invoke("airbyte:index/getSourceStripe:getSourceStripe", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceStripe.
type LookupSourceStripeArgs struct {
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceStripe.
type LookupSourceStripeResult struct {
	Configuration GetSourceStripeConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	Name        string  `pulumi:"name"`
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceStripeOutput(ctx *pulumi.Context, args LookupSourceStripeOutputArgs, opts ...pulumi.InvokeOption) LookupSourceStripeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceStripeResult, error) {
			args := v.(LookupSourceStripeArgs)
			r, err := LookupSourceStripe(ctx, &args, opts...)
			var s LookupSourceStripeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceStripeResultOutput)
}

// A collection of arguments for invoking getSourceStripe.
type LookupSourceStripeOutputArgs struct {
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceStripeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceStripeArgs)(nil)).Elem()
}

// A collection of values returned by getSourceStripe.
type LookupSourceStripeResultOutput struct{ *pulumi.OutputState }

func (LookupSourceStripeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceStripeResult)(nil)).Elem()
}

func (o LookupSourceStripeResultOutput) ToLookupSourceStripeResultOutput() LookupSourceStripeResultOutput {
	return o
}

func (o LookupSourceStripeResultOutput) ToLookupSourceStripeResultOutputWithContext(ctx context.Context) LookupSourceStripeResultOutput {
	return o
}

func (o LookupSourceStripeResultOutput) Configuration() GetSourceStripeConfigurationOutput {
	return o.ApplyT(func(v LookupSourceStripeResult) GetSourceStripeConfiguration { return v.Configuration }).(GetSourceStripeConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceStripeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceStripeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceStripeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceStripeResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceStripeResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceStripeResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceStripeResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceStripeResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceStripeResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceStripeResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceStripeResultOutput{})
}
