// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// SourceSlack DataSource
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
//			_, err := airbyte.LookupSourceSlack(ctx, &airbyte.LookupSourceSlackArgs{
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
func LookupSourceSlack(ctx *pulumi.Context, args *LookupSourceSlackArgs, opts ...pulumi.InvokeOption) (*LookupSourceSlackResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceSlackResult
	err := ctx.Invoke("airbyte:index/getSourceSlack:getSourceSlack", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceSlack.
type LookupSourceSlackArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceSlack.
type LookupSourceSlackResult struct {
	Configuration GetSourceSlackConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceSlackOutput(ctx *pulumi.Context, args LookupSourceSlackOutputArgs, opts ...pulumi.InvokeOption) LookupSourceSlackResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceSlackResult, error) {
			args := v.(LookupSourceSlackArgs)
			r, err := LookupSourceSlack(ctx, &args, opts...)
			var s LookupSourceSlackResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceSlackResultOutput)
}

// A collection of arguments for invoking getSourceSlack.
type LookupSourceSlackOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceSlackOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceSlackArgs)(nil)).Elem()
}

// A collection of values returned by getSourceSlack.
type LookupSourceSlackResultOutput struct{ *pulumi.OutputState }

func (LookupSourceSlackResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceSlackResult)(nil)).Elem()
}

func (o LookupSourceSlackResultOutput) ToLookupSourceSlackResultOutput() LookupSourceSlackResultOutput {
	return o
}

func (o LookupSourceSlackResultOutput) ToLookupSourceSlackResultOutputWithContext(ctx context.Context) LookupSourceSlackResultOutput {
	return o
}

func (o LookupSourceSlackResultOutput) Configuration() GetSourceSlackConfigurationOutput {
	return o.ApplyT(func(v LookupSourceSlackResult) GetSourceSlackConfiguration { return v.Configuration }).(GetSourceSlackConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceSlackResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSlackResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceSlackResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSlackResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceSlackResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceSlackResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceSlackResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSlackResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceSlackResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSlackResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceSlackResultOutput{})
}
