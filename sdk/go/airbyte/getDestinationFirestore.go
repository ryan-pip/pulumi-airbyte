// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"internal"
)

// DestinationFirestore DataSource
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
// _, err := airbyte.LookupDestinationFirestore(ctx, %!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference), nil);
// if err != nil {
// return err
// }
// return nil
// })
// }
// ```
func LookupDestinationFirestore(ctx *pulumi.Context, args *LookupDestinationFirestoreArgs, opts ...pulumi.InvokeOption) (*LookupDestinationFirestoreResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDestinationFirestoreResult
	err := ctx.Invoke("airbyte:index/getDestinationFirestore:getDestinationFirestore", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDestinationFirestore.
type LookupDestinationFirestoreArgs struct {
	DestinationId string `pulumi:"destinationId"`
}

// A collection of values returned by getDestinationFirestore.
type LookupDestinationFirestoreResult struct {
	Configuration GetDestinationFirestoreConfiguration `pulumi:"configuration"`
	DestinationId string                               `pulumi:"destinationId"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupDestinationFirestoreOutput(ctx *pulumi.Context, args LookupDestinationFirestoreOutputArgs, opts ...pulumi.InvokeOption) LookupDestinationFirestoreResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDestinationFirestoreResult, error) {
			args := v.(LookupDestinationFirestoreArgs)
			r, err := LookupDestinationFirestore(ctx, &args, opts...)
			var s LookupDestinationFirestoreResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDestinationFirestoreResultOutput)
}

// A collection of arguments for invoking getDestinationFirestore.
type LookupDestinationFirestoreOutputArgs struct {
	DestinationId pulumi.StringInput `pulumi:"destinationId"`
}

func (LookupDestinationFirestoreOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDestinationFirestoreArgs)(nil)).Elem()
}

// A collection of values returned by getDestinationFirestore.
type LookupDestinationFirestoreResultOutput struct{ *pulumi.OutputState }

func (LookupDestinationFirestoreResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDestinationFirestoreResult)(nil)).Elem()
}

func (o LookupDestinationFirestoreResultOutput) ToLookupDestinationFirestoreResultOutput() LookupDestinationFirestoreResultOutput {
	return o
}

func (o LookupDestinationFirestoreResultOutput) ToLookupDestinationFirestoreResultOutputWithContext(ctx context.Context) LookupDestinationFirestoreResultOutput {
	return o
}

func (o LookupDestinationFirestoreResultOutput) Configuration() GetDestinationFirestoreConfigurationOutput {
	return o.ApplyT(func(v LookupDestinationFirestoreResult) GetDestinationFirestoreConfiguration { return v.Configuration }).(GetDestinationFirestoreConfigurationOutput)
}

func (o LookupDestinationFirestoreResultOutput) DestinationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationFirestoreResult) string { return v.DestinationId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDestinationFirestoreResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationFirestoreResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDestinationFirestoreResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationFirestoreResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDestinationFirestoreResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationFirestoreResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDestinationFirestoreResultOutput{})
}
