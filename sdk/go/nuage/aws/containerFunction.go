// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ContainerFunction struct {
	pulumi.ResourceState

	Arn          pulumi.StringOutput `pulumi:"arn"`
	Function_url pulumi.StringOutput `pulumi:"function_url"`
	Name         pulumi.StringOutput `pulumi:"name"`
}

// NewContainerFunction registers a new resource with the given unique name, arguments, and options.
func NewContainerFunction(ctx *pulumi.Context,
	name string, args *ContainerFunctionArgs, opts ...pulumi.ResourceOption) (*ContainerFunction, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EcrRepositoryName == nil {
		return nil, errors.New("invalid value for required argument 'EcrRepositoryName'")
	}
	var resource ContainerFunction
	err := ctx.RegisterRemoteComponentResource("nuage:aws:ContainerFunction", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type containerFunctionArgs struct {
	// Architecture, either `X86_64` or `ARM64`. Defaults to `x86_64`
	Architecture *string `pulumi:"architecture"`
	// Dockerfile context path.
	Context *string `pulumi:"context"`
	// Description of the function.
	Description *string `pulumi:"description"`
	// Dockerfile path. Defaults to `./Dockerfile`
	Dockerfile *string `pulumi:"dockerfile"`
	// ECR repository name for new definition.
	EcrRepositoryName string `pulumi:"ecrRepositoryName"`
	// Environment Variables
	Environment map[string]string `pulumi:"environment"`
	// Keep warm by refreshing the lambda function every 5 minutes. Defaults to `false`
	KeepWarm *bool `pulumi:"keepWarm"`
	// Amount of memory in MB your Lambda Function can use at runtime. Defaults to `512`.
	MemorySize *float64 `pulumi:"memorySize"`
	// Policy Document for lambda.
	PolicyDocument *string `pulumi:"policyDocument"`
	// Existing ECR repository name
	Repository *string `pulumi:"repository"`
	// Amount of time your Lambda Function has to run in seconds. Defaults to `3`
	Timeout *float64 `pulumi:"timeout"`
	// Use Lambda URL. Defaults to `false`
	Url *bool `pulumi:"url"`
}

// The set of arguments for constructing a ContainerFunction resource.
type ContainerFunctionArgs struct {
	// Architecture, either `X86_64` or `ARM64`. Defaults to `x86_64`
	Architecture pulumi.StringPtrInput
	// Dockerfile context path.
	Context pulumi.StringPtrInput
	// Description of the function.
	Description pulumi.StringPtrInput
	// Dockerfile path. Defaults to `./Dockerfile`
	Dockerfile pulumi.StringPtrInput
	// ECR repository name for new definition.
	EcrRepositoryName pulumi.StringInput
	// Environment Variables
	Environment pulumi.StringMapInput
	// Keep warm by refreshing the lambda function every 5 minutes. Defaults to `false`
	KeepWarm pulumi.BoolPtrInput
	// Amount of memory in MB your Lambda Function can use at runtime. Defaults to `512`.
	MemorySize pulumi.Float64PtrInput
	// Policy Document for lambda.
	PolicyDocument pulumi.StringPtrInput
	// Existing ECR repository name
	Repository pulumi.StringPtrInput
	// Amount of time your Lambda Function has to run in seconds. Defaults to `3`
	Timeout pulumi.Float64PtrInput
	// Use Lambda URL. Defaults to `false`
	Url pulumi.BoolPtrInput
}

func (ContainerFunctionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerFunctionArgs)(nil)).Elem()
}

type ContainerFunctionInput interface {
	pulumi.Input

	ToContainerFunctionOutput() ContainerFunctionOutput
	ToContainerFunctionOutputWithContext(ctx context.Context) ContainerFunctionOutput
}

func (*ContainerFunction) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerFunction)(nil)).Elem()
}

func (i *ContainerFunction) ToContainerFunctionOutput() ContainerFunctionOutput {
	return i.ToContainerFunctionOutputWithContext(context.Background())
}

func (i *ContainerFunction) ToContainerFunctionOutputWithContext(ctx context.Context) ContainerFunctionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerFunctionOutput)
}

