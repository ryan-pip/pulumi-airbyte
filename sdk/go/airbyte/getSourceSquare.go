// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// SourceSquare DataSource
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
//			_, err := airbyte.LookupSourceSquare(ctx, &airbyte.LookupSourceSquareArgs{
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
func LookupSourceSquare(ctx *pulumi.Context, args *LookupSourceSquareArgs, opts ...pulumi.InvokeOption) (*LookupSourceSquareResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceSquareResult
	err := ctx.Invoke("airbyte:index/getSourceSquare:getSourceSquare", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceSquare.
type LookupSourceSquareArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceSquare.
type LookupSourceSquareResult struct {
	Configuration GetSourceSquareConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceSquareOutput(ctx *pulumi.Context, args LookupSourceSquareOutputArgs, opts ...pulumi.InvokeOption) LookupSourceSquareResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceSquareResult, error) {
			args := v.(LookupSourceSquareArgs)
			r, err := LookupSourceSquare(ctx, &args, opts...)
			var s LookupSourceSquareResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceSquareResultOutput)
}

// A collection of arguments for invoking getSourceSquare.
type LookupSourceSquareOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceSquareOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceSquareArgs)(nil)).Elem()
}

// A collection of values returned by getSourceSquare.
type LookupSourceSquareResultOutput struct{ *pulumi.OutputState }

func (LookupSourceSquareResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceSquareResult)(nil)).Elem()
}

func (o LookupSourceSquareResultOutput) ToLookupSourceSquareResultOutput() LookupSourceSquareResultOutput {
	return o
}

func (o LookupSourceSquareResultOutput) ToLookupSourceSquareResultOutputWithContext(ctx context.Context) LookupSourceSquareResultOutput {
	return o
}

func (o LookupSourceSquareResultOutput) Configuration() GetSourceSquareConfigurationOutput {
	return o.ApplyT(func(v LookupSourceSquareResult) GetSourceSquareConfiguration { return v.Configuration }).(GetSourceSquareConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceSquareResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSquareResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceSquareResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSquareResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceSquareResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceSquareResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceSquareResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSquareResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceSquareResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSquareResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceSquareResultOutput{})
}
