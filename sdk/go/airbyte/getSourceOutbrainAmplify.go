// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"internal"
)

// SourceOutbrainAmplify DataSource
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
// _, err := airbyte.LookupSourceOutbrainAmplify(ctx, %!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference), nil);
// if err != nil {
// return err
// }
// return nil
// })
// }
// ```
func LookupSourceOutbrainAmplify(ctx *pulumi.Context, args *LookupSourceOutbrainAmplifyArgs, opts ...pulumi.InvokeOption) (*LookupSourceOutbrainAmplifyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceOutbrainAmplifyResult
	err := ctx.Invoke("airbyte:index/getSourceOutbrainAmplify:getSourceOutbrainAmplify", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceOutbrainAmplify.
type LookupSourceOutbrainAmplifyArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceOutbrainAmplify.
type LookupSourceOutbrainAmplifyResult struct {
	Configuration GetSourceOutbrainAmplifyConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceOutbrainAmplifyOutput(ctx *pulumi.Context, args LookupSourceOutbrainAmplifyOutputArgs, opts ...pulumi.InvokeOption) LookupSourceOutbrainAmplifyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceOutbrainAmplifyResult, error) {
			args := v.(LookupSourceOutbrainAmplifyArgs)
			r, err := LookupSourceOutbrainAmplify(ctx, &args, opts...)
			var s LookupSourceOutbrainAmplifyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceOutbrainAmplifyResultOutput)
}

// A collection of arguments for invoking getSourceOutbrainAmplify.
type LookupSourceOutbrainAmplifyOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceOutbrainAmplifyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceOutbrainAmplifyArgs)(nil)).Elem()
}

// A collection of values returned by getSourceOutbrainAmplify.
type LookupSourceOutbrainAmplifyResultOutput struct{ *pulumi.OutputState }

func (LookupSourceOutbrainAmplifyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceOutbrainAmplifyResult)(nil)).Elem()
}

func (o LookupSourceOutbrainAmplifyResultOutput) ToLookupSourceOutbrainAmplifyResultOutput() LookupSourceOutbrainAmplifyResultOutput {
	return o
}

func (o LookupSourceOutbrainAmplifyResultOutput) ToLookupSourceOutbrainAmplifyResultOutputWithContext(ctx context.Context) LookupSourceOutbrainAmplifyResultOutput {
	return o
}

func (o LookupSourceOutbrainAmplifyResultOutput) Configuration() GetSourceOutbrainAmplifyConfigurationOutput {
	return o.ApplyT(func(v LookupSourceOutbrainAmplifyResult) GetSourceOutbrainAmplifyConfiguration {
		return v.Configuration
	}).(GetSourceOutbrainAmplifyConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceOutbrainAmplifyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceOutbrainAmplifyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceOutbrainAmplifyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceOutbrainAmplifyResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceOutbrainAmplifyResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceOutbrainAmplifyResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceOutbrainAmplifyResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceOutbrainAmplifyResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceOutbrainAmplifyResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceOutbrainAmplifyResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceOutbrainAmplifyResultOutput{})
}
