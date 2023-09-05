// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// DestinationAwsDatalake DataSource
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
//			_, err := airbyte.LookupDestinationAWSDatalake(ctx, &airbyte.LookupDestinationAWSDatalakeArgs{
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
func LookupDestinationAWSDatalake(ctx *pulumi.Context, args *LookupDestinationAWSDatalakeArgs, opts ...pulumi.InvokeOption) (*LookupDestinationAWSDatalakeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDestinationAWSDatalakeResult
	err := ctx.Invoke("airbyte:index/getDestinationAWSDatalake:getDestinationAWSDatalake", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDestinationAWSDatalake.
type LookupDestinationAWSDatalakeArgs struct {
	DestinationId string `pulumi:"destinationId"`
}

// A collection of values returned by getDestinationAWSDatalake.
type LookupDestinationAWSDatalakeResult struct {
	Configuration GetDestinationAWSDatalakeConfiguration `pulumi:"configuration"`
	DestinationId string                                 `pulumi:"destinationId"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupDestinationAWSDatalakeOutput(ctx *pulumi.Context, args LookupDestinationAWSDatalakeOutputArgs, opts ...pulumi.InvokeOption) LookupDestinationAWSDatalakeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDestinationAWSDatalakeResult, error) {
			args := v.(LookupDestinationAWSDatalakeArgs)
			r, err := LookupDestinationAWSDatalake(ctx, &args, opts...)
			var s LookupDestinationAWSDatalakeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDestinationAWSDatalakeResultOutput)
}

// A collection of arguments for invoking getDestinationAWSDatalake.
type LookupDestinationAWSDatalakeOutputArgs struct {
	DestinationId pulumi.StringInput `pulumi:"destinationId"`
}

func (LookupDestinationAWSDatalakeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDestinationAWSDatalakeArgs)(nil)).Elem()
}

// A collection of values returned by getDestinationAWSDatalake.
type LookupDestinationAWSDatalakeResultOutput struct{ *pulumi.OutputState }

func (LookupDestinationAWSDatalakeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDestinationAWSDatalakeResult)(nil)).Elem()
}

func (o LookupDestinationAWSDatalakeResultOutput) ToLookupDestinationAWSDatalakeResultOutput() LookupDestinationAWSDatalakeResultOutput {
	return o
}

func (o LookupDestinationAWSDatalakeResultOutput) ToLookupDestinationAWSDatalakeResultOutputWithContext(ctx context.Context) LookupDestinationAWSDatalakeResultOutput {
	return o
}

func (o LookupDestinationAWSDatalakeResultOutput) Configuration() GetDestinationAWSDatalakeConfigurationOutput {
	return o.ApplyT(func(v LookupDestinationAWSDatalakeResult) GetDestinationAWSDatalakeConfiguration {
		return v.Configuration
	}).(GetDestinationAWSDatalakeConfigurationOutput)
}

func (o LookupDestinationAWSDatalakeResultOutput) DestinationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationAWSDatalakeResult) string { return v.DestinationId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDestinationAWSDatalakeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationAWSDatalakeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDestinationAWSDatalakeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationAWSDatalakeResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDestinationAWSDatalakeResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationAWSDatalakeResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDestinationAWSDatalakeResultOutput{})
}
