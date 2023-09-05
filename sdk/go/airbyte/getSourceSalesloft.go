// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"internal"
)

// SourceSalesloft DataSource
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-airbyte/sdk/go/airbyte"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
// func main() {
// pulumi.Run(func(ctx *pulumi.Context) error {
// _, err := airbyte.LookupSourceSalesloft(ctx, %!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference), nil);
// if err != nil {
// return err
// }
// return nil
// })
// }
// ```
func LookupSourceSalesloft(ctx *pulumi.Context, args *LookupSourceSalesloftArgs, opts ...pulumi.InvokeOption) (*LookupSourceSalesloftResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceSalesloftResult
	err := ctx.Invoke("airbyte:index/getSourceSalesloft:getSourceSalesloft", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceSalesloft.
type LookupSourceSalesloftArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceSalesloft.
type LookupSourceSalesloftResult struct {
	Configuration GetSourceSalesloftConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceSalesloftOutput(ctx *pulumi.Context, args LookupSourceSalesloftOutputArgs, opts ...pulumi.InvokeOption) LookupSourceSalesloftResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceSalesloftResult, error) {
			args := v.(LookupSourceSalesloftArgs)
			r, err := LookupSourceSalesloft(ctx, &args, opts...)
			var s LookupSourceSalesloftResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceSalesloftResultOutput)
}

// A collection of arguments for invoking getSourceSalesloft.
type LookupSourceSalesloftOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceSalesloftOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceSalesloftArgs)(nil)).Elem()
}

// A collection of values returned by getSourceSalesloft.
type LookupSourceSalesloftResultOutput struct{ *pulumi.OutputState }

func (LookupSourceSalesloftResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceSalesloftResult)(nil)).Elem()
}

func (o LookupSourceSalesloftResultOutput) ToLookupSourceSalesloftResultOutput() LookupSourceSalesloftResultOutput {
	return o
}

func (o LookupSourceSalesloftResultOutput) ToLookupSourceSalesloftResultOutputWithContext(ctx context.Context) LookupSourceSalesloftResultOutput {
	return o
}

func (o LookupSourceSalesloftResultOutput) Configuration() GetSourceSalesloftConfigurationOutput {
	return o.ApplyT(func(v LookupSourceSalesloftResult) GetSourceSalesloftConfiguration { return v.Configuration }).(GetSourceSalesloftConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceSalesloftResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSalesloftResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceSalesloftResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSalesloftResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceSalesloftResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceSalesloftResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceSalesloftResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSalesloftResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceSalesloftResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSalesloftResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceSalesloftResultOutput{})
}
