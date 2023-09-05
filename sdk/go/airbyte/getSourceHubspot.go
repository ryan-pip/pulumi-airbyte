// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-airbyte/sdk/go/airbyte/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SourceHubspot DataSource
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
//			_, err := airbyte.LookupSourceHubspot(ctx, &airbyte.LookupSourceHubspotArgs{
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
func LookupSourceHubspot(ctx *pulumi.Context, args *LookupSourceHubspotArgs, opts ...pulumi.InvokeOption) (*LookupSourceHubspotResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceHubspotResult
	err := ctx.Invoke("airbyte:index/getSourceHubspot:getSourceHubspot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceHubspot.
type LookupSourceHubspotArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceHubspot.
type LookupSourceHubspotResult struct {
	Configuration GetSourceHubspotConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceHubspotOutput(ctx *pulumi.Context, args LookupSourceHubspotOutputArgs, opts ...pulumi.InvokeOption) LookupSourceHubspotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceHubspotResult, error) {
			args := v.(LookupSourceHubspotArgs)
			r, err := LookupSourceHubspot(ctx, &args, opts...)
			var s LookupSourceHubspotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceHubspotResultOutput)
}

// A collection of arguments for invoking getSourceHubspot.
type LookupSourceHubspotOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceHubspotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceHubspotArgs)(nil)).Elem()
}

// A collection of values returned by getSourceHubspot.
type LookupSourceHubspotResultOutput struct{ *pulumi.OutputState }

func (LookupSourceHubspotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceHubspotResult)(nil)).Elem()
}

func (o LookupSourceHubspotResultOutput) ToLookupSourceHubspotResultOutput() LookupSourceHubspotResultOutput {
	return o
}

func (o LookupSourceHubspotResultOutput) ToLookupSourceHubspotResultOutputWithContext(ctx context.Context) LookupSourceHubspotResultOutput {
	return o
}

func (o LookupSourceHubspotResultOutput) Configuration() GetSourceHubspotConfigurationOutput {
	return o.ApplyT(func(v LookupSourceHubspotResult) GetSourceHubspotConfiguration { return v.Configuration }).(GetSourceHubspotConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceHubspotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceHubspotResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceHubspotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceHubspotResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceHubspotResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceHubspotResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceHubspotResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceHubspotResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceHubspotResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceHubspotResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceHubspotResultOutput{})
}
