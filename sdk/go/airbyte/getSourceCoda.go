// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// SourceCoda DataSource
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
//			_, err := airbyte.LookupSourceCoda(ctx, &airbyte.LookupSourceCodaArgs{
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
func LookupSourceCoda(ctx *pulumi.Context, args *LookupSourceCodaArgs, opts ...pulumi.InvokeOption) (*LookupSourceCodaResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceCodaResult
	err := ctx.Invoke("airbyte:index/getSourceCoda:getSourceCoda", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceCoda.
type LookupSourceCodaArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceCoda.
type LookupSourceCodaResult struct {
	Configuration GetSourceCodaConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceCodaOutput(ctx *pulumi.Context, args LookupSourceCodaOutputArgs, opts ...pulumi.InvokeOption) LookupSourceCodaResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceCodaResult, error) {
			args := v.(LookupSourceCodaArgs)
			r, err := LookupSourceCoda(ctx, &args, opts...)
			var s LookupSourceCodaResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceCodaResultOutput)
}

// A collection of arguments for invoking getSourceCoda.
type LookupSourceCodaOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceCodaOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceCodaArgs)(nil)).Elem()
}

// A collection of values returned by getSourceCoda.
type LookupSourceCodaResultOutput struct{ *pulumi.OutputState }

func (LookupSourceCodaResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceCodaResult)(nil)).Elem()
}

func (o LookupSourceCodaResultOutput) ToLookupSourceCodaResultOutput() LookupSourceCodaResultOutput {
	return o
}

func (o LookupSourceCodaResultOutput) ToLookupSourceCodaResultOutputWithContext(ctx context.Context) LookupSourceCodaResultOutput {
	return o
}

func (o LookupSourceCodaResultOutput) Configuration() GetSourceCodaConfigurationOutput {
	return o.ApplyT(func(v LookupSourceCodaResult) GetSourceCodaConfiguration { return v.Configuration }).(GetSourceCodaConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceCodaResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceCodaResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceCodaResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceCodaResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceCodaResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceCodaResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceCodaResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceCodaResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceCodaResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceCodaResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceCodaResultOutput{})
}
