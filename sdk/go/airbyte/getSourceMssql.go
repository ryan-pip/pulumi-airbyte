// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// SourceMssql DataSource
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
//			_, err := airbyte.LookupSourceMssql(ctx, &airbyte.LookupSourceMssqlArgs{
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
func LookupSourceMssql(ctx *pulumi.Context, args *LookupSourceMssqlArgs, opts ...pulumi.InvokeOption) (*LookupSourceMssqlResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceMssqlResult
	err := ctx.Invoke("airbyte:index/getSourceMssql:getSourceMssql", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceMssql.
type LookupSourceMssqlArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceMssql.
type LookupSourceMssqlResult struct {
	Configuration GetSourceMssqlConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceMssqlOutput(ctx *pulumi.Context, args LookupSourceMssqlOutputArgs, opts ...pulumi.InvokeOption) LookupSourceMssqlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceMssqlResult, error) {
			args := v.(LookupSourceMssqlArgs)
			r, err := LookupSourceMssql(ctx, &args, opts...)
			var s LookupSourceMssqlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceMssqlResultOutput)
}

// A collection of arguments for invoking getSourceMssql.
type LookupSourceMssqlOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceMssqlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceMssqlArgs)(nil)).Elem()
}

// A collection of values returned by getSourceMssql.
type LookupSourceMssqlResultOutput struct{ *pulumi.OutputState }

func (LookupSourceMssqlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceMssqlResult)(nil)).Elem()
}

func (o LookupSourceMssqlResultOutput) ToLookupSourceMssqlResultOutput() LookupSourceMssqlResultOutput {
	return o
}

func (o LookupSourceMssqlResultOutput) ToLookupSourceMssqlResultOutputWithContext(ctx context.Context) LookupSourceMssqlResultOutput {
	return o
}

func (o LookupSourceMssqlResultOutput) Configuration() GetSourceMssqlConfigurationOutput {
	return o.ApplyT(func(v LookupSourceMssqlResult) GetSourceMssqlConfiguration { return v.Configuration }).(GetSourceMssqlConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceMssqlResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceMssqlResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceMssqlResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceMssqlResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceMssqlResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceMssqlResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceMssqlResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceMssqlResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceMssqlResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceMssqlResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceMssqlResultOutput{})
}
