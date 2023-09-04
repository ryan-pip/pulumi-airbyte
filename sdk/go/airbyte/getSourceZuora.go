// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSourceZuora(ctx *pulumi.Context, args *LookupSourceZuoraArgs, opts ...pulumi.InvokeOption) (*LookupSourceZuoraResult, error) {
	var rv LookupSourceZuoraResult
	err := ctx.Invoke("airbyte:index/getSourceZuora:getSourceZuora", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceZuora.
type LookupSourceZuoraArgs struct {
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceZuora.
type LookupSourceZuoraResult struct {
	Configuration GetSourceZuoraConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	Name        string  `pulumi:"name"`
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceZuoraOutput(ctx *pulumi.Context, args LookupSourceZuoraOutputArgs, opts ...pulumi.InvokeOption) LookupSourceZuoraResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceZuoraResult, error) {
			args := v.(LookupSourceZuoraArgs)
			r, err := LookupSourceZuora(ctx, &args, opts...)
			var s LookupSourceZuoraResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceZuoraResultOutput)
}

// A collection of arguments for invoking getSourceZuora.
type LookupSourceZuoraOutputArgs struct {
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceZuoraOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceZuoraArgs)(nil)).Elem()
}

// A collection of values returned by getSourceZuora.
type LookupSourceZuoraResultOutput struct{ *pulumi.OutputState }

func (LookupSourceZuoraResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceZuoraResult)(nil)).Elem()
}

func (o LookupSourceZuoraResultOutput) ToLookupSourceZuoraResultOutput() LookupSourceZuoraResultOutput {
	return o
}

func (o LookupSourceZuoraResultOutput) ToLookupSourceZuoraResultOutputWithContext(ctx context.Context) LookupSourceZuoraResultOutput {
	return o
}

func (o LookupSourceZuoraResultOutput) Configuration() GetSourceZuoraConfigurationOutput {
	return o.ApplyT(func(v LookupSourceZuoraResult) GetSourceZuoraConfiguration { return v.Configuration }).(GetSourceZuoraConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceZuoraResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceZuoraResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceZuoraResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceZuoraResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceZuoraResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceZuoraResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceZuoraResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceZuoraResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceZuoraResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceZuoraResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceZuoraResultOutput{})
}