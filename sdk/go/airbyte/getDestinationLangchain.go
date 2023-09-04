// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDestinationLangchain(ctx *pulumi.Context, args *LookupDestinationLangchainArgs, opts ...pulumi.InvokeOption) (*LookupDestinationLangchainResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupDestinationLangchainResult
	err := ctx.Invoke("airbyte:index/getDestinationLangchain:getDestinationLangchain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDestinationLangchain.
type LookupDestinationLangchainArgs struct {
	DestinationId string `pulumi:"destinationId"`
}

// A collection of values returned by getDestinationLangchain.
type LookupDestinationLangchainResult struct {
	Configuration GetDestinationLangchainConfiguration `pulumi:"configuration"`
	DestinationId string                               `pulumi:"destinationId"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupDestinationLangchainOutput(ctx *pulumi.Context, args LookupDestinationLangchainOutputArgs, opts ...pulumi.InvokeOption) LookupDestinationLangchainResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDestinationLangchainResult, error) {
			args := v.(LookupDestinationLangchainArgs)
			r, err := LookupDestinationLangchain(ctx, &args, opts...)
			var s LookupDestinationLangchainResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDestinationLangchainResultOutput)
}

// A collection of arguments for invoking getDestinationLangchain.
type LookupDestinationLangchainOutputArgs struct {
	DestinationId pulumi.StringInput `pulumi:"destinationId"`
}

func (LookupDestinationLangchainOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDestinationLangchainArgs)(nil)).Elem()
}

// A collection of values returned by getDestinationLangchain.
type LookupDestinationLangchainResultOutput struct{ *pulumi.OutputState }

func (LookupDestinationLangchainResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDestinationLangchainResult)(nil)).Elem()
}

func (o LookupDestinationLangchainResultOutput) ToLookupDestinationLangchainResultOutput() LookupDestinationLangchainResultOutput {
	return o
}

func (o LookupDestinationLangchainResultOutput) ToLookupDestinationLangchainResultOutputWithContext(ctx context.Context) LookupDestinationLangchainResultOutput {
	return o
}

func (o LookupDestinationLangchainResultOutput) Configuration() GetDestinationLangchainConfigurationOutput {
	return o.ApplyT(func(v LookupDestinationLangchainResult) GetDestinationLangchainConfiguration { return v.Configuration }).(GetDestinationLangchainConfigurationOutput)
}

func (o LookupDestinationLangchainResultOutput) DestinationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationLangchainResult) string { return v.DestinationId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDestinationLangchainResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationLangchainResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDestinationLangchainResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationLangchainResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDestinationLangchainResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationLangchainResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDestinationLangchainResultOutput{})
}
