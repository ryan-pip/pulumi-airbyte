// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"internal"
)

// DestinationMysql DataSource
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
// func main() {
// pulumi.Run(func(ctx *pulumi.Context) error {
// _, err := airbyte.LookupDestinationMysql(ctx, %!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference), nil);
// if err != nil {
// return err
// }
// return nil
// })
// }
// ```
func LookupDestinationMysql(ctx *pulumi.Context, args *LookupDestinationMysqlArgs, opts ...pulumi.InvokeOption) (*LookupDestinationMysqlResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDestinationMysqlResult
	err := ctx.Invoke("airbyte:index/getDestinationMysql:getDestinationMysql", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDestinationMysql.
type LookupDestinationMysqlArgs struct {
	DestinationId string `pulumi:"destinationId"`
}

// A collection of values returned by getDestinationMysql.
type LookupDestinationMysqlResult struct {
	Configuration GetDestinationMysqlConfiguration `pulumi:"configuration"`
	DestinationId string                           `pulumi:"destinationId"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupDestinationMysqlOutput(ctx *pulumi.Context, args LookupDestinationMysqlOutputArgs, opts ...pulumi.InvokeOption) LookupDestinationMysqlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDestinationMysqlResult, error) {
			args := v.(LookupDestinationMysqlArgs)
			r, err := LookupDestinationMysql(ctx, &args, opts...)
			var s LookupDestinationMysqlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDestinationMysqlResultOutput)
}

// A collection of arguments for invoking getDestinationMysql.
type LookupDestinationMysqlOutputArgs struct {
	DestinationId pulumi.StringInput `pulumi:"destinationId"`
}

func (LookupDestinationMysqlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDestinationMysqlArgs)(nil)).Elem()
}

// A collection of values returned by getDestinationMysql.
type LookupDestinationMysqlResultOutput struct{ *pulumi.OutputState }

func (LookupDestinationMysqlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDestinationMysqlResult)(nil)).Elem()
}

func (o LookupDestinationMysqlResultOutput) ToLookupDestinationMysqlResultOutput() LookupDestinationMysqlResultOutput {
	return o
}

func (o LookupDestinationMysqlResultOutput) ToLookupDestinationMysqlResultOutputWithContext(ctx context.Context) LookupDestinationMysqlResultOutput {
	return o
}

func (o LookupDestinationMysqlResultOutput) Configuration() GetDestinationMysqlConfigurationOutput {
	return o.ApplyT(func(v LookupDestinationMysqlResult) GetDestinationMysqlConfiguration { return v.Configuration }).(GetDestinationMysqlConfigurationOutput)
}

func (o LookupDestinationMysqlResultOutput) DestinationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationMysqlResult) string { return v.DestinationId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDestinationMysqlResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationMysqlResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDestinationMysqlResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationMysqlResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDestinationMysqlResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationMysqlResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDestinationMysqlResultOutput{})
}
