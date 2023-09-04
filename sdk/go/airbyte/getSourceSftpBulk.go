// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSourceSftpBulk(ctx *pulumi.Context, args *LookupSourceSftpBulkArgs, opts ...pulumi.InvokeOption) (*LookupSourceSftpBulkResult, error) {
	var rv LookupSourceSftpBulkResult
	err := ctx.Invoke("airbyte:index/getSourceSftpBulk:getSourceSftpBulk", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceSftpBulk.
type LookupSourceSftpBulkArgs struct {
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceSftpBulk.
type LookupSourceSftpBulkResult struct {
	Configuration GetSourceSftpBulkConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	Name        string  `pulumi:"name"`
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceSftpBulkOutput(ctx *pulumi.Context, args LookupSourceSftpBulkOutputArgs, opts ...pulumi.InvokeOption) LookupSourceSftpBulkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceSftpBulkResult, error) {
			args := v.(LookupSourceSftpBulkArgs)
			r, err := LookupSourceSftpBulk(ctx, &args, opts...)
			var s LookupSourceSftpBulkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceSftpBulkResultOutput)
}

// A collection of arguments for invoking getSourceSftpBulk.
type LookupSourceSftpBulkOutputArgs struct {
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceSftpBulkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceSftpBulkArgs)(nil)).Elem()
}

// A collection of values returned by getSourceSftpBulk.
type LookupSourceSftpBulkResultOutput struct{ *pulumi.OutputState }

func (LookupSourceSftpBulkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceSftpBulkResult)(nil)).Elem()
}

func (o LookupSourceSftpBulkResultOutput) ToLookupSourceSftpBulkResultOutput() LookupSourceSftpBulkResultOutput {
	return o
}

func (o LookupSourceSftpBulkResultOutput) ToLookupSourceSftpBulkResultOutputWithContext(ctx context.Context) LookupSourceSftpBulkResultOutput {
	return o
}

func (o LookupSourceSftpBulkResultOutput) Configuration() GetSourceSftpBulkConfigurationOutput {
	return o.ApplyT(func(v LookupSourceSftpBulkResult) GetSourceSftpBulkConfiguration { return v.Configuration }).(GetSourceSftpBulkConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceSftpBulkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSftpBulkResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceSftpBulkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSftpBulkResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceSftpBulkResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceSftpBulkResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceSftpBulkResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSftpBulkResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceSftpBulkResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSftpBulkResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceSftpBulkResultOutput{})
}