// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// SourceDremio DataSource
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := airbyte.LookupSourceDremio(ctx, &airbyte.LookupSourceDremioArgs{
//				SecretId: pulumi.StringRef("...my_secret_id..."),
//				SourceId: "...my_source_id...",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupSourceDremio(ctx *pulumi.Context, args *LookupSourceDremioArgs, opts ...pulumi.InvokeOption) (*LookupSourceDremioResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceDremioResult
	err := ctx.Invoke("airbyte:index/getSourceDremio:getSourceDremio", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceDremio.
type LookupSourceDremioArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceDremio.
type LookupSourceDremioResult struct {
	Configuration GetSourceDremioConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceDremioOutput(ctx *pulumi.Context, args LookupSourceDremioOutputArgs, opts ...pulumi.InvokeOption) LookupSourceDremioResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceDremioResult, error) {
			args := v.(LookupSourceDremioArgs)
			r, err := LookupSourceDremio(ctx, &args, opts...)
			var s LookupSourceDremioResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceDremioResultOutput)
}

// A collection of arguments for invoking getSourceDremio.
type LookupSourceDremioOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceDremioOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceDremioArgs)(nil)).Elem()
}

// A collection of values returned by getSourceDremio.
type LookupSourceDremioResultOutput struct{ *pulumi.OutputState }

func (LookupSourceDremioResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceDremioResult)(nil)).Elem()
}

func (o LookupSourceDremioResultOutput) ToLookupSourceDremioResultOutput() LookupSourceDremioResultOutput {
	return o
}

func (o LookupSourceDremioResultOutput) ToLookupSourceDremioResultOutputWithContext(ctx context.Context) LookupSourceDremioResultOutput {
	return o
}

func (o LookupSourceDremioResultOutput) Configuration() GetSourceDremioConfigurationOutput {
	return o.ApplyT(func(v LookupSourceDremioResult) GetSourceDremioConfiguration { return v.Configuration }).(GetSourceDremioConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceDremioResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceDremioResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceDremioResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceDremioResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceDremioResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceDremioResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceDremioResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceDremioResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceDremioResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceDremioResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceDremioResultOutput{})
}
