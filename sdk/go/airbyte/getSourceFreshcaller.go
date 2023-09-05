// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-airbyte/sdk/go/airbyte/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SourceFreshcaller DataSource
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
//			_, err := airbyte.LookupSourceFreshcaller(ctx, &airbyte.LookupSourceFreshcallerArgs{
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
func LookupSourceFreshcaller(ctx *pulumi.Context, args *LookupSourceFreshcallerArgs, opts ...pulumi.InvokeOption) (*LookupSourceFreshcallerResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceFreshcallerResult
	err := ctx.Invoke("airbyte:index/getSourceFreshcaller:getSourceFreshcaller", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceFreshcaller.
type LookupSourceFreshcallerArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceFreshcaller.
type LookupSourceFreshcallerResult struct {
	Configuration GetSourceFreshcallerConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceFreshcallerOutput(ctx *pulumi.Context, args LookupSourceFreshcallerOutputArgs, opts ...pulumi.InvokeOption) LookupSourceFreshcallerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceFreshcallerResult, error) {
			args := v.(LookupSourceFreshcallerArgs)
			r, err := LookupSourceFreshcaller(ctx, &args, opts...)
			var s LookupSourceFreshcallerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceFreshcallerResultOutput)
}

// A collection of arguments for invoking getSourceFreshcaller.
type LookupSourceFreshcallerOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceFreshcallerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceFreshcallerArgs)(nil)).Elem()
}

// A collection of values returned by getSourceFreshcaller.
type LookupSourceFreshcallerResultOutput struct{ *pulumi.OutputState }

func (LookupSourceFreshcallerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceFreshcallerResult)(nil)).Elem()
}

func (o LookupSourceFreshcallerResultOutput) ToLookupSourceFreshcallerResultOutput() LookupSourceFreshcallerResultOutput {
	return o
}

func (o LookupSourceFreshcallerResultOutput) ToLookupSourceFreshcallerResultOutputWithContext(ctx context.Context) LookupSourceFreshcallerResultOutput {
	return o
}

func (o LookupSourceFreshcallerResultOutput) Configuration() GetSourceFreshcallerConfigurationOutput {
	return o.ApplyT(func(v LookupSourceFreshcallerResult) GetSourceFreshcallerConfiguration { return v.Configuration }).(GetSourceFreshcallerConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceFreshcallerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceFreshcallerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceFreshcallerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceFreshcallerResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceFreshcallerResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceFreshcallerResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceFreshcallerResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceFreshcallerResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceFreshcallerResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceFreshcallerResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceFreshcallerResultOutput{})
}
