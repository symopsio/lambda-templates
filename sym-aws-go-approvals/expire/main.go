package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/serverless-templates/sym-aws-go-approvals/internal"
	"github.com/symopsio/types/go/enums"
	"github.com/symopsio/types/go/messages"
)

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, expiration *messages.Expiration) (*messages.ExpirationResponse, error) {
	identity, err := internal.FindIdentity(expiration.Target.User, enums.Service_SLACK)
	if err != nil {
		return &messages.ExpirationResponse{
			Ok: false, Error: err.Error(),
		}, nil
	}
	fmt.Printf("Target user slack id: %s\n", identity.Id)
	fmt.Printf("Target resource: %s\n", expiration.Target.Resource)

	return &messages.ExpirationResponse{Ok: true}, nil
}

func main() {
	lambda.Start(Handler)
}
