// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSourceFirebolt(ctx *pulumi.Context, args *LookupSourceFireboltArgs, opts ...pulumi.InvokeOption) (*LookupSourceFireboltResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSourceFireboltResult
	err := ctx.Invoke("airbyte:index/getSourceFirebolt:getSourceFirebolt", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceFirebolt.
type LookupSourceFireboltArgs struct {
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceFirebolt.
type LookupSourceFireboltResult struct {
	Configuration GetSourceFireboltConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	Name        string  `pulumi:"name"`
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceFireboltOutput(ctx *pulumi.Context, args LookupSourceFireboltOutputArgs, opts ...pulumi.InvokeOption) LookupSourceFireboltResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceFireboltResult, error) {
			args := v.(LookupSourceFireboltArgs)
			r, err := LookupSourceFirebolt(ctx, &args, opts...)
			var s LookupSourceFireboltResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceFireboltResultOutput)
}

// A collection of arguments for invoking getSourceFirebolt.
type LookupSourceFireboltOutputArgs struct {
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceFireboltOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceFireboltArgs)(nil)).Elem()
}

// A collection of values returned by getSourceFirebolt.
type LookupSourceFireboltResultOutput struct{ *pulumi.OutputState }

func (LookupSourceFireboltResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceFireboltResult)(nil)).Elem()
}

func (o LookupSourceFireboltResultOutput) ToLookupSourceFireboltResultOutput() LookupSourceFireboltResultOutput {
	return o
}

func (o LookupSourceFireboltResultOutput) ToLookupSourceFireboltResultOutputWithContext(ctx context.Context) LookupSourceFireboltResultOutput {
	return o
}

func (o LookupSourceFireboltResultOutput) Configuration() GetSourceFireboltConfigurationOutput {
	return o.ApplyT(func(v LookupSourceFireboltResult) GetSourceFireboltConfiguration { return v.Configuration }).(GetSourceFireboltConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceFireboltResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceFireboltResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceFireboltResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceFireboltResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceFireboltResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceFireboltResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceFireboltResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceFireboltResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceFireboltResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceFireboltResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceFireboltResultOutput{})
}
