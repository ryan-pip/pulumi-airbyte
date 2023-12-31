// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// SourceDockerhub DataSource
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
//			_, err := airbyte.LookupSourceDockerhub(ctx, &airbyte.LookupSourceDockerhubArgs{
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
func LookupSourceDockerhub(ctx *pulumi.Context, args *LookupSourceDockerhubArgs, opts ...pulumi.InvokeOption) (*LookupSourceDockerhubResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceDockerhubResult
	err := ctx.Invoke("airbyte:index/getSourceDockerhub:getSourceDockerhub", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceDockerhub.
type LookupSourceDockerhubArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceDockerhub.
type LookupSourceDockerhubResult struct {
	Configuration GetSourceDockerhubConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceDockerhubOutput(ctx *pulumi.Context, args LookupSourceDockerhubOutputArgs, opts ...pulumi.InvokeOption) LookupSourceDockerhubResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceDockerhubResult, error) {
			args := v.(LookupSourceDockerhubArgs)
			r, err := LookupSourceDockerhub(ctx, &args, opts...)
			var s LookupSourceDockerhubResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceDockerhubResultOutput)
}

// A collection of arguments for invoking getSourceDockerhub.
type LookupSourceDockerhubOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceDockerhubOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceDockerhubArgs)(nil)).Elem()
}

// A collection of values returned by getSourceDockerhub.
type LookupSourceDockerhubResultOutput struct{ *pulumi.OutputState }

func (LookupSourceDockerhubResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceDockerhubResult)(nil)).Elem()
}

func (o LookupSourceDockerhubResultOutput) ToLookupSourceDockerhubResultOutput() LookupSourceDockerhubResultOutput {
	return o
}

func (o LookupSourceDockerhubResultOutput) ToLookupSourceDockerhubResultOutputWithContext(ctx context.Context) LookupSourceDockerhubResultOutput {
	return o
}

func (o LookupSourceDockerhubResultOutput) Configuration() GetSourceDockerhubConfigurationOutput {
	return o.ApplyT(func(v LookupSourceDockerhubResult) GetSourceDockerhubConfiguration { return v.Configuration }).(GetSourceDockerhubConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceDockerhubResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceDockerhubResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceDockerhubResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceDockerhubResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceDockerhubResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceDockerhubResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceDockerhubResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceDockerhubResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceDockerhubResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceDockerhubResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceDockerhubResultOutput{})
}
