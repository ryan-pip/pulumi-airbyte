// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// SourceXkcd DataSource
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
//			_, err := airbyte.LookupSourceXkcd(ctx, &airbyte.LookupSourceXkcdArgs{
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
func LookupSourceXkcd(ctx *pulumi.Context, args *LookupSourceXkcdArgs, opts ...pulumi.InvokeOption) (*LookupSourceXkcdResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceXkcdResult
	err := ctx.Invoke("airbyte:index/getSourceXkcd:getSourceXkcd", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceXkcd.
type LookupSourceXkcdArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceXkcd.
type LookupSourceXkcdResult struct {
	Configuration GetSourceXkcdConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceXkcdOutput(ctx *pulumi.Context, args LookupSourceXkcdOutputArgs, opts ...pulumi.InvokeOption) LookupSourceXkcdResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceXkcdResult, error) {
			args := v.(LookupSourceXkcdArgs)
			r, err := LookupSourceXkcd(ctx, &args, opts...)
			var s LookupSourceXkcdResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceXkcdResultOutput)
}

// A collection of arguments for invoking getSourceXkcd.
type LookupSourceXkcdOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceXkcdOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceXkcdArgs)(nil)).Elem()
}

// A collection of values returned by getSourceXkcd.
type LookupSourceXkcdResultOutput struct{ *pulumi.OutputState }

func (LookupSourceXkcdResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceXkcdResult)(nil)).Elem()
}

func (o LookupSourceXkcdResultOutput) ToLookupSourceXkcdResultOutput() LookupSourceXkcdResultOutput {
	return o
}

func (o LookupSourceXkcdResultOutput) ToLookupSourceXkcdResultOutputWithContext(ctx context.Context) LookupSourceXkcdResultOutput {
	return o
}

func (o LookupSourceXkcdResultOutput) Configuration() GetSourceXkcdConfigurationOutput {
	return o.ApplyT(func(v LookupSourceXkcdResult) GetSourceXkcdConfiguration { return v.Configuration }).(GetSourceXkcdConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceXkcdResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceXkcdResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceXkcdResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceXkcdResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceXkcdResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceXkcdResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceXkcdResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceXkcdResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceXkcdResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceXkcdResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceXkcdResultOutput{})
}
