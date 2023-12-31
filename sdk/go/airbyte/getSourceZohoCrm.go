// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// SourceZohoCrm DataSource
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
//			_, err := airbyte.LookupSourceZohoCrm(ctx, &airbyte.LookupSourceZohoCrmArgs{
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
func LookupSourceZohoCrm(ctx *pulumi.Context, args *LookupSourceZohoCrmArgs, opts ...pulumi.InvokeOption) (*LookupSourceZohoCrmResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceZohoCrmResult
	err := ctx.Invoke("airbyte:index/getSourceZohoCrm:getSourceZohoCrm", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceZohoCrm.
type LookupSourceZohoCrmArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceZohoCrm.
type LookupSourceZohoCrmResult struct {
	Configuration GetSourceZohoCrmConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceZohoCrmOutput(ctx *pulumi.Context, args LookupSourceZohoCrmOutputArgs, opts ...pulumi.InvokeOption) LookupSourceZohoCrmResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceZohoCrmResult, error) {
			args := v.(LookupSourceZohoCrmArgs)
			r, err := LookupSourceZohoCrm(ctx, &args, opts...)
			var s LookupSourceZohoCrmResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceZohoCrmResultOutput)
}

// A collection of arguments for invoking getSourceZohoCrm.
type LookupSourceZohoCrmOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceZohoCrmOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceZohoCrmArgs)(nil)).Elem()
}

// A collection of values returned by getSourceZohoCrm.
type LookupSourceZohoCrmResultOutput struct{ *pulumi.OutputState }

func (LookupSourceZohoCrmResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceZohoCrmResult)(nil)).Elem()
}

func (o LookupSourceZohoCrmResultOutput) ToLookupSourceZohoCrmResultOutput() LookupSourceZohoCrmResultOutput {
	return o
}

func (o LookupSourceZohoCrmResultOutput) ToLookupSourceZohoCrmResultOutputWithContext(ctx context.Context) LookupSourceZohoCrmResultOutput {
	return o
}

func (o LookupSourceZohoCrmResultOutput) Configuration() GetSourceZohoCrmConfigurationOutput {
	return o.ApplyT(func(v LookupSourceZohoCrmResult) GetSourceZohoCrmConfiguration { return v.Configuration }).(GetSourceZohoCrmConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceZohoCrmResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceZohoCrmResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceZohoCrmResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceZohoCrmResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceZohoCrmResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceZohoCrmResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceZohoCrmResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceZohoCrmResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceZohoCrmResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceZohoCrmResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceZohoCrmResultOutput{})
}
