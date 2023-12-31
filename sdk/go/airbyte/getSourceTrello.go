// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// SourceTrello DataSource
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
//			_, err := airbyte.LookupSourceTrello(ctx, &airbyte.LookupSourceTrelloArgs{
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
func LookupSourceTrello(ctx *pulumi.Context, args *LookupSourceTrelloArgs, opts ...pulumi.InvokeOption) (*LookupSourceTrelloResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceTrelloResult
	err := ctx.Invoke("airbyte:index/getSourceTrello:getSourceTrello", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceTrello.
type LookupSourceTrelloArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceTrello.
type LookupSourceTrelloResult struct {
	Configuration GetSourceTrelloConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceTrelloOutput(ctx *pulumi.Context, args LookupSourceTrelloOutputArgs, opts ...pulumi.InvokeOption) LookupSourceTrelloResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceTrelloResult, error) {
			args := v.(LookupSourceTrelloArgs)
			r, err := LookupSourceTrello(ctx, &args, opts...)
			var s LookupSourceTrelloResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceTrelloResultOutput)
}

// A collection of arguments for invoking getSourceTrello.
type LookupSourceTrelloOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceTrelloOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceTrelloArgs)(nil)).Elem()
}

// A collection of values returned by getSourceTrello.
type LookupSourceTrelloResultOutput struct{ *pulumi.OutputState }

func (LookupSourceTrelloResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceTrelloResult)(nil)).Elem()
}

func (o LookupSourceTrelloResultOutput) ToLookupSourceTrelloResultOutput() LookupSourceTrelloResultOutput {
	return o
}

func (o LookupSourceTrelloResultOutput) ToLookupSourceTrelloResultOutputWithContext(ctx context.Context) LookupSourceTrelloResultOutput {
	return o
}

func (o LookupSourceTrelloResultOutput) Configuration() GetSourceTrelloConfigurationOutput {
	return o.ApplyT(func(v LookupSourceTrelloResult) GetSourceTrelloConfiguration { return v.Configuration }).(GetSourceTrelloConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceTrelloResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceTrelloResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceTrelloResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceTrelloResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceTrelloResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceTrelloResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceTrelloResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceTrelloResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceTrelloResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceTrelloResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceTrelloResultOutput{})
}
