// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"internal"
)

// SourcePipedrive DataSource
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
// _, err := airbyte.LookupSourcePipedrive(ctx, %!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference), nil);
// if err != nil {
// return err
// }
// return nil
// })
// }
// ```
func LookupSourcePipedrive(ctx *pulumi.Context, args *LookupSourcePipedriveArgs, opts ...pulumi.InvokeOption) (*LookupSourcePipedriveResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourcePipedriveResult
	err := ctx.Invoke("airbyte:index/getSourcePipedrive:getSourcePipedrive", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourcePipedrive.
type LookupSourcePipedriveArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourcePipedrive.
type LookupSourcePipedriveResult struct {
	Configuration GetSourcePipedriveConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourcePipedriveOutput(ctx *pulumi.Context, args LookupSourcePipedriveOutputArgs, opts ...pulumi.InvokeOption) LookupSourcePipedriveResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourcePipedriveResult, error) {
			args := v.(LookupSourcePipedriveArgs)
			r, err := LookupSourcePipedrive(ctx, &args, opts...)
			var s LookupSourcePipedriveResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourcePipedriveResultOutput)
}

// A collection of arguments for invoking getSourcePipedrive.
type LookupSourcePipedriveOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourcePipedriveOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourcePipedriveArgs)(nil)).Elem()
}

// A collection of values returned by getSourcePipedrive.
type LookupSourcePipedriveResultOutput struct{ *pulumi.OutputState }

func (LookupSourcePipedriveResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourcePipedriveResult)(nil)).Elem()
}

func (o LookupSourcePipedriveResultOutput) ToLookupSourcePipedriveResultOutput() LookupSourcePipedriveResultOutput {
	return o
}

func (o LookupSourcePipedriveResultOutput) ToLookupSourcePipedriveResultOutputWithContext(ctx context.Context) LookupSourcePipedriveResultOutput {
	return o
}

func (o LookupSourcePipedriveResultOutput) Configuration() GetSourcePipedriveConfigurationOutput {
	return o.ApplyT(func(v LookupSourcePipedriveResult) GetSourcePipedriveConfiguration { return v.Configuration }).(GetSourcePipedriveConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourcePipedriveResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourcePipedriveResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourcePipedriveResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourcePipedriveResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourcePipedriveResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourcePipedriveResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourcePipedriveResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourcePipedriveResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourcePipedriveResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourcePipedriveResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourcePipedriveResultOutput{})
}
