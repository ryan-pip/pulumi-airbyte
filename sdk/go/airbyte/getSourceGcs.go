// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"internal"
)

// SourceGcs DataSource
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
// _, err := airbyte.LookupSourceGcs(ctx, %!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference), nil);
// if err != nil {
// return err
// }
// return nil
// })
// }
// ```
func LookupSourceGcs(ctx *pulumi.Context, args *LookupSourceGcsArgs, opts ...pulumi.InvokeOption) (*LookupSourceGcsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceGcsResult
	err := ctx.Invoke("airbyte:index/getSourceGcs:getSourceGcs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceGcs.
type LookupSourceGcsArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceGcs.
type LookupSourceGcsResult struct {
	Configuration GetSourceGcsConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceGcsOutput(ctx *pulumi.Context, args LookupSourceGcsOutputArgs, opts ...pulumi.InvokeOption) LookupSourceGcsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceGcsResult, error) {
			args := v.(LookupSourceGcsArgs)
			r, err := LookupSourceGcs(ctx, &args, opts...)
			var s LookupSourceGcsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceGcsResultOutput)
}

// A collection of arguments for invoking getSourceGcs.
type LookupSourceGcsOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceGcsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceGcsArgs)(nil)).Elem()
}

// A collection of values returned by getSourceGcs.
type LookupSourceGcsResultOutput struct{ *pulumi.OutputState }

func (LookupSourceGcsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceGcsResult)(nil)).Elem()
}

func (o LookupSourceGcsResultOutput) ToLookupSourceGcsResultOutput() LookupSourceGcsResultOutput {
	return o
}

func (o LookupSourceGcsResultOutput) ToLookupSourceGcsResultOutputWithContext(ctx context.Context) LookupSourceGcsResultOutput {
	return o
}

func (o LookupSourceGcsResultOutput) Configuration() GetSourceGcsConfigurationOutput {
	return o.ApplyT(func(v LookupSourceGcsResult) GetSourceGcsConfiguration { return v.Configuration }).(GetSourceGcsConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceGcsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceGcsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceGcsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceGcsResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceGcsResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceGcsResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceGcsResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceGcsResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceGcsResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceGcsResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceGcsResultOutput{})
}
