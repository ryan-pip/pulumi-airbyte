// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// SourceFreshdesk DataSource
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
//			_, err := airbyte.LookupSourceFreshdesk(ctx, &airbyte.LookupSourceFreshdeskArgs{
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
func LookupSourceFreshdesk(ctx *pulumi.Context, args *LookupSourceFreshdeskArgs, opts ...pulumi.InvokeOption) (*LookupSourceFreshdeskResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceFreshdeskResult
	err := ctx.Invoke("airbyte:index/getSourceFreshdesk:getSourceFreshdesk", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceFreshdesk.
type LookupSourceFreshdeskArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceFreshdesk.
type LookupSourceFreshdeskResult struct {
	Configuration GetSourceFreshdeskConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceFreshdeskOutput(ctx *pulumi.Context, args LookupSourceFreshdeskOutputArgs, opts ...pulumi.InvokeOption) LookupSourceFreshdeskResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceFreshdeskResult, error) {
			args := v.(LookupSourceFreshdeskArgs)
			r, err := LookupSourceFreshdesk(ctx, &args, opts...)
			var s LookupSourceFreshdeskResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceFreshdeskResultOutput)
}

// A collection of arguments for invoking getSourceFreshdesk.
type LookupSourceFreshdeskOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceFreshdeskOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceFreshdeskArgs)(nil)).Elem()
}

// A collection of values returned by getSourceFreshdesk.
type LookupSourceFreshdeskResultOutput struct{ *pulumi.OutputState }

func (LookupSourceFreshdeskResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceFreshdeskResult)(nil)).Elem()
}

func (o LookupSourceFreshdeskResultOutput) ToLookupSourceFreshdeskResultOutput() LookupSourceFreshdeskResultOutput {
	return o
}

func (o LookupSourceFreshdeskResultOutput) ToLookupSourceFreshdeskResultOutputWithContext(ctx context.Context) LookupSourceFreshdeskResultOutput {
	return o
}

func (o LookupSourceFreshdeskResultOutput) Configuration() GetSourceFreshdeskConfigurationOutput {
	return o.ApplyT(func(v LookupSourceFreshdeskResult) GetSourceFreshdeskConfiguration { return v.Configuration }).(GetSourceFreshdeskConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceFreshdeskResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceFreshdeskResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceFreshdeskResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceFreshdeskResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceFreshdeskResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceFreshdeskResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceFreshdeskResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceFreshdeskResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceFreshdeskResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceFreshdeskResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceFreshdeskResultOutput{})
}
