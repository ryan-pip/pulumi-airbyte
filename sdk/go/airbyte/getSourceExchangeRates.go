// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"internal"
)

// SourceExchangeRates DataSource
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
// _, err := airbyte.LookupSourceExchangeRates(ctx, %!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference), nil);
// if err != nil {
// return err
// }
// return nil
// })
// }
// ```
func LookupSourceExchangeRates(ctx *pulumi.Context, args *LookupSourceExchangeRatesArgs, opts ...pulumi.InvokeOption) (*LookupSourceExchangeRatesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceExchangeRatesResult
	err := ctx.Invoke("airbyte:index/getSourceExchangeRates:getSourceExchangeRates", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceExchangeRates.
type LookupSourceExchangeRatesArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceExchangeRates.
type LookupSourceExchangeRatesResult struct {
	Configuration GetSourceExchangeRatesConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceExchangeRatesOutput(ctx *pulumi.Context, args LookupSourceExchangeRatesOutputArgs, opts ...pulumi.InvokeOption) LookupSourceExchangeRatesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceExchangeRatesResult, error) {
			args := v.(LookupSourceExchangeRatesArgs)
			r, err := LookupSourceExchangeRates(ctx, &args, opts...)
			var s LookupSourceExchangeRatesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceExchangeRatesResultOutput)
}

// A collection of arguments for invoking getSourceExchangeRates.
type LookupSourceExchangeRatesOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceExchangeRatesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceExchangeRatesArgs)(nil)).Elem()
}

// A collection of values returned by getSourceExchangeRates.
type LookupSourceExchangeRatesResultOutput struct{ *pulumi.OutputState }

func (LookupSourceExchangeRatesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceExchangeRatesResult)(nil)).Elem()
}

func (o LookupSourceExchangeRatesResultOutput) ToLookupSourceExchangeRatesResultOutput() LookupSourceExchangeRatesResultOutput {
	return o
}

func (o LookupSourceExchangeRatesResultOutput) ToLookupSourceExchangeRatesResultOutputWithContext(ctx context.Context) LookupSourceExchangeRatesResultOutput {
	return o
}

func (o LookupSourceExchangeRatesResultOutput) Configuration() GetSourceExchangeRatesConfigurationOutput {
	return o.ApplyT(func(v LookupSourceExchangeRatesResult) GetSourceExchangeRatesConfiguration { return v.Configuration }).(GetSourceExchangeRatesConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceExchangeRatesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceExchangeRatesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceExchangeRatesResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceExchangeRatesResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceExchangeRatesResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceExchangeRatesResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceExchangeRatesResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceExchangeRatesResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceExchangeRatesResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceExchangeRatesResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceExchangeRatesResultOutput{})
}
