package main

import (
	"context"
	"os"
	"strconv"

	runtime "github.com/aws/aws-lambda-go/lambda"
	"github.com/wandersoulz/randomname"
)

func init() {
	contextSizeStr := os.Getenv("CONTEXT_SIZE")
	contextSize, _ := strconv.Atoi(contextSizeStr)
	randomname.Init("first-names.txt", contextSize)
}

func handleRequest(ctx context.Context) (string, error) {
	return randomname.GetName(), nil
}

func main() {
	runtime.Start(handleRequest)
}
