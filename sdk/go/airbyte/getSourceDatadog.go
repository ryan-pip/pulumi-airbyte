// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"internal"
)

// SourceDatadog DataSource
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
// _, err := airbyte.LookupSourceDatadog(ctx, %!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference), nil);
// if err != nil {
// return err
// }
// return nil
// })
// }
// ```
func LookupSourceDatadog(ctx *pulumi.Context, args *LookupSourceDatadogArgs, opts ...pulumi.InvokeOption) (*LookupSourceDatadogResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceDatadogResult
	err := ctx.Invoke("airbyte:index/getSourceDatadog:getSourceDatadog", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceDatadog.
type LookupSourceDatadogArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceDatadog.
type LookupSourceDatadogResult struct {
	Configuration GetSourceDatadogConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceDatadogOutput(ctx *pulumi.Context, args LookupSourceDatadogOutputArgs, opts ...pulumi.InvokeOption) LookupSourceDatadogResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceDatadogResult, error) {
			args := v.(LookupSourceDatadogArgs)
			r, err := LookupSourceDatadog(ctx, &args, opts...)
			var s LookupSourceDatadogResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceDatadogResultOutput)
}

// A collection of arguments for invoking getSourceDatadog.
type LookupSourceDatadogOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceDatadogOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceDatadogArgs)(nil)).Elem()
}

// A collection of values returned by getSourceDatadog.
type LookupSourceDatadogResultOutput struct{ *pulumi.OutputState }

func (LookupSourceDatadogResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceDatadogResult)(nil)).Elem()
}

func (o LookupSourceDatadogResultOutput) ToLookupSourceDatadogResultOutput() LookupSourceDatadogResultOutput {
	return o
}

func (o LookupSourceDatadogResultOutput) ToLookupSourceDatadogResultOutputWithContext(ctx context.Context) LookupSourceDatadogResultOutput {
	return o
}

func (o LookupSourceDatadogResultOutput) Configuration() GetSourceDatadogConfigurationOutput {
	return o.ApplyT(func(v LookupSourceDatadogResult) GetSourceDatadogConfiguration { return v.Configuration }).(GetSourceDatadogConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceDatadogResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceDatadogResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceDatadogResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceDatadogResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceDatadogResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceDatadogResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceDatadogResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceDatadogResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceDatadogResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceDatadogResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceDatadogResultOutput{})
}
