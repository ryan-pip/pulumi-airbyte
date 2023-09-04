// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSourceInstatus(ctx *pulumi.Context, args *LookupSourceInstatusArgs, opts ...pulumi.InvokeOption) (*LookupSourceInstatusResult, error) {
	var rv LookupSourceInstatusResult
	err := ctx.Invoke("airbyte:index/getSourceInstatus:getSourceInstatus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceInstatus.
type LookupSourceInstatusArgs struct {
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceInstatus.
type LookupSourceInstatusResult struct {
	Configuration GetSourceInstatusConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	Name        string  `pulumi:"name"`
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceInstatusOutput(ctx *pulumi.Context, args LookupSourceInstatusOutputArgs, opts ...pulumi.InvokeOption) LookupSourceInstatusResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceInstatusResult, error) {
			args := v.(LookupSourceInstatusArgs)
			r, err := LookupSourceInstatus(ctx, &args, opts...)
			var s LookupSourceInstatusResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceInstatusResultOutput)
}

// A collection of arguments for invoking getSourceInstatus.
type LookupSourceInstatusOutputArgs struct {
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceInstatusOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceInstatusArgs)(nil)).Elem()
}

// A collection of values returned by getSourceInstatus.
type LookupSourceInstatusResultOutput struct{ *pulumi.OutputState }

func (LookupSourceInstatusResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceInstatusResult)(nil)).Elem()
}

func (o LookupSourceInstatusResultOutput) ToLookupSourceInstatusResultOutput() LookupSourceInstatusResultOutput {
	return o
}

func (o LookupSourceInstatusResultOutput) ToLookupSourceInstatusResultOutputWithContext(ctx context.Context) LookupSourceInstatusResultOutput {
	return o
}

func (o LookupSourceInstatusResultOutput) Configuration() GetSourceInstatusConfigurationOutput {
	return o.ApplyT(func(v LookupSourceInstatusResult) GetSourceInstatusConfiguration { return v.Configuration }).(GetSourceInstatusConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceInstatusResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceInstatusResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceInstatusResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceInstatusResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceInstatusResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceInstatusResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceInstatusResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceInstatusResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceInstatusResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceInstatusResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceInstatusResultOutput{})
}