// ContainerFunctionArrayInput is an input type that accepts ContainerFunctionArray and ContainerFunctionArrayOutput values.
// You can construct a concrete instance of `ContainerFunctionArrayInput` via:
//
//	ContainerFunctionArray{ ContainerFunctionArgs{...} }
type ContainerFunctionArrayInput interface {
	pulumi.Input

	ToContainerFunctionArrayOutput() ContainerFunctionArrayOutput
	ToContainerFunctionArrayOutputWithContext(context.Context) ContainerFunctionArrayOutput
}

type ContainerFunctionArray []ContainerFunctionInput

func (ContainerFunctionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContainerFunction)(nil)).Elem()
}

func (i ContainerFunctionArray) ToContainerFunctionArrayOutput() ContainerFunctionArrayOutput {
	return i.ToContainerFunctionArrayOutputWithContext(context.Background())
}

func (i ContainerFunctionArray) ToContainerFunctionArrayOutputWithContext(ctx context.Context) ContainerFunctionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerFunctionArrayOutput)
}

// ContainerFunctionMapInput is an input type that accepts ContainerFunctionMap and ContainerFunctionMapOutput values.
// You can construct a concrete instance of `ContainerFunctionMapInput` via:
//
//	ContainerFunctionMap{ "key": ContainerFunctionArgs{...} }
type ContainerFunctionMapInput interface {
	pulumi.Input

	ToContainerFunctionMapOutput() ContainerFunctionMapOutput
	ToContainerFunctionMapOutputWithContext(context.Context) ContainerFunctionMapOutput
}

type ContainerFunctionMap map[string]ContainerFunctionInput

func (ContainerFunctionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContainerFunction)(nil)).Elem()
}

func (i ContainerFunctionMap) ToContainerFunctionMapOutput() ContainerFunctionMapOutput {
	return i.ToContainerFunctionMapOutputWithContext(context.Background())
}

func (i ContainerFunctionMap) ToContainerFunctionMapOutputWithContext(ctx context.Context) ContainerFunctionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerFunctionMapOutput)
}

type ContainerFunctionOutput struct{ *pulumi.OutputState }

func (ContainerFunctionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerFunction)(nil)).Elem()
}

func (o ContainerFunctionOutput) ToContainerFunctionOutput() ContainerFunctionOutput {
	return o
}

func (o ContainerFunctionOutput) ToContainerFunctionOutputWithContext(ctx context.Context) ContainerFunctionOutput {
	return o
}

type ContainerFunctionArrayOutput struct{ *pulumi.OutputState }

func (ContainerFunctionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContainerFunction)(nil)).Elem()
}

func (o ContainerFunctionArrayOutput) ToContainerFunctionArrayOutput() ContainerFunctionArrayOutput {
	return o
}

func (o ContainerFunctionArrayOutput) ToContainerFunctionArrayOutputWithContext(ctx context.Context) ContainerFunctionArrayOutput {
	return o
}

func (o ContainerFunctionArrayOutput) Index(i pulumi.IntInput) ContainerFunctionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ContainerFunction {
		return vs[0].([]*ContainerFunction)[vs[1].(int)]
	}).(ContainerFunctionOutput)
}

type ContainerFunctionMapOutput struct{ *pulumi.OutputState }

func (ContainerFunctionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContainerFunction)(nil)).Elem()
}

func (o ContainerFunctionMapOutput) ToContainerFunctionMapOutput() ContainerFunctionMapOutput {
	return o
}

func (o ContainerFunctionMapOutput) ToContainerFunctionMapOutputWithContext(ctx context.Context) ContainerFunctionMapOutput {
	return o
}

func (o ContainerFunctionMapOutput) MapIndex(k pulumi.StringInput) ContainerFunctionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ContainerFunction {
		return vs[0].(map[string]*ContainerFunction)[vs[1].(string)]
	}).(ContainerFunctionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerFunctionInput)(nil)).Elem(), &ContainerFunction{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerFunctionArrayInput)(nil)).Elem(), ContainerFunctionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerFunctionMapInput)(nil)).Elem(), ContainerFunctionMap{})
	pulumi.RegisterOutputType(ContainerFunctionOutput{})
	pulumi.RegisterOutputType(ContainerFunctionArrayOutput{})
	pulumi.RegisterOutputType(ContainerFunctionMapOutput{})
}