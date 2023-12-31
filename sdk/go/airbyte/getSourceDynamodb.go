// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// SourceDynamodb DataSource
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
//			_, err := airbyte.LookupSourceDynamodb(ctx, &airbyte.LookupSourceDynamodbArgs{
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
func LookupSourceDynamodb(ctx *pulumi.Context, args *LookupSourceDynamodbArgs, opts ...pulumi.InvokeOption) (*LookupSourceDynamodbResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceDynamodbResult
	err := ctx.Invoke("airbyte:index/getSourceDynamodb:getSourceDynamodb", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceDynamodb.
type LookupSourceDynamodbArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceDynamodb.
type LookupSourceDynamodbResult struct {
	Configuration GetSourceDynamodbConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceDynamodbOutput(ctx *pulumi.Context, args LookupSourceDynamodbOutputArgs, opts ...pulumi.InvokeOption) LookupSourceDynamodbResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceDynamodbResult, error) {
			args := v.(LookupSourceDynamodbArgs)
			r, err := LookupSourceDynamodb(ctx, &args, opts...)
			var s LookupSourceDynamodbResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceDynamodbResultOutput)
}

// A collection of arguments for invoking getSourceDynamodb.
type LookupSourceDynamodbOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceDynamodbOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceDynamodbArgs)(nil)).Elem()
}

// A collection of values returned by getSourceDynamodb.
type LookupSourceDynamodbResultOutput struct{ *pulumi.OutputState }

func (LookupSourceDynamodbResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceDynamodbResult)(nil)).Elem()
}

func (o LookupSourceDynamodbResultOutput) ToLookupSourceDynamodbResultOutput() LookupSourceDynamodbResultOutput {
	return o
}

func (o LookupSourceDynamodbResultOutput) ToLookupSourceDynamodbResultOutputWithContext(ctx context.Context) LookupSourceDynamodbResultOutput {
	return o
}

func (o LookupSourceDynamodbResultOutput) Configuration() GetSourceDynamodbConfigurationOutput {
	return o.ApplyT(func(v LookupSourceDynamodbResult) GetSourceDynamodbConfiguration { return v.Configuration }).(GetSourceDynamodbConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceDynamodbResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceDynamodbResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceDynamodbResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceDynamodbResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceDynamodbResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceDynamodbResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceDynamodbResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceDynamodbResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceDynamodbResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceDynamodbResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceDynamodbResultOutput{})
}
