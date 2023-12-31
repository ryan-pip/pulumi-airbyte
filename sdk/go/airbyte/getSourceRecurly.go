// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// SourceRecurly DataSource
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
//			_, err := airbyte.LookupSourceRecurly(ctx, &airbyte.LookupSourceRecurlyArgs{
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
func LookupSourceRecurly(ctx *pulumi.Context, args *LookupSourceRecurlyArgs, opts ...pulumi.InvokeOption) (*LookupSourceRecurlyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceRecurlyResult
	err := ctx.Invoke("airbyte:index/getSourceRecurly:getSourceRecurly", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceRecurly.
type LookupSourceRecurlyArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceRecurly.
type LookupSourceRecurlyResult struct {
	Configuration GetSourceRecurlyConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceRecurlyOutput(ctx *pulumi.Context, args LookupSourceRecurlyOutputArgs, opts ...pulumi.InvokeOption) LookupSourceRecurlyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceRecurlyResult, error) {
			args := v.(LookupSourceRecurlyArgs)
			r, err := LookupSourceRecurly(ctx, &args, opts...)
			var s LookupSourceRecurlyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceRecurlyResultOutput)
}

// A collection of arguments for invoking getSourceRecurly.
type LookupSourceRecurlyOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceRecurlyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceRecurlyArgs)(nil)).Elem()
}

// A collection of values returned by getSourceRecurly.
type LookupSourceRecurlyResultOutput struct{ *pulumi.OutputState }

func (LookupSourceRecurlyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceRecurlyResult)(nil)).Elem()
}

func (o LookupSourceRecurlyResultOutput) ToLookupSourceRecurlyResultOutput() LookupSourceRecurlyResultOutput {
	return o
}

func (o LookupSourceRecurlyResultOutput) ToLookupSourceRecurlyResultOutputWithContext(ctx context.Context) LookupSourceRecurlyResultOutput {
	return o
}

func (o LookupSourceRecurlyResultOutput) Configuration() GetSourceRecurlyConfigurationOutput {
	return o.ApplyT(func(v LookupSourceRecurlyResult) GetSourceRecurlyConfiguration { return v.Configuration }).(GetSourceRecurlyConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceRecurlyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceRecurlyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceRecurlyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceRecurlyResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceRecurlyResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceRecurlyResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceRecurlyResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceRecurlyResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceRecurlyResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceRecurlyResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceRecurlyResultOutput{})
}
