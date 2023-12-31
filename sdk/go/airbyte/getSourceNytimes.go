// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// SourceNytimes DataSource
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
//			_, err := airbyte.LookupSourceNytimes(ctx, &airbyte.LookupSourceNytimesArgs{
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
func LookupSourceNytimes(ctx *pulumi.Context, args *LookupSourceNytimesArgs, opts ...pulumi.InvokeOption) (*LookupSourceNytimesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceNytimesResult
	err := ctx.Invoke("airbyte:index/getSourceNytimes:getSourceNytimes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceNytimes.
type LookupSourceNytimesArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceNytimes.
type LookupSourceNytimesResult struct {
	Configuration GetSourceNytimesConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceNytimesOutput(ctx *pulumi.Context, args LookupSourceNytimesOutputArgs, opts ...pulumi.InvokeOption) LookupSourceNytimesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceNytimesResult, error) {
			args := v.(LookupSourceNytimesArgs)
			r, err := LookupSourceNytimes(ctx, &args, opts...)
			var s LookupSourceNytimesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceNytimesResultOutput)
}

// A collection of arguments for invoking getSourceNytimes.
type LookupSourceNytimesOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceNytimesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceNytimesArgs)(nil)).Elem()
}

// A collection of values returned by getSourceNytimes.
type LookupSourceNytimesResultOutput struct{ *pulumi.OutputState }

func (LookupSourceNytimesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceNytimesResult)(nil)).Elem()
}

func (o LookupSourceNytimesResultOutput) ToLookupSourceNytimesResultOutput() LookupSourceNytimesResultOutput {
	return o
}

func (o LookupSourceNytimesResultOutput) ToLookupSourceNytimesResultOutputWithContext(ctx context.Context) LookupSourceNytimesResultOutput {
	return o
}

func (o LookupSourceNytimesResultOutput) Configuration() GetSourceNytimesConfigurationOutput {
	return o.ApplyT(func(v LookupSourceNytimesResult) GetSourceNytimesConfiguration { return v.Configuration }).(GetSourceNytimesConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceNytimesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceNytimesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceNytimesResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceNytimesResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceNytimesResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceNytimesResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceNytimesResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceNytimesResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceNytimesResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceNytimesResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceNytimesResultOutput{})
}
