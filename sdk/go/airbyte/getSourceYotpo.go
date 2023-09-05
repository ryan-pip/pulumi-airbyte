// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"internal"
)

// SourceYotpo DataSource
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
// _, err := airbyte.LookupSourceYotpo(ctx, %!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference), nil);
// if err != nil {
// return err
// }
// return nil
// })
// }
// ```
func LookupSourceYotpo(ctx *pulumi.Context, args *LookupSourceYotpoArgs, opts ...pulumi.InvokeOption) (*LookupSourceYotpoResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceYotpoResult
	err := ctx.Invoke("airbyte:index/getSourceYotpo:getSourceYotpo", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceYotpo.
type LookupSourceYotpoArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceYotpo.
type LookupSourceYotpoResult struct {
	Configuration GetSourceYotpoConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceYotpoOutput(ctx *pulumi.Context, args LookupSourceYotpoOutputArgs, opts ...pulumi.InvokeOption) LookupSourceYotpoResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceYotpoResult, error) {
			args := v.(LookupSourceYotpoArgs)
			r, err := LookupSourceYotpo(ctx, &args, opts...)
			var s LookupSourceYotpoResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceYotpoResultOutput)
}

// A collection of arguments for invoking getSourceYotpo.
type LookupSourceYotpoOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceYotpoOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceYotpoArgs)(nil)).Elem()
}

// A collection of values returned by getSourceYotpo.
type LookupSourceYotpoResultOutput struct{ *pulumi.OutputState }

func (LookupSourceYotpoResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceYotpoResult)(nil)).Elem()
}

func (o LookupSourceYotpoResultOutput) ToLookupSourceYotpoResultOutput() LookupSourceYotpoResultOutput {
	return o
}

func (o LookupSourceYotpoResultOutput) ToLookupSourceYotpoResultOutputWithContext(ctx context.Context) LookupSourceYotpoResultOutput {
	return o
}

func (o LookupSourceYotpoResultOutput) Configuration() GetSourceYotpoConfigurationOutput {
	return o.ApplyT(func(v LookupSourceYotpoResult) GetSourceYotpoConfiguration { return v.Configuration }).(GetSourceYotpoConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceYotpoResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceYotpoResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceYotpoResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceYotpoResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceYotpoResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceYotpoResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceYotpoResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceYotpoResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceYotpoResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceYotpoResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceYotpoResultOutput{})
}
