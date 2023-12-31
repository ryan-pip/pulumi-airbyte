// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// SourceOkta DataSource
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
//			_, err := airbyte.LookupSourceOkta(ctx, &airbyte.LookupSourceOktaArgs{
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
func LookupSourceOkta(ctx *pulumi.Context, args *LookupSourceOktaArgs, opts ...pulumi.InvokeOption) (*LookupSourceOktaResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceOktaResult
	err := ctx.Invoke("airbyte:index/getSourceOkta:getSourceOkta", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceOkta.
type LookupSourceOktaArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceOkta.
type LookupSourceOktaResult struct {
	Configuration GetSourceOktaConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceOktaOutput(ctx *pulumi.Context, args LookupSourceOktaOutputArgs, opts ...pulumi.InvokeOption) LookupSourceOktaResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceOktaResult, error) {
			args := v.(LookupSourceOktaArgs)
			r, err := LookupSourceOkta(ctx, &args, opts...)
			var s LookupSourceOktaResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceOktaResultOutput)
}

// A collection of arguments for invoking getSourceOkta.
type LookupSourceOktaOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceOktaOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceOktaArgs)(nil)).Elem()
}

// A collection of values returned by getSourceOkta.
type LookupSourceOktaResultOutput struct{ *pulumi.OutputState }

func (LookupSourceOktaResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceOktaResult)(nil)).Elem()
}

func (o LookupSourceOktaResultOutput) ToLookupSourceOktaResultOutput() LookupSourceOktaResultOutput {
	return o
}

func (o LookupSourceOktaResultOutput) ToLookupSourceOktaResultOutputWithContext(ctx context.Context) LookupSourceOktaResultOutput {
	return o
}

func (o LookupSourceOktaResultOutput) Configuration() GetSourceOktaConfigurationOutput {
	return o.ApplyT(func(v LookupSourceOktaResult) GetSourceOktaConfiguration { return v.Configuration }).(GetSourceOktaConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceOktaResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceOktaResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceOktaResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceOktaResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceOktaResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceOktaResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceOktaResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceOktaResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceOktaResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceOktaResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceOktaResultOutput{})
}
