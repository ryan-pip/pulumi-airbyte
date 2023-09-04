// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSourceConfluence(ctx *pulumi.Context, args *LookupSourceConfluenceArgs, opts ...pulumi.InvokeOption) (*LookupSourceConfluenceResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSourceConfluenceResult
	err := ctx.Invoke("airbyte:index/getSourceConfluence:getSourceConfluence", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceConfluence.
type LookupSourceConfluenceArgs struct {
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceConfluence.
type LookupSourceConfluenceResult struct {
	Configuration GetSourceConfluenceConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	Name        string  `pulumi:"name"`
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceConfluenceOutput(ctx *pulumi.Context, args LookupSourceConfluenceOutputArgs, opts ...pulumi.InvokeOption) LookupSourceConfluenceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceConfluenceResult, error) {
			args := v.(LookupSourceConfluenceArgs)
			r, err := LookupSourceConfluence(ctx, &args, opts...)
			var s LookupSourceConfluenceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceConfluenceResultOutput)
}

// A collection of arguments for invoking getSourceConfluence.
type LookupSourceConfluenceOutputArgs struct {
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceConfluenceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceConfluenceArgs)(nil)).Elem()
}

// A collection of values returned by getSourceConfluence.
type LookupSourceConfluenceResultOutput struct{ *pulumi.OutputState }

func (LookupSourceConfluenceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceConfluenceResult)(nil)).Elem()
}

func (o LookupSourceConfluenceResultOutput) ToLookupSourceConfluenceResultOutput() LookupSourceConfluenceResultOutput {
	return o
}

func (o LookupSourceConfluenceResultOutput) ToLookupSourceConfluenceResultOutputWithContext(ctx context.Context) LookupSourceConfluenceResultOutput {
	return o
}

func (o LookupSourceConfluenceResultOutput) Configuration() GetSourceConfluenceConfigurationOutput {
	return o.ApplyT(func(v LookupSourceConfluenceResult) GetSourceConfluenceConfiguration { return v.Configuration }).(GetSourceConfluenceConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceConfluenceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceConfluenceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceConfluenceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceConfluenceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceConfluenceResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceConfluenceResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceConfluenceResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceConfluenceResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceConfluenceResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceConfluenceResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceConfluenceResultOutput{})
}
