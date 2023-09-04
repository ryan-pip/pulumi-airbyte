// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SourceCoinmarketcap struct {
	pulumi.CustomResourceState

	Configuration SourceCoinmarketcapConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput                    `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceCoinmarketcap registers a new resource with the given unique name, arguments, and options.
func NewSourceCoinmarketcap(ctx *pulumi.Context,
	name string, args *SourceCoinmarketcapArgs, opts ...pulumi.ResourceOption) (*SourceCoinmarketcap, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Configuration == nil {
		return nil, errors.New("invalid value for required argument 'Configuration'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.WorkspaceId == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SourceCoinmarketcap
	err := ctx.RegisterResource("airbyte:index/sourceCoinmarketcap:SourceCoinmarketcap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceCoinmarketcap gets an existing SourceCoinmarketcap resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceCoinmarketcap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceCoinmarketcapState, opts ...pulumi.ResourceOption) (*SourceCoinmarketcap, error) {
	var resource SourceCoinmarketcap
	err := ctx.ReadResource("airbyte:index/sourceCoinmarketcap:SourceCoinmarketcap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceCoinmarketcap resources.
type sourceCoinmarketcapState struct {
	Configuration *SourceCoinmarketcapConfiguration `pulumi:"configuration"`
	Name          *string                           `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceCoinmarketcapState struct {
	Configuration SourceCoinmarketcapConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceCoinmarketcapState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceCoinmarketcapState)(nil)).Elem()
}

type sourceCoinmarketcapArgs struct {
	Configuration SourceCoinmarketcapConfiguration `pulumi:"configuration"`
	Name          string                           `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceCoinmarketcap resource.
type SourceCoinmarketcapArgs struct {
	Configuration SourceCoinmarketcapConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceCoinmarketcapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceCoinmarketcapArgs)(nil)).Elem()
}

type SourceCoinmarketcapInput interface {
	pulumi.Input

	ToSourceCoinmarketcapOutput() SourceCoinmarketcapOutput
	ToSourceCoinmarketcapOutputWithContext(ctx context.Context) SourceCoinmarketcapOutput
}

func (*SourceCoinmarketcap) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceCoinmarketcap)(nil)).Elem()
}

func (i *SourceCoinmarketcap) ToSourceCoinmarketcapOutput() SourceCoinmarketcapOutput {
	return i.ToSourceCoinmarketcapOutputWithContext(context.Background())
}

func (i *SourceCoinmarketcap) ToSourceCoinmarketcapOutputWithContext(ctx context.Context) SourceCoinmarketcapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceCoinmarketcapOutput)
}

type SourceCoinmarketcapOutput struct{ *pulumi.OutputState }

func (SourceCoinmarketcapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceCoinmarketcap)(nil)).Elem()
}

func (o SourceCoinmarketcapOutput) ToSourceCoinmarketcapOutput() SourceCoinmarketcapOutput {
	return o
}

func (o SourceCoinmarketcapOutput) ToSourceCoinmarketcapOutputWithContext(ctx context.Context) SourceCoinmarketcapOutput {
	return o
}

func (o SourceCoinmarketcapOutput) Configuration() SourceCoinmarketcapConfigurationOutput {
	return o.ApplyT(func(v *SourceCoinmarketcap) SourceCoinmarketcapConfigurationOutput { return v.Configuration }).(SourceCoinmarketcapConfigurationOutput)
}

func (o SourceCoinmarketcapOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceCoinmarketcap) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceCoinmarketcapOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceCoinmarketcap) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceCoinmarketcapOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceCoinmarketcap) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceCoinmarketcapOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceCoinmarketcap) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceCoinmarketcapOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceCoinmarketcap) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceCoinmarketcapInput)(nil)).Elem(), &SourceCoinmarketcap{})
	pulumi.RegisterOutputType(SourceCoinmarketcapOutput{})
}
