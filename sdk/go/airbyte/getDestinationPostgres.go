// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// DestinationPostgres DataSource
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
//			_, err := airbyte.LookupDestinationPostgres(ctx, &airbyte.LookupDestinationPostgresArgs{
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
func LookupDestinationPostgres(ctx *pulumi.Context, args *LookupDestinationPostgresArgs, opts ...pulumi.InvokeOption) (*LookupDestinationPostgresResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDestinationPostgresResult
	err := ctx.Invoke("airbyte:index/getDestinationPostgres:getDestinationPostgres", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDestinationPostgres.
type LookupDestinationPostgresArgs struct {
	DestinationId string `pulumi:"destinationId"`
}

// A collection of values returned by getDestinationPostgres.
type LookupDestinationPostgresResult struct {
	Configuration GetDestinationPostgresConfiguration `pulumi:"configuration"`
	DestinationId string                              `pulumi:"destinationId"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupDestinationPostgresOutput(ctx *pulumi.Context, args LookupDestinationPostgresOutputArgs, opts ...pulumi.InvokeOption) LookupDestinationPostgresResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDestinationPostgresResult, error) {
			args := v.(LookupDestinationPostgresArgs)
			r, err := LookupDestinationPostgres(ctx, &args, opts...)
			var s LookupDestinationPostgresResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDestinationPostgresResultOutput)
}

// A collection of arguments for invoking getDestinationPostgres.
type LookupDestinationPostgresOutputArgs struct {
	DestinationId pulumi.StringInput `pulumi:"destinationId"`
}

func (LookupDestinationPostgresOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDestinationPostgresArgs)(nil)).Elem()
}

// A collection of values returned by getDestinationPostgres.
type LookupDestinationPostgresResultOutput struct{ *pulumi.OutputState }

func (LookupDestinationPostgresResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDestinationPostgresResult)(nil)).Elem()
}

func (o LookupDestinationPostgresResultOutput) ToLookupDestinationPostgresResultOutput() LookupDestinationPostgresResultOutput {
	return o
}

func (o LookupDestinationPostgresResultOutput) ToLookupDestinationPostgresResultOutputWithContext(ctx context.Context) LookupDestinationPostgresResultOutput {
	return o
}

func (o LookupDestinationPostgresResultOutput) Configuration() GetDestinationPostgresConfigurationOutput {
	return o.ApplyT(func(v LookupDestinationPostgresResult) GetDestinationPostgresConfiguration { return v.Configuration }).(GetDestinationPostgresConfigurationOutput)
}

func (o LookupDestinationPostgresResultOutput) DestinationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationPostgresResult) string { return v.DestinationId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDestinationPostgresResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationPostgresResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDestinationPostgresResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationPostgresResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDestinationPostgresResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationPostgresResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDestinationPostgresResultOutput{})
}
