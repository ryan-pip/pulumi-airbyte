// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// SourceMonday DataSource
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
//			_, err := airbyte.LookupSourceMonday(ctx, &airbyte.LookupSourceMondayArgs{
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
func LookupSourceMonday(ctx *pulumi.Context, args *LookupSourceMondayArgs, opts ...pulumi.InvokeOption) (*LookupSourceMondayResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceMondayResult
	err := ctx.Invoke("airbyte:index/getSourceMonday:getSourceMonday", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceMonday.
type LookupSourceMondayArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceMonday.
type LookupSourceMondayResult struct {
	Configuration GetSourceMondayConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceMondayOutput(ctx *pulumi.Context, args LookupSourceMondayOutputArgs, opts ...pulumi.InvokeOption) LookupSourceMondayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceMondayResult, error) {
			args := v.(LookupSourceMondayArgs)
			r, err := LookupSourceMonday(ctx, &args, opts...)
			var s LookupSourceMondayResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceMondayResultOutput)
}

// A collection of arguments for invoking getSourceMonday.
type LookupSourceMondayOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceMondayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceMondayArgs)(nil)).Elem()
}

// A collection of values returned by getSourceMonday.
type LookupSourceMondayResultOutput struct{ *pulumi.OutputState }

func (LookupSourceMondayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceMondayResult)(nil)).Elem()
}

func (o LookupSourceMondayResultOutput) ToLookupSourceMondayResultOutput() LookupSourceMondayResultOutput {
	return o
}

func (o LookupSourceMondayResultOutput) ToLookupSourceMondayResultOutputWithContext(ctx context.Context) LookupSourceMondayResultOutput {
	return o
}

func (o LookupSourceMondayResultOutput) Configuration() GetSourceMondayConfigurationOutput {
	return o.ApplyT(func(v LookupSourceMondayResult) GetSourceMondayConfiguration { return v.Configuration }).(GetSourceMondayConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceMondayResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceMondayResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceMondayResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceMondayResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceMondayResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceMondayResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceMondayResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceMondayResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceMondayResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceMondayResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceMondayResultOutput{})
}
