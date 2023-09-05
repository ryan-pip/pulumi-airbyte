// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// SourceGetlago DataSource
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
//			_, err := airbyte.LookupSourceGetlago(ctx, &airbyte.LookupSourceGetlagoArgs{
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
func LookupSourceGetlago(ctx *pulumi.Context, args *LookupSourceGetlagoArgs, opts ...pulumi.InvokeOption) (*LookupSourceGetlagoResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceGetlagoResult
	err := ctx.Invoke("airbyte:index/getSourceGetlago:getSourceGetlago", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceGetlago.
type LookupSourceGetlagoArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceGetlago.
type LookupSourceGetlagoResult struct {
	Configuration GetSourceGetlagoConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceGetlagoOutput(ctx *pulumi.Context, args LookupSourceGetlagoOutputArgs, opts ...pulumi.InvokeOption) LookupSourceGetlagoResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceGetlagoResult, error) {
			args := v.(LookupSourceGetlagoArgs)
			r, err := LookupSourceGetlago(ctx, &args, opts...)
			var s LookupSourceGetlagoResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceGetlagoResultOutput)
}

// A collection of arguments for invoking getSourceGetlago.
type LookupSourceGetlagoOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceGetlagoOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceGetlagoArgs)(nil)).Elem()
}

// A collection of values returned by getSourceGetlago.
type LookupSourceGetlagoResultOutput struct{ *pulumi.OutputState }

func (LookupSourceGetlagoResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceGetlagoResult)(nil)).Elem()
}

func (o LookupSourceGetlagoResultOutput) ToLookupSourceGetlagoResultOutput() LookupSourceGetlagoResultOutput {
	return o
}

func (o LookupSourceGetlagoResultOutput) ToLookupSourceGetlagoResultOutputWithContext(ctx context.Context) LookupSourceGetlagoResultOutput {
	return o
}

func (o LookupSourceGetlagoResultOutput) Configuration() GetSourceGetlagoConfigurationOutput {
	return o.ApplyT(func(v LookupSourceGetlagoResult) GetSourceGetlagoConfiguration { return v.Configuration }).(GetSourceGetlagoConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceGetlagoResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceGetlagoResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceGetlagoResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceGetlagoResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceGetlagoResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceGetlagoResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceGetlagoResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceGetlagoResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceGetlagoResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceGetlagoResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceGetlagoResultOutput{})
}
