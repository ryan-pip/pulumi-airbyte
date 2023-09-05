// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"internal"
)

// SourceGreenhouse DataSource
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
// _, err := airbyte.LookupSourceGreenhouse(ctx, %!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference), nil);
// if err != nil {
// return err
// }
// return nil
// })
// }
// ```
func LookupSourceGreenhouse(ctx *pulumi.Context, args *LookupSourceGreenhouseArgs, opts ...pulumi.InvokeOption) (*LookupSourceGreenhouseResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceGreenhouseResult
	err := ctx.Invoke("airbyte:index/getSourceGreenhouse:getSourceGreenhouse", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceGreenhouse.
type LookupSourceGreenhouseArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceGreenhouse.
type LookupSourceGreenhouseResult struct {
	Configuration GetSourceGreenhouseConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceGreenhouseOutput(ctx *pulumi.Context, args LookupSourceGreenhouseOutputArgs, opts ...pulumi.InvokeOption) LookupSourceGreenhouseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceGreenhouseResult, error) {
			args := v.(LookupSourceGreenhouseArgs)
			r, err := LookupSourceGreenhouse(ctx, &args, opts...)
			var s LookupSourceGreenhouseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceGreenhouseResultOutput)
}

// A collection of arguments for invoking getSourceGreenhouse.
type LookupSourceGreenhouseOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceGreenhouseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceGreenhouseArgs)(nil)).Elem()
}

// A collection of values returned by getSourceGreenhouse.
type LookupSourceGreenhouseResultOutput struct{ *pulumi.OutputState }

func (LookupSourceGreenhouseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceGreenhouseResult)(nil)).Elem()
}

func (o LookupSourceGreenhouseResultOutput) ToLookupSourceGreenhouseResultOutput() LookupSourceGreenhouseResultOutput {
	return o
}

func (o LookupSourceGreenhouseResultOutput) ToLookupSourceGreenhouseResultOutputWithContext(ctx context.Context) LookupSourceGreenhouseResultOutput {
	return o
}

func (o LookupSourceGreenhouseResultOutput) Configuration() GetSourceGreenhouseConfigurationOutput {
	return o.ApplyT(func(v LookupSourceGreenhouseResult) GetSourceGreenhouseConfiguration { return v.Configuration }).(GetSourceGreenhouseConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceGreenhouseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceGreenhouseResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceGreenhouseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceGreenhouseResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceGreenhouseResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceGreenhouseResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceGreenhouseResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceGreenhouseResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceGreenhouseResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceGreenhouseResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceGreenhouseResultOutput{})
}
