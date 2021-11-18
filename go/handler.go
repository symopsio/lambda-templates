package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func handleEscalate(log *SymLogEntry) error {
	flow := ParsedSrn{Srn: log.Run.Flow}
	target := ParsedSrn{Srn: log.Fields.Target.Srn}

	if flow.Slug() != "my-lambda-flow" {
		return errors.New("Unknown flow")
	}

	// Call your business logic, passing, for example, the username and the target
	// doSomething(log.actor.username, target.srn.slug);
	fmt.Printf("Target: %+v", target)

	return nil
}

func handleDeescalate(log *SymLogEntry) error {
	target := ParsedSrn{Srn: log.Fields.Target.Srn}

	// Call your business logic, passing, for example, the username and the target
	// doSomething(log.actor.username, target.srn.slug);
	fmt.Printf("Target: %+v", target)

	return nil
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, log *SymLogEntry) error {
	if log.Event.Type == "escalate" {
		return handleEscalate(log)
	} else if log.Event.Type == "deescalate" {
		return handleDeescalate(log)
	} else {
		return errors.New("Unknown event type")
	}
}

func main() {
	lambda.Start(Handler)
}
