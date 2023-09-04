// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSourceSftp(ctx *pulumi.Context, args *LookupSourceSftpArgs, opts ...pulumi.InvokeOption) (*LookupSourceSftpResult, error) {
	var rv LookupSourceSftpResult
	err := ctx.Invoke("airbyte:index/getSourceSftp:getSourceSftp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceSftp.
type LookupSourceSftpArgs struct {
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceSftp.
type LookupSourceSftpResult struct {
	Configuration GetSourceSftpConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	Name        string  `pulumi:"name"`
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceSftpOutput(ctx *pulumi.Context, args LookupSourceSftpOutputArgs, opts ...pulumi.InvokeOption) LookupSourceSftpResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceSftpResult, error) {
			args := v.(LookupSourceSftpArgs)
			r, err := LookupSourceSftp(ctx, &args, opts...)
			var s LookupSourceSftpResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceSftpResultOutput)
}

// A collection of arguments for invoking getSourceSftp.
type LookupSourceSftpOutputArgs struct {
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceSftpOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceSftpArgs)(nil)).Elem()
}

// A collection of values returned by getSourceSftp.
type LookupSourceSftpResultOutput struct{ *pulumi.OutputState }

func (LookupSourceSftpResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceSftpResult)(nil)).Elem()
}

func (o LookupSourceSftpResultOutput) ToLookupSourceSftpResultOutput() LookupSourceSftpResultOutput {
	return o
}

func (o LookupSourceSftpResultOutput) ToLookupSourceSftpResultOutputWithContext(ctx context.Context) LookupSourceSftpResultOutput {
	return o
}

func (o LookupSourceSftpResultOutput) Configuration() GetSourceSftpConfigurationOutput {
	return o.ApplyT(func(v LookupSourceSftpResult) GetSourceSftpConfiguration { return v.Configuration }).(GetSourceSftpConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceSftpResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSftpResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceSftpResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSftpResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceSftpResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceSftpResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceSftpResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSftpResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceSftpResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSftpResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceSftpResultOutput{})
}