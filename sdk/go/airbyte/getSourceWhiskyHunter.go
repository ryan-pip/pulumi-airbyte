// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSourceWhiskyHunter(ctx *pulumi.Context, args *LookupSourceWhiskyHunterArgs, opts ...pulumi.InvokeOption) (*LookupSourceWhiskyHunterResult, error) {
	var rv LookupSourceWhiskyHunterResult
	err := ctx.Invoke("airbyte:index/getSourceWhiskyHunter:getSourceWhiskyHunter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceWhiskyHunter.
type LookupSourceWhiskyHunterArgs struct {
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceWhiskyHunter.
type LookupSourceWhiskyHunterResult struct {
	Configuration GetSourceWhiskyHunterConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	Name        string  `pulumi:"name"`
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceWhiskyHunterOutput(ctx *pulumi.Context, args LookupSourceWhiskyHunterOutputArgs, opts ...pulumi.InvokeOption) LookupSourceWhiskyHunterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceWhiskyHunterResult, error) {
			args := v.(LookupSourceWhiskyHunterArgs)
			r, err := LookupSourceWhiskyHunter(ctx, &args, opts...)
			var s LookupSourceWhiskyHunterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceWhiskyHunterResultOutput)
}

// A collection of arguments for invoking getSourceWhiskyHunter.
type LookupSourceWhiskyHunterOutputArgs struct {
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceWhiskyHunterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceWhiskyHunterArgs)(nil)).Elem()
}

// A collection of values returned by getSourceWhiskyHunter.
type LookupSourceWhiskyHunterResultOutput struct{ *pulumi.OutputState }

func (LookupSourceWhiskyHunterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceWhiskyHunterResult)(nil)).Elem()
}

func (o LookupSourceWhiskyHunterResultOutput) ToLookupSourceWhiskyHunterResultOutput() LookupSourceWhiskyHunterResultOutput {
	return o
}

func (o LookupSourceWhiskyHunterResultOutput) ToLookupSourceWhiskyHunterResultOutputWithContext(ctx context.Context) LookupSourceWhiskyHunterResultOutput {
	return o
}

func (o LookupSourceWhiskyHunterResultOutput) Configuration() GetSourceWhiskyHunterConfigurationOutput {
	return o.ApplyT(func(v LookupSourceWhiskyHunterResult) GetSourceWhiskyHunterConfiguration { return v.Configuration }).(GetSourceWhiskyHunterConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceWhiskyHunterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceWhiskyHunterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceWhiskyHunterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceWhiskyHunterResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceWhiskyHunterResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceWhiskyHunterResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceWhiskyHunterResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceWhiskyHunterResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceWhiskyHunterResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceWhiskyHunterResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceWhiskyHunterResultOutput{})
}