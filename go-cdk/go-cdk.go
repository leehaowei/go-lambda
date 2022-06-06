package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type GoCdkStackProps struct {
	awscdk.StackProps
}

func NewGoCdkStack(scope constructs.Construct, id string, props *GoCdkStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// The code that defines your stack goes here
	awslambda.NewFunction(stack, jsii.String("KomaGoLambda"), &awslambda.FunctionProps{
		Code:    awslambda.NewAssetCode(jsii.String("lambda"), nil),
		Handler: jsii.String("hello.handler"),
		Timeout: awscdk.Duration_Seconds(jsii.Number(15)),
		Runtime: awslambda.Runtime_NODEJS_14_X(),
	})

	return stack
}

func main() {
	app := awscdk.NewApp(nil)

	NewGoCdkStack(app, "GoCdkStack", &GoCdkStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return &awscdk.Environment{
		// Account: jsii.String("123456789012"),
		Region: jsii.String("eu-central-1"),
	}
}
