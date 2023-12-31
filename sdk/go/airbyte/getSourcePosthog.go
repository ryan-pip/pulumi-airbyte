// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// SourcePosthog DataSource
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
//			_, err := airbyte.LookupSourcePosthog(ctx, &airbyte.LookupSourcePosthogArgs{
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
func LookupSourcePosthog(ctx *pulumi.Context, args *LookupSourcePosthogArgs, opts ...pulumi.InvokeOption) (*LookupSourcePosthogResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourcePosthogResult
	err := ctx.Invoke("airbyte:index/getSourcePosthog:getSourcePosthog", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourcePosthog.
type LookupSourcePosthogArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourcePosthog.
type LookupSourcePosthogResult struct {
	Configuration GetSourcePosthogConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourcePosthogOutput(ctx *pulumi.Context, args LookupSourcePosthogOutputArgs, opts ...pulumi.InvokeOption) LookupSourcePosthogResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourcePosthogResult, error) {
			args := v.(LookupSourcePosthogArgs)
			r, err := LookupSourcePosthog(ctx, &args, opts...)
			var s LookupSourcePosthogResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourcePosthogResultOutput)
}

// A collection of arguments for invoking getSourcePosthog.
type LookupSourcePosthogOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourcePosthogOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourcePosthogArgs)(nil)).Elem()
}

// A collection of values returned by getSourcePosthog.
type LookupSourcePosthogResultOutput struct{ *pulumi.OutputState }

func (LookupSourcePosthogResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourcePosthogResult)(nil)).Elem()
}

func (o LookupSourcePosthogResultOutput) ToLookupSourcePosthogResultOutput() LookupSourcePosthogResultOutput {
	return o
}

func (o LookupSourcePosthogResultOutput) ToLookupSourcePosthogResultOutputWithContext(ctx context.Context) LookupSourcePosthogResultOutput {
	return o
}

func (o LookupSourcePosthogResultOutput) Configuration() GetSourcePosthogConfigurationOutput {
	return o.ApplyT(func(v LookupSourcePosthogResult) GetSourcePosthogConfiguration { return v.Configuration }).(GetSourcePosthogConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourcePosthogResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourcePosthogResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourcePosthogResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourcePosthogResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourcePosthogResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourcePosthogResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourcePosthogResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourcePosthogResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourcePosthogResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourcePosthogResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourcePosthogResultOutput{})
}
