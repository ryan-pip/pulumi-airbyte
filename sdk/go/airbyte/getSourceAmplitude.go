// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-airbyte/sdk/go/airbyte/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SourceAmplitude DataSource
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
//			_, err := airbyte.LookupSourceAmplitude(ctx, &airbyte.LookupSourceAmplitudeArgs{
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
func LookupSourceAmplitude(ctx *pulumi.Context, args *LookupSourceAmplitudeArgs, opts ...pulumi.InvokeOption) (*LookupSourceAmplitudeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceAmplitudeResult
	err := ctx.Invoke("airbyte:index/getSourceAmplitude:getSourceAmplitude", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceAmplitude.
type LookupSourceAmplitudeArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceAmplitude.
type LookupSourceAmplitudeResult struct {
	Configuration GetSourceAmplitudeConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceAmplitudeOutput(ctx *pulumi.Context, args LookupSourceAmplitudeOutputArgs, opts ...pulumi.InvokeOption) LookupSourceAmplitudeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceAmplitudeResult, error) {
			args := v.(LookupSourceAmplitudeArgs)
			r, err := LookupSourceAmplitude(ctx, &args, opts...)
			var s LookupSourceAmplitudeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceAmplitudeResultOutput)
}

// A collection of arguments for invoking getSourceAmplitude.
type LookupSourceAmplitudeOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceAmplitudeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceAmplitudeArgs)(nil)).Elem()
}

// A collection of values returned by getSourceAmplitude.
type LookupSourceAmplitudeResultOutput struct{ *pulumi.OutputState }

func (LookupSourceAmplitudeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceAmplitudeResult)(nil)).Elem()
}

func (o LookupSourceAmplitudeResultOutput) ToLookupSourceAmplitudeResultOutput() LookupSourceAmplitudeResultOutput {
	return o
}

func (o LookupSourceAmplitudeResultOutput) ToLookupSourceAmplitudeResultOutputWithContext(ctx context.Context) LookupSourceAmplitudeResultOutput {
	return o
}

func (o LookupSourceAmplitudeResultOutput) Configuration() GetSourceAmplitudeConfigurationOutput {
	return o.ApplyT(func(v LookupSourceAmplitudeResult) GetSourceAmplitudeConfiguration { return v.Configuration }).(GetSourceAmplitudeConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceAmplitudeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceAmplitudeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceAmplitudeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceAmplitudeResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceAmplitudeResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceAmplitudeResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceAmplitudeResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceAmplitudeResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceAmplitudeResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceAmplitudeResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceAmplitudeResultOutput{})
}
