// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// SourceQualaroo DataSource
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
//			_, err := airbyte.LookupSourceQualaroo(ctx, &airbyte.LookupSourceQualarooArgs{
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
func LookupSourceQualaroo(ctx *pulumi.Context, args *LookupSourceQualarooArgs, opts ...pulumi.InvokeOption) (*LookupSourceQualarooResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceQualarooResult
	err := ctx.Invoke("airbyte:index/getSourceQualaroo:getSourceQualaroo", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceQualaroo.
type LookupSourceQualarooArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceQualaroo.
type LookupSourceQualarooResult struct {
	Configuration GetSourceQualarooConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceQualarooOutput(ctx *pulumi.Context, args LookupSourceQualarooOutputArgs, opts ...pulumi.InvokeOption) LookupSourceQualarooResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceQualarooResult, error) {
			args := v.(LookupSourceQualarooArgs)
			r, err := LookupSourceQualaroo(ctx, &args, opts...)
			var s LookupSourceQualarooResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceQualarooResultOutput)
}

// A collection of arguments for invoking getSourceQualaroo.
type LookupSourceQualarooOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceQualarooOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceQualarooArgs)(nil)).Elem()
}

// A collection of values returned by getSourceQualaroo.
type LookupSourceQualarooResultOutput struct{ *pulumi.OutputState }

func (LookupSourceQualarooResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceQualarooResult)(nil)).Elem()
}

func (o LookupSourceQualarooResultOutput) ToLookupSourceQualarooResultOutput() LookupSourceQualarooResultOutput {
	return o
}

func (o LookupSourceQualarooResultOutput) ToLookupSourceQualarooResultOutputWithContext(ctx context.Context) LookupSourceQualarooResultOutput {
	return o
}

func (o LookupSourceQualarooResultOutput) Configuration() GetSourceQualarooConfigurationOutput {
	return o.ApplyT(func(v LookupSourceQualarooResult) GetSourceQualarooConfiguration { return v.Configuration }).(GetSourceQualarooConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceQualarooResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceQualarooResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceQualarooResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceQualarooResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceQualarooResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceQualarooResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceQualarooResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceQualarooResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceQualarooResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceQualarooResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceQualarooResultOutput{})
}
