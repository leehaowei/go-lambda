package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

type ResponseResult struct {
	Result string `json:"result"`
}

func HandleRequest(ctx context.Context, name MyEvent) (ResponseResult, error) {
	return ResponseResult{
		Result: fmt.Sprintf("Hello %s!", name.Name),
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
