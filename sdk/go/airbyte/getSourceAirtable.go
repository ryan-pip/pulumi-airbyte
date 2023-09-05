// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// SourceAirtable DataSource
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
//			_, err := airbyte.LookupSourceAirtable(ctx, &airbyte.LookupSourceAirtableArgs{
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
func LookupSourceAirtable(ctx *pulumi.Context, args *LookupSourceAirtableArgs, opts ...pulumi.InvokeOption) (*LookupSourceAirtableResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceAirtableResult
	err := ctx.Invoke("airbyte:index/getSourceAirtable:getSourceAirtable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceAirtable.
type LookupSourceAirtableArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceAirtable.
type LookupSourceAirtableResult struct {
	Configuration GetSourceAirtableConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceAirtableOutput(ctx *pulumi.Context, args LookupSourceAirtableOutputArgs, opts ...pulumi.InvokeOption) LookupSourceAirtableResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceAirtableResult, error) {
			args := v.(LookupSourceAirtableArgs)
			r, err := LookupSourceAirtable(ctx, &args, opts...)
			var s LookupSourceAirtableResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceAirtableResultOutput)
}

// A collection of arguments for invoking getSourceAirtable.
type LookupSourceAirtableOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceAirtableOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceAirtableArgs)(nil)).Elem()
}

// A collection of values returned by getSourceAirtable.
type LookupSourceAirtableResultOutput struct{ *pulumi.OutputState }

func (LookupSourceAirtableResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceAirtableResult)(nil)).Elem()
}

func (o LookupSourceAirtableResultOutput) ToLookupSourceAirtableResultOutput() LookupSourceAirtableResultOutput {
	return o
}

func (o LookupSourceAirtableResultOutput) ToLookupSourceAirtableResultOutputWithContext(ctx context.Context) LookupSourceAirtableResultOutput {
	return o
}

func (o LookupSourceAirtableResultOutput) Configuration() GetSourceAirtableConfigurationOutput {
	return o.ApplyT(func(v LookupSourceAirtableResult) GetSourceAirtableConfiguration { return v.Configuration }).(GetSourceAirtableConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceAirtableResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceAirtableResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceAirtableResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceAirtableResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceAirtableResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceAirtableResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceAirtableResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceAirtableResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceAirtableResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceAirtableResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceAirtableResultOutput{})
}
