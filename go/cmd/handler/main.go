package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	i "github.com/symopsio/lambda-templates/go/internal"
)

func HandleRequest(ctx context.Context, event i.SymEvent) (*i.SymResult, error) {
	return i.HandleSymEvent(event)
}

func main() {
	lambda.Start(HandleRequest)
}
