// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"internal"
)

// SourceFacebookMarketing DataSource
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
// _, err := airbyte.LookupSourceFacebookMarketing(ctx, %!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference), nil);
// if err != nil {
// return err
// }
// return nil
// })
// }
// ```
func LookupSourceFacebookMarketing(ctx *pulumi.Context, args *LookupSourceFacebookMarketingArgs, opts ...pulumi.InvokeOption) (*LookupSourceFacebookMarketingResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceFacebookMarketingResult
	err := ctx.Invoke("airbyte:index/getSourceFacebookMarketing:getSourceFacebookMarketing", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceFacebookMarketing.
type LookupSourceFacebookMarketingArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceFacebookMarketing.
type LookupSourceFacebookMarketingResult struct {
	Configuration GetSourceFacebookMarketingConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceFacebookMarketingOutput(ctx *pulumi.Context, args LookupSourceFacebookMarketingOutputArgs, opts ...pulumi.InvokeOption) LookupSourceFacebookMarketingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceFacebookMarketingResult, error) {
			args := v.(LookupSourceFacebookMarketingArgs)
			r, err := LookupSourceFacebookMarketing(ctx, &args, opts...)
			var s LookupSourceFacebookMarketingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceFacebookMarketingResultOutput)
}

// A collection of arguments for invoking getSourceFacebookMarketing.
type LookupSourceFacebookMarketingOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceFacebookMarketingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceFacebookMarketingArgs)(nil)).Elem()
}

// A collection of values returned by getSourceFacebookMarketing.
type LookupSourceFacebookMarketingResultOutput struct{ *pulumi.OutputState }

func (LookupSourceFacebookMarketingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceFacebookMarketingResult)(nil)).Elem()
}

func (o LookupSourceFacebookMarketingResultOutput) ToLookupSourceFacebookMarketingResultOutput() LookupSourceFacebookMarketingResultOutput {
	return o
}

func (o LookupSourceFacebookMarketingResultOutput) ToLookupSourceFacebookMarketingResultOutputWithContext(ctx context.Context) LookupSourceFacebookMarketingResultOutput {
	return o
}

func (o LookupSourceFacebookMarketingResultOutput) Configuration() GetSourceFacebookMarketingConfigurationOutput {
	return o.ApplyT(func(v LookupSourceFacebookMarketingResult) GetSourceFacebookMarketingConfiguration {
		return v.Configuration
	}).(GetSourceFacebookMarketingConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceFacebookMarketingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceFacebookMarketingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceFacebookMarketingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceFacebookMarketingResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceFacebookMarketingResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceFacebookMarketingResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceFacebookMarketingResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceFacebookMarketingResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceFacebookMarketingResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceFacebookMarketingResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceFacebookMarketingResultOutput{})
}
