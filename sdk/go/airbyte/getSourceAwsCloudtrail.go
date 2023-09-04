// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSourceAwsCloudtrail(ctx *pulumi.Context, args *LookupSourceAwsCloudtrailArgs, opts ...pulumi.InvokeOption) (*LookupSourceAwsCloudtrailResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSourceAwsCloudtrailResult
	err := ctx.Invoke("airbyte:index/getSourceAwsCloudtrail:getSourceAwsCloudtrail", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceAwsCloudtrail.
type LookupSourceAwsCloudtrailArgs struct {
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceAwsCloudtrail.
type LookupSourceAwsCloudtrailResult struct {
	Configuration GetSourceAwsCloudtrailConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	Name        string  `pulumi:"name"`
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceAwsCloudtrailOutput(ctx *pulumi.Context, args LookupSourceAwsCloudtrailOutputArgs, opts ...pulumi.InvokeOption) LookupSourceAwsCloudtrailResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceAwsCloudtrailResult, error) {
			args := v.(LookupSourceAwsCloudtrailArgs)
			r, err := LookupSourceAwsCloudtrail(ctx, &args, opts...)
			var s LookupSourceAwsCloudtrailResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceAwsCloudtrailResultOutput)
}

// A collection of arguments for invoking getSourceAwsCloudtrail.
type LookupSourceAwsCloudtrailOutputArgs struct {
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceAwsCloudtrailOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceAwsCloudtrailArgs)(nil)).Elem()
}

// A collection of values returned by getSourceAwsCloudtrail.
type LookupSourceAwsCloudtrailResultOutput struct{ *pulumi.OutputState }

func (LookupSourceAwsCloudtrailResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceAwsCloudtrailResult)(nil)).Elem()
}

func (o LookupSourceAwsCloudtrailResultOutput) ToLookupSourceAwsCloudtrailResultOutput() LookupSourceAwsCloudtrailResultOutput {
	return o
}

func (o LookupSourceAwsCloudtrailResultOutput) ToLookupSourceAwsCloudtrailResultOutputWithContext(ctx context.Context) LookupSourceAwsCloudtrailResultOutput {
	return o
}

func (o LookupSourceAwsCloudtrailResultOutput) Configuration() GetSourceAwsCloudtrailConfigurationOutput {
	return o.ApplyT(func(v LookupSourceAwsCloudtrailResult) GetSourceAwsCloudtrailConfiguration { return v.Configuration }).(GetSourceAwsCloudtrailConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceAwsCloudtrailResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceAwsCloudtrailResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceAwsCloudtrailResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceAwsCloudtrailResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceAwsCloudtrailResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceAwsCloudtrailResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceAwsCloudtrailResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceAwsCloudtrailResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceAwsCloudtrailResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceAwsCloudtrailResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceAwsCloudtrailResultOutput{})
}
