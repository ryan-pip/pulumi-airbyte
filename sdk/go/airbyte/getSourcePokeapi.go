// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSourcePokeapi(ctx *pulumi.Context, args *LookupSourcePokeapiArgs, opts ...pulumi.InvokeOption) (*LookupSourcePokeapiResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSourcePokeapiResult
	err := ctx.Invoke("airbyte:index/getSourcePokeapi:getSourcePokeapi", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourcePokeapi.
type LookupSourcePokeapiArgs struct {
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourcePokeapi.
type LookupSourcePokeapiResult struct {
	Configuration GetSourcePokeapiConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	Name        string  `pulumi:"name"`
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourcePokeapiOutput(ctx *pulumi.Context, args LookupSourcePokeapiOutputArgs, opts ...pulumi.InvokeOption) LookupSourcePokeapiResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourcePokeapiResult, error) {
			args := v.(LookupSourcePokeapiArgs)
			r, err := LookupSourcePokeapi(ctx, &args, opts...)
			var s LookupSourcePokeapiResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourcePokeapiResultOutput)
}

// A collection of arguments for invoking getSourcePokeapi.
type LookupSourcePokeapiOutputArgs struct {
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourcePokeapiOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourcePokeapiArgs)(nil)).Elem()
}

// A collection of values returned by getSourcePokeapi.
type LookupSourcePokeapiResultOutput struct{ *pulumi.OutputState }

func (LookupSourcePokeapiResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourcePokeapiResult)(nil)).Elem()
}

func (o LookupSourcePokeapiResultOutput) ToLookupSourcePokeapiResultOutput() LookupSourcePokeapiResultOutput {
	return o
}

func (o LookupSourcePokeapiResultOutput) ToLookupSourcePokeapiResultOutputWithContext(ctx context.Context) LookupSourcePokeapiResultOutput {
	return o
}

func (o LookupSourcePokeapiResultOutput) Configuration() GetSourcePokeapiConfigurationOutput {
	return o.ApplyT(func(v LookupSourcePokeapiResult) GetSourcePokeapiConfiguration { return v.Configuration }).(GetSourcePokeapiConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourcePokeapiResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourcePokeapiResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourcePokeapiResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourcePokeapiResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourcePokeapiResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourcePokeapiResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourcePokeapiResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourcePokeapiResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourcePokeapiResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourcePokeapiResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourcePokeapiResultOutput{})
}
