// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// SourceBraze DataSource
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
//			_, err := airbyte.LookupSourceBraze(ctx, &airbyte.LookupSourceBrazeArgs{
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
func LookupSourceBraze(ctx *pulumi.Context, args *LookupSourceBrazeArgs, opts ...pulumi.InvokeOption) (*LookupSourceBrazeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceBrazeResult
	err := ctx.Invoke("airbyte:index/getSourceBraze:getSourceBraze", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceBraze.
type LookupSourceBrazeArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceBraze.
type LookupSourceBrazeResult struct {
	Configuration GetSourceBrazeConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceBrazeOutput(ctx *pulumi.Context, args LookupSourceBrazeOutputArgs, opts ...pulumi.InvokeOption) LookupSourceBrazeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceBrazeResult, error) {
			args := v.(LookupSourceBrazeArgs)
			r, err := LookupSourceBraze(ctx, &args, opts...)
			var s LookupSourceBrazeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceBrazeResultOutput)
}

// A collection of arguments for invoking getSourceBraze.
type LookupSourceBrazeOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceBrazeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceBrazeArgs)(nil)).Elem()
}

// A collection of values returned by getSourceBraze.
type LookupSourceBrazeResultOutput struct{ *pulumi.OutputState }

func (LookupSourceBrazeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceBrazeResult)(nil)).Elem()
}

func (o LookupSourceBrazeResultOutput) ToLookupSourceBrazeResultOutput() LookupSourceBrazeResultOutput {
	return o
}

func (o LookupSourceBrazeResultOutput) ToLookupSourceBrazeResultOutputWithContext(ctx context.Context) LookupSourceBrazeResultOutput {
	return o
}

func (o LookupSourceBrazeResultOutput) Configuration() GetSourceBrazeConfigurationOutput {
	return o.ApplyT(func(v LookupSourceBrazeResult) GetSourceBrazeConfiguration { return v.Configuration }).(GetSourceBrazeConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceBrazeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceBrazeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceBrazeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceBrazeResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceBrazeResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceBrazeResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceBrazeResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceBrazeResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceBrazeResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceBrazeResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceBrazeResultOutput{})
}
