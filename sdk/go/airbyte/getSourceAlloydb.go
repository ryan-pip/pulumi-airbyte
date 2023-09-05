// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// SourceAlloydb DataSource
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
//			_, err := airbyte.LookupSourceAlloydb(ctx, &airbyte.LookupSourceAlloydbArgs{
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
func LookupSourceAlloydb(ctx *pulumi.Context, args *LookupSourceAlloydbArgs, opts ...pulumi.InvokeOption) (*LookupSourceAlloydbResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceAlloydbResult
	err := ctx.Invoke("airbyte:index/getSourceAlloydb:getSourceAlloydb", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceAlloydb.
type LookupSourceAlloydbArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceAlloydb.
type LookupSourceAlloydbResult struct {
	Configuration GetSourceAlloydbConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceAlloydbOutput(ctx *pulumi.Context, args LookupSourceAlloydbOutputArgs, opts ...pulumi.InvokeOption) LookupSourceAlloydbResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceAlloydbResult, error) {
			args := v.(LookupSourceAlloydbArgs)
			r, err := LookupSourceAlloydb(ctx, &args, opts...)
			var s LookupSourceAlloydbResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceAlloydbResultOutput)
}

// A collection of arguments for invoking getSourceAlloydb.
type LookupSourceAlloydbOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceAlloydbOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceAlloydbArgs)(nil)).Elem()
}

// A collection of values returned by getSourceAlloydb.
type LookupSourceAlloydbResultOutput struct{ *pulumi.OutputState }

func (LookupSourceAlloydbResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceAlloydbResult)(nil)).Elem()
}

func (o LookupSourceAlloydbResultOutput) ToLookupSourceAlloydbResultOutput() LookupSourceAlloydbResultOutput {
	return o
}

func (o LookupSourceAlloydbResultOutput) ToLookupSourceAlloydbResultOutputWithContext(ctx context.Context) LookupSourceAlloydbResultOutput {
	return o
}

func (o LookupSourceAlloydbResultOutput) Configuration() GetSourceAlloydbConfigurationOutput {
	return o.ApplyT(func(v LookupSourceAlloydbResult) GetSourceAlloydbConfiguration { return v.Configuration }).(GetSourceAlloydbConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceAlloydbResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceAlloydbResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceAlloydbResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceAlloydbResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceAlloydbResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceAlloydbResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceAlloydbResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceAlloydbResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceAlloydbResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceAlloydbResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceAlloydbResultOutput{})
}
