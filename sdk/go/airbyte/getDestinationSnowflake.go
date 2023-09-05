// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-airbyte/sdk/go/airbyte/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// DestinationSnowflake DataSource
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
//			_, err := airbyte.LookupDestinationSnowflake(ctx, &airbyte.LookupDestinationSnowflakeArgs{
//				DestinationId: "...my_destination_id...",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupDestinationSnowflake(ctx *pulumi.Context, args *LookupDestinationSnowflakeArgs, opts ...pulumi.InvokeOption) (*LookupDestinationSnowflakeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDestinationSnowflakeResult
	err := ctx.Invoke("airbyte:index/getDestinationSnowflake:getDestinationSnowflake", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDestinationSnowflake.
type LookupDestinationSnowflakeArgs struct {
	DestinationId string `pulumi:"destinationId"`
}

// A collection of values returned by getDestinationSnowflake.
type LookupDestinationSnowflakeResult struct {
	Configuration GetDestinationSnowflakeConfiguration `pulumi:"configuration"`
	DestinationId string                               `pulumi:"destinationId"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupDestinationSnowflakeOutput(ctx *pulumi.Context, args LookupDestinationSnowflakeOutputArgs, opts ...pulumi.InvokeOption) LookupDestinationSnowflakeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDestinationSnowflakeResult, error) {
			args := v.(LookupDestinationSnowflakeArgs)
			r, err := LookupDestinationSnowflake(ctx, &args, opts...)
			var s LookupDestinationSnowflakeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDestinationSnowflakeResultOutput)
}

// A collection of arguments for invoking getDestinationSnowflake.
type LookupDestinationSnowflakeOutputArgs struct {
	DestinationId pulumi.StringInput `pulumi:"destinationId"`
}

func (LookupDestinationSnowflakeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDestinationSnowflakeArgs)(nil)).Elem()
}

// A collection of values returned by getDestinationSnowflake.
type LookupDestinationSnowflakeResultOutput struct{ *pulumi.OutputState }

func (LookupDestinationSnowflakeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDestinationSnowflakeResult)(nil)).Elem()
}

func (o LookupDestinationSnowflakeResultOutput) ToLookupDestinationSnowflakeResultOutput() LookupDestinationSnowflakeResultOutput {
	return o
}

func (o LookupDestinationSnowflakeResultOutput) ToLookupDestinationSnowflakeResultOutputWithContext(ctx context.Context) LookupDestinationSnowflakeResultOutput {
	return o
}

func (o LookupDestinationSnowflakeResultOutput) Configuration() GetDestinationSnowflakeConfigurationOutput {
	return o.ApplyT(func(v LookupDestinationSnowflakeResult) GetDestinationSnowflakeConfiguration { return v.Configuration }).(GetDestinationSnowflakeConfigurationOutput)
}

func (o LookupDestinationSnowflakeResultOutput) DestinationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationSnowflakeResult) string { return v.DestinationId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDestinationSnowflakeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationSnowflakeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDestinationSnowflakeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationSnowflakeResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDestinationSnowflakeResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationSnowflakeResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDestinationSnowflakeResultOutput{})
}
