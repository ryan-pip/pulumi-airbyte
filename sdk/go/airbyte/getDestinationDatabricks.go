// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// DestinationDatabricks DataSource
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
//			_, err := airbyte.LookupDestinationDatabricks(ctx, &airbyte.LookupDestinationDatabricksArgs{
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
func LookupDestinationDatabricks(ctx *pulumi.Context, args *LookupDestinationDatabricksArgs, opts ...pulumi.InvokeOption) (*LookupDestinationDatabricksResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDestinationDatabricksResult
	err := ctx.Invoke("airbyte:index/getDestinationDatabricks:getDestinationDatabricks", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDestinationDatabricks.
type LookupDestinationDatabricksArgs struct {
	DestinationId string `pulumi:"destinationId"`
}

// A collection of values returned by getDestinationDatabricks.
type LookupDestinationDatabricksResult struct {
	Configuration GetDestinationDatabricksConfiguration `pulumi:"configuration"`
	DestinationId string                                `pulumi:"destinationId"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupDestinationDatabricksOutput(ctx *pulumi.Context, args LookupDestinationDatabricksOutputArgs, opts ...pulumi.InvokeOption) LookupDestinationDatabricksResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDestinationDatabricksResult, error) {
			args := v.(LookupDestinationDatabricksArgs)
			r, err := LookupDestinationDatabricks(ctx, &args, opts...)
			var s LookupDestinationDatabricksResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDestinationDatabricksResultOutput)
}

// A collection of arguments for invoking getDestinationDatabricks.
type LookupDestinationDatabricksOutputArgs struct {
	DestinationId pulumi.StringInput `pulumi:"destinationId"`
}

func (LookupDestinationDatabricksOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDestinationDatabricksArgs)(nil)).Elem()
}

// A collection of values returned by getDestinationDatabricks.
type LookupDestinationDatabricksResultOutput struct{ *pulumi.OutputState }

func (LookupDestinationDatabricksResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDestinationDatabricksResult)(nil)).Elem()
}

func (o LookupDestinationDatabricksResultOutput) ToLookupDestinationDatabricksResultOutput() LookupDestinationDatabricksResultOutput {
	return o
}

func (o LookupDestinationDatabricksResultOutput) ToLookupDestinationDatabricksResultOutputWithContext(ctx context.Context) LookupDestinationDatabricksResultOutput {
	return o
}

func (o LookupDestinationDatabricksResultOutput) Configuration() GetDestinationDatabricksConfigurationOutput {
	return o.ApplyT(func(v LookupDestinationDatabricksResult) GetDestinationDatabricksConfiguration {
		return v.Configuration
	}).(GetDestinationDatabricksConfigurationOutput)
}

func (o LookupDestinationDatabricksResultOutput) DestinationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationDatabricksResult) string { return v.DestinationId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDestinationDatabricksResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationDatabricksResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDestinationDatabricksResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationDatabricksResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDestinationDatabricksResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationDatabricksResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDestinationDatabricksResultOutput{})
}
