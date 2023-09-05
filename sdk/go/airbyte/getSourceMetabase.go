// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-airbyte/sdk/go/airbyte/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SourceMetabase DataSource
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-airbyte/sdk/go/airbyte"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := airbyte.LookupSourceMetabase(ctx, &airbyte.LookupSourceMetabaseArgs{
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
func LookupSourceMetabase(ctx *pulumi.Context, args *LookupSourceMetabaseArgs, opts ...pulumi.InvokeOption) (*LookupSourceMetabaseResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceMetabaseResult
	err := ctx.Invoke("airbyte:index/getSourceMetabase:getSourceMetabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceMetabase.
type LookupSourceMetabaseArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceMetabase.
type LookupSourceMetabaseResult struct {
	Configuration GetSourceMetabaseConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceMetabaseOutput(ctx *pulumi.Context, args LookupSourceMetabaseOutputArgs, opts ...pulumi.InvokeOption) LookupSourceMetabaseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceMetabaseResult, error) {
			args := v.(LookupSourceMetabaseArgs)
			r, err := LookupSourceMetabase(ctx, &args, opts...)
			var s LookupSourceMetabaseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceMetabaseResultOutput)
}

// A collection of arguments for invoking getSourceMetabase.
type LookupSourceMetabaseOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceMetabaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceMetabaseArgs)(nil)).Elem()
}

// A collection of values returned by getSourceMetabase.
type LookupSourceMetabaseResultOutput struct{ *pulumi.OutputState }

func (LookupSourceMetabaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceMetabaseResult)(nil)).Elem()
}

func (o LookupSourceMetabaseResultOutput) ToLookupSourceMetabaseResultOutput() LookupSourceMetabaseResultOutput {
	return o
}

func (o LookupSourceMetabaseResultOutput) ToLookupSourceMetabaseResultOutputWithContext(ctx context.Context) LookupSourceMetabaseResultOutput {
	return o
}

func (o LookupSourceMetabaseResultOutput) Configuration() GetSourceMetabaseConfigurationOutput {
	return o.ApplyT(func(v LookupSourceMetabaseResult) GetSourceMetabaseConfiguration { return v.Configuration }).(GetSourceMetabaseConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceMetabaseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceMetabaseResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceMetabaseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceMetabaseResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceMetabaseResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceMetabaseResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceMetabaseResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceMetabaseResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceMetabaseResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceMetabaseResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceMetabaseResultOutput{})
}
