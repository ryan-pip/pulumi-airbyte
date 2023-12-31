// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// SourceMailgun DataSource
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
//			_, err := airbyte.LookupSourceMailgun(ctx, &airbyte.LookupSourceMailgunArgs{
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
func LookupSourceMailgun(ctx *pulumi.Context, args *LookupSourceMailgunArgs, opts ...pulumi.InvokeOption) (*LookupSourceMailgunResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceMailgunResult
	err := ctx.Invoke("airbyte:index/getSourceMailgun:getSourceMailgun", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceMailgun.
type LookupSourceMailgunArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceMailgun.
type LookupSourceMailgunResult struct {
	Configuration GetSourceMailgunConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceMailgunOutput(ctx *pulumi.Context, args LookupSourceMailgunOutputArgs, opts ...pulumi.InvokeOption) LookupSourceMailgunResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceMailgunResult, error) {
			args := v.(LookupSourceMailgunArgs)
			r, err := LookupSourceMailgun(ctx, &args, opts...)
			var s LookupSourceMailgunResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceMailgunResultOutput)
}

// A collection of arguments for invoking getSourceMailgun.
type LookupSourceMailgunOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceMailgunOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceMailgunArgs)(nil)).Elem()
}

// A collection of values returned by getSourceMailgun.
type LookupSourceMailgunResultOutput struct{ *pulumi.OutputState }

func (LookupSourceMailgunResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceMailgunResult)(nil)).Elem()
}

func (o LookupSourceMailgunResultOutput) ToLookupSourceMailgunResultOutput() LookupSourceMailgunResultOutput {
	return o
}

func (o LookupSourceMailgunResultOutput) ToLookupSourceMailgunResultOutputWithContext(ctx context.Context) LookupSourceMailgunResultOutput {
	return o
}

func (o LookupSourceMailgunResultOutput) Configuration() GetSourceMailgunConfigurationOutput {
	return o.ApplyT(func(v LookupSourceMailgunResult) GetSourceMailgunConfiguration { return v.Configuration }).(GetSourceMailgunConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceMailgunResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceMailgunResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceMailgunResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceMailgunResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceMailgunResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceMailgunResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceMailgunResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceMailgunResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceMailgunResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceMailgunResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceMailgunResultOutput{})
}
