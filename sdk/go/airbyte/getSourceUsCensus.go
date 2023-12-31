// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// SourceUsCensus DataSource
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
//			_, err := airbyte.LookupSourceUsCensus(ctx, &airbyte.LookupSourceUsCensusArgs{
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
func LookupSourceUsCensus(ctx *pulumi.Context, args *LookupSourceUsCensusArgs, opts ...pulumi.InvokeOption) (*LookupSourceUsCensusResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceUsCensusResult
	err := ctx.Invoke("airbyte:index/getSourceUsCensus:getSourceUsCensus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceUsCensus.
type LookupSourceUsCensusArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceUsCensus.
type LookupSourceUsCensusResult struct {
	Configuration GetSourceUsCensusConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceUsCensusOutput(ctx *pulumi.Context, args LookupSourceUsCensusOutputArgs, opts ...pulumi.InvokeOption) LookupSourceUsCensusResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceUsCensusResult, error) {
			args := v.(LookupSourceUsCensusArgs)
			r, err := LookupSourceUsCensus(ctx, &args, opts...)
			var s LookupSourceUsCensusResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceUsCensusResultOutput)
}

// A collection of arguments for invoking getSourceUsCensus.
type LookupSourceUsCensusOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceUsCensusOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceUsCensusArgs)(nil)).Elem()
}

// A collection of values returned by getSourceUsCensus.
type LookupSourceUsCensusResultOutput struct{ *pulumi.OutputState }

func (LookupSourceUsCensusResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceUsCensusResult)(nil)).Elem()
}

func (o LookupSourceUsCensusResultOutput) ToLookupSourceUsCensusResultOutput() LookupSourceUsCensusResultOutput {
	return o
}

func (o LookupSourceUsCensusResultOutput) ToLookupSourceUsCensusResultOutputWithContext(ctx context.Context) LookupSourceUsCensusResultOutput {
	return o
}

func (o LookupSourceUsCensusResultOutput) Configuration() GetSourceUsCensusConfigurationOutput {
	return o.ApplyT(func(v LookupSourceUsCensusResult) GetSourceUsCensusConfiguration { return v.Configuration }).(GetSourceUsCensusConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceUsCensusResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceUsCensusResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceUsCensusResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceUsCensusResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceUsCensusResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceUsCensusResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceUsCensusResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceUsCensusResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceUsCensusResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceUsCensusResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceUsCensusResultOutput{})
}
