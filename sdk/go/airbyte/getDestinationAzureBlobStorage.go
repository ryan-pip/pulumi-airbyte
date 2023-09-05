// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-airbyte/sdk/go/airbyte/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// DestinationAzureBlobStorage DataSource
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
//			_, err := airbyte.LookupDestinationAzureBlobStorage(ctx, &airbyte.LookupDestinationAzureBlobStorageArgs{
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
func LookupDestinationAzureBlobStorage(ctx *pulumi.Context, args *LookupDestinationAzureBlobStorageArgs, opts ...pulumi.InvokeOption) (*LookupDestinationAzureBlobStorageResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDestinationAzureBlobStorageResult
	err := ctx.Invoke("airbyte:index/getDestinationAzureBlobStorage:getDestinationAzureBlobStorage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDestinationAzureBlobStorage.
type LookupDestinationAzureBlobStorageArgs struct {
	DestinationId string `pulumi:"destinationId"`
}

// A collection of values returned by getDestinationAzureBlobStorage.
type LookupDestinationAzureBlobStorageResult struct {
	Configuration GetDestinationAzureBlobStorageConfiguration `pulumi:"configuration"`
	DestinationId string                                      `pulumi:"destinationId"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupDestinationAzureBlobStorageOutput(ctx *pulumi.Context, args LookupDestinationAzureBlobStorageOutputArgs, opts ...pulumi.InvokeOption) LookupDestinationAzureBlobStorageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDestinationAzureBlobStorageResult, error) {
			args := v.(LookupDestinationAzureBlobStorageArgs)
			r, err := LookupDestinationAzureBlobStorage(ctx, &args, opts...)
			var s LookupDestinationAzureBlobStorageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDestinationAzureBlobStorageResultOutput)
}

// A collection of arguments for invoking getDestinationAzureBlobStorage.
type LookupDestinationAzureBlobStorageOutputArgs struct {
	DestinationId pulumi.StringInput `pulumi:"destinationId"`
}

func (LookupDestinationAzureBlobStorageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDestinationAzureBlobStorageArgs)(nil)).Elem()
}

// A collection of values returned by getDestinationAzureBlobStorage.
type LookupDestinationAzureBlobStorageResultOutput struct{ *pulumi.OutputState }

func (LookupDestinationAzureBlobStorageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDestinationAzureBlobStorageResult)(nil)).Elem()
}

func (o LookupDestinationAzureBlobStorageResultOutput) ToLookupDestinationAzureBlobStorageResultOutput() LookupDestinationAzureBlobStorageResultOutput {
	return o
}

func (o LookupDestinationAzureBlobStorageResultOutput) ToLookupDestinationAzureBlobStorageResultOutputWithContext(ctx context.Context) LookupDestinationAzureBlobStorageResultOutput {
	return o
}

func (o LookupDestinationAzureBlobStorageResultOutput) Configuration() GetDestinationAzureBlobStorageConfigurationOutput {
	return o.ApplyT(func(v LookupDestinationAzureBlobStorageResult) GetDestinationAzureBlobStorageConfiguration {
		return v.Configuration
	}).(GetDestinationAzureBlobStorageConfigurationOutput)
}

func (o LookupDestinationAzureBlobStorageResultOutput) DestinationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationAzureBlobStorageResult) string { return v.DestinationId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDestinationAzureBlobStorageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationAzureBlobStorageResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDestinationAzureBlobStorageResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationAzureBlobStorageResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDestinationAzureBlobStorageResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationAzureBlobStorageResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDestinationAzureBlobStorageResultOutput{})
}
