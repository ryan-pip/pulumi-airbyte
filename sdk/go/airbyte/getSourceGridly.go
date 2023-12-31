// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// SourceGridly DataSource
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
//			_, err := airbyte.LookupSourceGridly(ctx, &airbyte.LookupSourceGridlyArgs{
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
func LookupSourceGridly(ctx *pulumi.Context, args *LookupSourceGridlyArgs, opts ...pulumi.InvokeOption) (*LookupSourceGridlyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceGridlyResult
	err := ctx.Invoke("airbyte:index/getSourceGridly:getSourceGridly", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceGridly.
type LookupSourceGridlyArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceGridly.
type LookupSourceGridlyResult struct {
	Configuration GetSourceGridlyConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceGridlyOutput(ctx *pulumi.Context, args LookupSourceGridlyOutputArgs, opts ...pulumi.InvokeOption) LookupSourceGridlyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceGridlyResult, error) {
			args := v.(LookupSourceGridlyArgs)
			r, err := LookupSourceGridly(ctx, &args, opts...)
			var s LookupSourceGridlyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceGridlyResultOutput)
}

// A collection of arguments for invoking getSourceGridly.
type LookupSourceGridlyOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceGridlyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceGridlyArgs)(nil)).Elem()
}

// A collection of values returned by getSourceGridly.
type LookupSourceGridlyResultOutput struct{ *pulumi.OutputState }

func (LookupSourceGridlyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceGridlyResult)(nil)).Elem()
}

func (o LookupSourceGridlyResultOutput) ToLookupSourceGridlyResultOutput() LookupSourceGridlyResultOutput {
	return o
}

func (o LookupSourceGridlyResultOutput) ToLookupSourceGridlyResultOutputWithContext(ctx context.Context) LookupSourceGridlyResultOutput {
	return o
}

func (o LookupSourceGridlyResultOutput) Configuration() GetSourceGridlyConfigurationOutput {
	return o.ApplyT(func(v LookupSourceGridlyResult) GetSourceGridlyConfiguration { return v.Configuration }).(GetSourceGridlyConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceGridlyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceGridlyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceGridlyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceGridlyResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceGridlyResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceGridlyResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceGridlyResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceGridlyResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceGridlyResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceGridlyResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceGridlyResultOutput{})
}
