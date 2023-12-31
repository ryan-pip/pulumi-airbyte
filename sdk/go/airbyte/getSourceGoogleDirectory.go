// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// SourceGoogleDirectory DataSource
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
//			_, err := airbyte.LookupSourceGoogleDirectory(ctx, &airbyte.LookupSourceGoogleDirectoryArgs{
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
func LookupSourceGoogleDirectory(ctx *pulumi.Context, args *LookupSourceGoogleDirectoryArgs, opts ...pulumi.InvokeOption) (*LookupSourceGoogleDirectoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceGoogleDirectoryResult
	err := ctx.Invoke("airbyte:index/getSourceGoogleDirectory:getSourceGoogleDirectory", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceGoogleDirectory.
type LookupSourceGoogleDirectoryArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceGoogleDirectory.
type LookupSourceGoogleDirectoryResult struct {
	Configuration GetSourceGoogleDirectoryConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceGoogleDirectoryOutput(ctx *pulumi.Context, args LookupSourceGoogleDirectoryOutputArgs, opts ...pulumi.InvokeOption) LookupSourceGoogleDirectoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceGoogleDirectoryResult, error) {
			args := v.(LookupSourceGoogleDirectoryArgs)
			r, err := LookupSourceGoogleDirectory(ctx, &args, opts...)
			var s LookupSourceGoogleDirectoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceGoogleDirectoryResultOutput)
}

// A collection of arguments for invoking getSourceGoogleDirectory.
type LookupSourceGoogleDirectoryOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceGoogleDirectoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceGoogleDirectoryArgs)(nil)).Elem()
}

// A collection of values returned by getSourceGoogleDirectory.
type LookupSourceGoogleDirectoryResultOutput struct{ *pulumi.OutputState }

func (LookupSourceGoogleDirectoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceGoogleDirectoryResult)(nil)).Elem()
}

func (o LookupSourceGoogleDirectoryResultOutput) ToLookupSourceGoogleDirectoryResultOutput() LookupSourceGoogleDirectoryResultOutput {
	return o
}

func (o LookupSourceGoogleDirectoryResultOutput) ToLookupSourceGoogleDirectoryResultOutputWithContext(ctx context.Context) LookupSourceGoogleDirectoryResultOutput {
	return o
}

func (o LookupSourceGoogleDirectoryResultOutput) Configuration() GetSourceGoogleDirectoryConfigurationOutput {
	return o.ApplyT(func(v LookupSourceGoogleDirectoryResult) GetSourceGoogleDirectoryConfiguration {
		return v.Configuration
	}).(GetSourceGoogleDirectoryConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceGoogleDirectoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceGoogleDirectoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceGoogleDirectoryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceGoogleDirectoryResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceGoogleDirectoryResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceGoogleDirectoryResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceGoogleDirectoryResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceGoogleDirectoryResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceGoogleDirectoryResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceGoogleDirectoryResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceGoogleDirectoryResultOutput{})
}
