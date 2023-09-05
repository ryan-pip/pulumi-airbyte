// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"internal"
)

// SourceMicrosoftTeams DataSource
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
// _, err := airbyte.LookupSourceMicrosoftTeams(ctx, %!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference), nil);
// if err != nil {
// return err
// }
// return nil
// })
// }
// ```
func LookupSourceMicrosoftTeams(ctx *pulumi.Context, args *LookupSourceMicrosoftTeamsArgs, opts ...pulumi.InvokeOption) (*LookupSourceMicrosoftTeamsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceMicrosoftTeamsResult
	err := ctx.Invoke("airbyte:index/getSourceMicrosoftTeams:getSourceMicrosoftTeams", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceMicrosoftTeams.
type LookupSourceMicrosoftTeamsArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId *string `pulumi:"secretId"`
	SourceId string  `pulumi:"sourceId"`
}

// A collection of values returned by getSourceMicrosoftTeams.
type LookupSourceMicrosoftTeamsResult struct {
	Configuration GetSourceMicrosoftTeamsConfiguration `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    string  `pulumi:"sourceId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

func LookupSourceMicrosoftTeamsOutput(ctx *pulumi.Context, args LookupSourceMicrosoftTeamsOutputArgs, opts ...pulumi.InvokeOption) LookupSourceMicrosoftTeamsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceMicrosoftTeamsResult, error) {
			args := v.(LookupSourceMicrosoftTeamsArgs)
			r, err := LookupSourceMicrosoftTeams(ctx, &args, opts...)
			var s LookupSourceMicrosoftTeamsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceMicrosoftTeamsResultOutput)
}

// A collection of arguments for invoking getSourceMicrosoftTeams.
type LookupSourceMicrosoftTeamsOutputArgs struct {
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceMicrosoftTeamsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceMicrosoftTeamsArgs)(nil)).Elem()
}

// A collection of values returned by getSourceMicrosoftTeams.
type LookupSourceMicrosoftTeamsResultOutput struct{ *pulumi.OutputState }

func (LookupSourceMicrosoftTeamsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceMicrosoftTeamsResult)(nil)).Elem()
}

func (o LookupSourceMicrosoftTeamsResultOutput) ToLookupSourceMicrosoftTeamsResultOutput() LookupSourceMicrosoftTeamsResultOutput {
	return o
}

func (o LookupSourceMicrosoftTeamsResultOutput) ToLookupSourceMicrosoftTeamsResultOutputWithContext(ctx context.Context) LookupSourceMicrosoftTeamsResultOutput {
	return o
}

func (o LookupSourceMicrosoftTeamsResultOutput) Configuration() GetSourceMicrosoftTeamsConfigurationOutput {
	return o.ApplyT(func(v LookupSourceMicrosoftTeamsResult) GetSourceMicrosoftTeamsConfiguration { return v.Configuration }).(GetSourceMicrosoftTeamsConfigurationOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceMicrosoftTeamsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceMicrosoftTeamsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceMicrosoftTeamsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceMicrosoftTeamsResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o LookupSourceMicrosoftTeamsResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceMicrosoftTeamsResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSourceMicrosoftTeamsResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceMicrosoftTeamsResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceMicrosoftTeamsResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceMicrosoftTeamsResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceMicrosoftTeamsResultOutput{})
}
