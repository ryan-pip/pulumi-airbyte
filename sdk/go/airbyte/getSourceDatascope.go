// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"internal"
)

// SourceDatascope DataSource
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
// _, err := airbyte.LookupSourceDatascope(ctx, %!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference), nil);
// if err != nil {
// return err
// }
// return nil
// })
// }
// ```
func LookupSourceDatascope(ctx *pulumi.Context, args *LookupSourceDatascopeArgs, opts ...pulumi.InvokeOption) (*LookupSourceDatascopeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceDatascopeResult
	err := ctx.Invoke("airbyte:index/getSourceDatascope:getSourceDatascope", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceDatascope.
type LookupSourceDatascopeArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceDatascope.
type LookupSourceDatascopeResult struct {
	Configuration GetSourceDatascopeConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceDatascopeOutput(ctx *pulumi.Context, args LookupSourceDatascopeOutputArgs, opts ...pulumi.InvokeOption) LookupSourceDatascopeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceDatascopeResult, error) {
			args := v.(LookupSourceDatascopeArgs)
			r, err := LookupSourceDatascope(ctx, &args, opts...)
			var s LookupSourceDatascopeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceDatascopeResultOutput)
}

// A collection of arguments for invoking getSourceDatascope.
type LookupSourceDatascopeOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceDatascopeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceDatascopeArgs)(nil)).Elem()
}

// A collection of values returned by getSourceDatascope.
type LookupSourceDatascopeResultOutput struct{ *pulumi.OutputState }

func (LookupSourceDatascopeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceDatascopeResult)(nil)).Elem()
}

func (o LookupSourceDatascopeResultOutput) ToLookupSourceDatascopeResultOutput() LookupSourceDatascopeResultOutput {
	return o
}

func (o LookupSourceDatascopeResultOutput) ToLookupSourceDatascopeResultOutputWithContext(ctx context.Context) LookupSourceDatascopeResultOutput {
	return o
}

func (o LookupSourceDatascopeResultOutput) Configuration() GetSourceDatascopeConfigurationOutput {
	return o.ApplyT(func(v LookupSourceDatascopeResult) GetSourceDatascopeConfiguration { return v.Configuration }).(GetSourceDatascopeConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceDatascopeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceDatascopeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceDatascopeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceDatascopeResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceDatascopeResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceDatascopeResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceDatascopeResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceDatascopeResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceDatascopeResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceDatascopeResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceDatascopeResultOutput{})
}
