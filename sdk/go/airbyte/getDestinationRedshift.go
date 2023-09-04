// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDestinationRedshift(ctx *pulumi.Context, args *LookupDestinationRedshiftArgs, opts ...pulumi.InvokeOption) (*LookupDestinationRedshiftResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupDestinationRedshiftResult
	err := ctx.Invoke("airbyte:index/getDestinationRedshift:getDestinationRedshift", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDestinationRedshift.
type LookupDestinationRedshiftArgs struct {
	DestinationId string `pulumi:"destinationId"`
}

// A collection of values returned by getDestinationRedshift.
type LookupDestinationRedshiftResult struct {
	Configuration GetDestinationRedshiftConfiguration `pulumi:"configuration"`
	DestinationId string                              `pulumi:"destinationId"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupDestinationRedshiftOutput(ctx *pulumi.Context, args LookupDestinationRedshiftOutputArgs, opts ...pulumi.InvokeOption) LookupDestinationRedshiftResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDestinationRedshiftResult, error) {
			args := v.(LookupDestinationRedshiftArgs)
			r, err := LookupDestinationRedshift(ctx, &args, opts...)
			var s LookupDestinationRedshiftResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDestinationRedshiftResultOutput)
}

// A collection of arguments for invoking getDestinationRedshift.
type LookupDestinationRedshiftOutputArgs struct {
	DestinationId pulumi.StringInput `pulumi:"destinationId"`
}

func (LookupDestinationRedshiftOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDestinationRedshiftArgs)(nil)).Elem()
}

// A collection of values returned by getDestinationRedshift.
type LookupDestinationRedshiftResultOutput struct{ *pulumi.OutputState }

func (LookupDestinationRedshiftResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDestinationRedshiftResult)(nil)).Elem()
}

func (o LookupDestinationRedshiftResultOutput) ToLookupDestinationRedshiftResultOutput() LookupDestinationRedshiftResultOutput {
	return o
}

func (o LookupDestinationRedshiftResultOutput) ToLookupDestinationRedshiftResultOutputWithContext(ctx context.Context) LookupDestinationRedshiftResultOutput {
	return o
}

func (o LookupDestinationRedshiftResultOutput) Configuration() GetDestinationRedshiftConfigurationOutput {
	return o.ApplyT(func(v LookupDestinationRedshiftResult) GetDestinationRedshiftConfiguration { return v.Configuration }).(GetDestinationRedshiftConfigurationOutput)
}

func (o LookupDestinationRedshiftResultOutput) DestinationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationRedshiftResult) string { return v.DestinationId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDestinationRedshiftResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationRedshiftResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDestinationRedshiftResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationRedshiftResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDestinationRedshiftResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationRedshiftResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDestinationRedshiftResultOutput{})
}
