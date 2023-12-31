// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// SourceEmailoctopus DataSource
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
//			_, err := airbyte.LookupSourceEmailoctopus(ctx, &airbyte.LookupSourceEmailoctopusArgs{
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
func LookupSourceEmailoctopus(ctx *pulumi.Context, args *LookupSourceEmailoctopusArgs, opts ...pulumi.InvokeOption) (*LookupSourceEmailoctopusResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceEmailoctopusResult
	err := ctx.Invoke("airbyte:index/getSourceEmailoctopus:getSourceEmailoctopus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceEmailoctopus.
type LookupSourceEmailoctopusArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceEmailoctopus.
type LookupSourceEmailoctopusResult struct {
	Configuration GetSourceEmailoctopusConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceEmailoctopusOutput(ctx *pulumi.Context, args LookupSourceEmailoctopusOutputArgs, opts ...pulumi.InvokeOption) LookupSourceEmailoctopusResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceEmailoctopusResult, error) {
			args := v.(LookupSourceEmailoctopusArgs)
			r, err := LookupSourceEmailoctopus(ctx, &args, opts...)
			var s LookupSourceEmailoctopusResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceEmailoctopusResultOutput)
}

// A collection of arguments for invoking getSourceEmailoctopus.
type LookupSourceEmailoctopusOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceEmailoctopusOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceEmailoctopusArgs)(nil)).Elem()
}

// A collection of values returned by getSourceEmailoctopus.
type LookupSourceEmailoctopusResultOutput struct{ *pulumi.OutputState }

func (LookupSourceEmailoctopusResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceEmailoctopusResult)(nil)).Elem()
}

func (o LookupSourceEmailoctopusResultOutput) ToLookupSourceEmailoctopusResultOutput() LookupSourceEmailoctopusResultOutput {
	return o
}

func (o LookupSourceEmailoctopusResultOutput) ToLookupSourceEmailoctopusResultOutputWithContext(ctx context.Context) LookupSourceEmailoctopusResultOutput {
	return o
}

func (o LookupSourceEmailoctopusResultOutput) Configuration() GetSourceEmailoctopusConfigurationOutput {
	return o.ApplyT(func(v LookupSourceEmailoctopusResult) GetSourceEmailoctopusConfiguration { return v.Configuration }).(GetSourceEmailoctopusConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceEmailoctopusResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceEmailoctopusResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceEmailoctopusResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceEmailoctopusResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceEmailoctopusResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceEmailoctopusResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceEmailoctopusResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceEmailoctopusResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceEmailoctopusResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceEmailoctopusResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceEmailoctopusResultOutput{})
}
