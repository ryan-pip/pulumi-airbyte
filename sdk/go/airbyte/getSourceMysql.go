// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// SourceMysql DataSource
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
//			_, err := airbyte.LookupSourceMysql(ctx, &airbyte.LookupSourceMysqlArgs{
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
func LookupSourceMysql(ctx *pulumi.Context, args *LookupSourceMysqlArgs, opts ...pulumi.InvokeOption) (*LookupSourceMysqlResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceMysqlResult
	err := ctx.Invoke("airbyte:index/getSourceMysql:getSourceMysql", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceMysql.
type LookupSourceMysqlArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceMysql.
type LookupSourceMysqlResult struct {
	Configuration GetSourceMysqlConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceMysqlOutput(ctx *pulumi.Context, args LookupSourceMysqlOutputArgs, opts ...pulumi.InvokeOption) LookupSourceMysqlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceMysqlResult, error) {
			args := v.(LookupSourceMysqlArgs)
			r, err := LookupSourceMysql(ctx, &args, opts...)
			var s LookupSourceMysqlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceMysqlResultOutput)
}

// A collection of arguments for invoking getSourceMysql.
type LookupSourceMysqlOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceMysqlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceMysqlArgs)(nil)).Elem()
}

// A collection of values returned by getSourceMysql.
type LookupSourceMysqlResultOutput struct{ *pulumi.OutputState }

func (LookupSourceMysqlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceMysqlResult)(nil)).Elem()
}

func (o LookupSourceMysqlResultOutput) ToLookupSourceMysqlResultOutput() LookupSourceMysqlResultOutput {
	return o
}

func (o LookupSourceMysqlResultOutput) ToLookupSourceMysqlResultOutputWithContext(ctx context.Context) LookupSourceMysqlResultOutput {
	return o
}

func (o LookupSourceMysqlResultOutput) Configuration() GetSourceMysqlConfigurationOutput {
	return o.ApplyT(func(v LookupSourceMysqlResult) GetSourceMysqlConfiguration { return v.Configuration }).(GetSourceMysqlConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceMysqlResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceMysqlResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceMysqlResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceMysqlResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceMysqlResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceMysqlResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceMysqlResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceMysqlResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceMysqlResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceMysqlResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceMysqlResultOutput{})
}
