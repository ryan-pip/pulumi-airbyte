// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSourceTvmazeSchedule(ctx *pulumi.Context, args *LookupSourceTvmazeScheduleArgs, opts ...pulumi.InvokeOption) (*LookupSourceTvmazeScheduleResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSourceTvmazeScheduleResult
	err := ctx.Invoke("airbyte:index/getSourceTvmazeSchedule:getSourceTvmazeSchedule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceTvmazeSchedule.
type LookupSourceTvmazeScheduleArgs struct {
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceTvmazeSchedule.
type LookupSourceTvmazeScheduleResult struct {
	Configuration GetSourceTvmazeScheduleConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	Name        string  `pulumi:"name"`
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceTvmazeScheduleOutput(ctx *pulumi.Context, args LookupSourceTvmazeScheduleOutputArgs, opts ...pulumi.InvokeOption) LookupSourceTvmazeScheduleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceTvmazeScheduleResult, error) {
			args := v.(LookupSourceTvmazeScheduleArgs)
			r, err := LookupSourceTvmazeSchedule(ctx, &args, opts...)
			var s LookupSourceTvmazeScheduleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceTvmazeScheduleResultOutput)
}

// A collection of arguments for invoking getSourceTvmazeSchedule.
type LookupSourceTvmazeScheduleOutputArgs struct {
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceTvmazeScheduleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceTvmazeScheduleArgs)(nil)).Elem()
}

// A collection of values returned by getSourceTvmazeSchedule.
type LookupSourceTvmazeScheduleResultOutput struct{ *pulumi.OutputState }

func (LookupSourceTvmazeScheduleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceTvmazeScheduleResult)(nil)).Elem()
}

func (o LookupSourceTvmazeScheduleResultOutput) ToLookupSourceTvmazeScheduleResultOutput() LookupSourceTvmazeScheduleResultOutput {
	return o
}

func (o LookupSourceTvmazeScheduleResultOutput) ToLookupSourceTvmazeScheduleResultOutputWithContext(ctx context.Context) LookupSourceTvmazeScheduleResultOutput {
	return o
}

func (o LookupSourceTvmazeScheduleResultOutput) Configuration() GetSourceTvmazeScheduleConfigurationOutput {
	return o.ApplyT(func(v LookupSourceTvmazeScheduleResult) GetSourceTvmazeScheduleConfiguration { return v.Configuration }).(GetSourceTvmazeScheduleConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceTvmazeScheduleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceTvmazeScheduleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceTvmazeScheduleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceTvmazeScheduleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceTvmazeScheduleResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceTvmazeScheduleResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceTvmazeScheduleResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceTvmazeScheduleResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceTvmazeScheduleResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceTvmazeScheduleResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceTvmazeScheduleResultOutput{})
}
