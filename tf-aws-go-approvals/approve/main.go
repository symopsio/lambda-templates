package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/symopsio/sym-approval-lambda/internal"

	"github.com/symopsio/types/go/enums"
	"github.com/symopsio/types/go/messages"
)

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, approval *messages.Approval) (*messages.ApprovalResponse, error) {
	identity, err := internal.FindIdentity(approval.Request.Target.User, enums.Service_OKTA)
	if err != nil {
		return &messages.ApprovalResponse{
			Ok: false, Error: err.Error(),
		}, nil
	}
	fmt.Printf("Target user slack id: %s\n", identity.Id)
	fmt.Printf("Target resource: %s\n", approval.Request.Target.Resource)
	fmt.Printf("Target reason: %s\n", approval.Request.Meta.Reason)

	return &messages.ApprovalResponse{Ok: true}, nil
}

func main() {
	lambda.Start(Handler)
}
