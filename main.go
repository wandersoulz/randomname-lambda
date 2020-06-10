package main

import (
	"context"
	"math/rand"
	"os"
	"strconv"
	"time"

	runtime "github.com/aws/aws-lambda-go/lambda"
	"github.com/wandersoulz/godes"
	"github.com/wandersoulz/randomname"
)

func generateRandomSeed() int64 {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 1337227
	return (int64)(rand.Intn(max-min+1) + min)
}

func init() {
	godes.SetSeed(generateRandomSeed())
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
