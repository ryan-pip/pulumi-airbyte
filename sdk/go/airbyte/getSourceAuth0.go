// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSourceAuth0(ctx *pulumi.Context, args *LookupSourceAuth0Args, opts ...pulumi.InvokeOption) (*LookupSourceAuth0Result, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSourceAuth0Result
	err := ctx.Invoke("airbyte:index/getSourceAuth0:getSourceAuth0", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceAuth0.
type LookupSourceAuth0Args struct {
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceAuth0.
type LookupSourceAuth0Result struct {
	Configuration GetSourceAuth0Configuration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	Name        string  `pulumi:"name"`
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceAuth0Output(ctx *pulumi.Context, args LookupSourceAuth0OutputArgs, opts ...pulumi.InvokeOption) LookupSourceAuth0ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceAuth0Result, error) {
			args := v.(LookupSourceAuth0Args)
			r, err := LookupSourceAuth0(ctx, &args, opts...)
			var s LookupSourceAuth0Result
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceAuth0ResultOutput)
}

// A collection of arguments for invoking getSourceAuth0.
type LookupSourceAuth0OutputArgs struct {
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceAuth0OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceAuth0Args)(nil)).Elem()
}

// A collection of values returned by getSourceAuth0.
type LookupSourceAuth0ResultOutput struct{ *pulumi.OutputState }

func (LookupSourceAuth0ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceAuth0Result)(nil)).Elem()
}

func (o LookupSourceAuth0ResultOutput) ToLookupSourceAuth0ResultOutput() LookupSourceAuth0ResultOutput {
	return o
}

func (o LookupSourceAuth0ResultOutput) ToLookupSourceAuth0ResultOutputWithContext(ctx context.Context) LookupSourceAuth0ResultOutput {
	return o
}

func (o LookupSourceAuth0ResultOutput) Configuration() GetSourceAuth0ConfigurationOutput {
	return o.ApplyT(func(v LookupSourceAuth0Result) GetSourceAuth0Configuration { return v.Configuration }).(GetSourceAuth0ConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceAuth0ResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceAuth0Result) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceAuth0ResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceAuth0Result) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceAuth0ResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceAuth0Result) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceAuth0ResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceAuth0Result) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceAuth0ResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceAuth0Result) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceAuth0ResultOutput{})
}
