// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// SourceZendeskSunshine DataSource
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
//			_, err := airbyte.LookupSourceZendeskSunshine(ctx, &airbyte.LookupSourceZendeskSunshineArgs{
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
func LookupSourceZendeskSunshine(ctx *pulumi.Context, args *LookupSourceZendeskSunshineArgs, opts ...pulumi.InvokeOption) (*LookupSourceZendeskSunshineResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceZendeskSunshineResult
	err := ctx.Invoke("airbyte:index/getSourceZendeskSunshine:getSourceZendeskSunshine", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceZendeskSunshine.
type LookupSourceZendeskSunshineArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceZendeskSunshine.
type LookupSourceZendeskSunshineResult struct {
	Configuration GetSourceZendeskSunshineConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceZendeskSunshineOutput(ctx *pulumi.Context, args LookupSourceZendeskSunshineOutputArgs, opts ...pulumi.InvokeOption) LookupSourceZendeskSunshineResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceZendeskSunshineResult, error) {
			args := v.(LookupSourceZendeskSunshineArgs)
			r, err := LookupSourceZendeskSunshine(ctx, &args, opts...)
			var s LookupSourceZendeskSunshineResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceZendeskSunshineResultOutput)
}

// A collection of arguments for invoking getSourceZendeskSunshine.
type LookupSourceZendeskSunshineOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceZendeskSunshineOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceZendeskSunshineArgs)(nil)).Elem()
}

// A collection of values returned by getSourceZendeskSunshine.
type LookupSourceZendeskSunshineResultOutput struct{ *pulumi.OutputState }

func (LookupSourceZendeskSunshineResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceZendeskSunshineResult)(nil)).Elem()
}

func (o LookupSourceZendeskSunshineResultOutput) ToLookupSourceZendeskSunshineResultOutput() LookupSourceZendeskSunshineResultOutput {
	return o
}

func (o LookupSourceZendeskSunshineResultOutput) ToLookupSourceZendeskSunshineResultOutputWithContext(ctx context.Context) LookupSourceZendeskSunshineResultOutput {
	return o
}

func (o LookupSourceZendeskSunshineResultOutput) Configuration() GetSourceZendeskSunshineConfigurationOutput {
	return o.ApplyT(func(v LookupSourceZendeskSunshineResult) GetSourceZendeskSunshineConfiguration {
		return v.Configuration
	}).(GetSourceZendeskSunshineConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceZendeskSunshineResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceZendeskSunshineResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceZendeskSunshineResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceZendeskSunshineResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceZendeskSunshineResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceZendeskSunshineResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceZendeskSunshineResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceZendeskSunshineResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceZendeskSunshineResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceZendeskSunshineResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceZendeskSunshineResultOutput{})
}
