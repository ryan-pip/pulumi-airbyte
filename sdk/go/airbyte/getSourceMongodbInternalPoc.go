// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// SourceMongodbInternalPoc DataSource
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
//			_, err := airbyte.LookupSourceMongodbInternalPoc(ctx, &airbyte.LookupSourceMongodbInternalPocArgs{
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
func LookupSourceMongodbInternalPoc(ctx *pulumi.Context, args *LookupSourceMongodbInternalPocArgs, opts ...pulumi.InvokeOption) (*LookupSourceMongodbInternalPocResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceMongodbInternalPocResult
	err := ctx.Invoke("airbyte:index/getSourceMongodbInternalPoc:getSourceMongodbInternalPoc", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceMongodbInternalPoc.
type LookupSourceMongodbInternalPocArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceMongodbInternalPoc.
type LookupSourceMongodbInternalPocResult struct {
	Configuration GetSourceMongodbInternalPocConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceMongodbInternalPocOutput(ctx *pulumi.Context, args LookupSourceMongodbInternalPocOutputArgs, opts ...pulumi.InvokeOption) LookupSourceMongodbInternalPocResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceMongodbInternalPocResult, error) {
			args := v.(LookupSourceMongodbInternalPocArgs)
			r, err := LookupSourceMongodbInternalPoc(ctx, &args, opts...)
			var s LookupSourceMongodbInternalPocResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceMongodbInternalPocResultOutput)
}

// A collection of arguments for invoking getSourceMongodbInternalPoc.
type LookupSourceMongodbInternalPocOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceMongodbInternalPocOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceMongodbInternalPocArgs)(nil)).Elem()
}

// A collection of values returned by getSourceMongodbInternalPoc.
type LookupSourceMongodbInternalPocResultOutput struct{ *pulumi.OutputState }

func (LookupSourceMongodbInternalPocResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceMongodbInternalPocResult)(nil)).Elem()
}

func (o LookupSourceMongodbInternalPocResultOutput) ToLookupSourceMongodbInternalPocResultOutput() LookupSourceMongodbInternalPocResultOutput {
	return o
}

func (o LookupSourceMongodbInternalPocResultOutput) ToLookupSourceMongodbInternalPocResultOutputWithContext(ctx context.Context) LookupSourceMongodbInternalPocResultOutput {
	return o
}

func (o LookupSourceMongodbInternalPocResultOutput) Configuration() GetSourceMongodbInternalPocConfigurationOutput {
	return o.ApplyT(func(v LookupSourceMongodbInternalPocResult) GetSourceMongodbInternalPocConfiguration {
		return v.Configuration
	}).(GetSourceMongodbInternalPocConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceMongodbInternalPocResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceMongodbInternalPocResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceMongodbInternalPocResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceMongodbInternalPocResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceMongodbInternalPocResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceMongodbInternalPocResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceMongodbInternalPocResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceMongodbInternalPocResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceMongodbInternalPocResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceMongodbInternalPocResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceMongodbInternalPocResultOutput{})
}
