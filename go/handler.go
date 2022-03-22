package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type SymLambdaResponse struct {
	Body   map[string]interface{} `json:"body"`
	Errors []string               `json:"errors"`
}

func handleEscalate(log *SymLogEntry) SymLambdaResponse {
	flow := ParsedSrn{Srn: log.Run.Flow}
	target := ParsedSrn{Srn: log.Fields.Target.Srn}

	if flow.Slug() != "my-lambda-flow" {
		return SymLambdaResponse{
			Body:   make(map[string]interface{}),
			Errors: []string{"Unknown flow"},
		}
	}

	// Call your business logic, passing, for example, the username and the target
	// doSomething(log.actor.username, target.srn.slug);
	fmt.Printf("Target: %+v", target)

	return SymLambdaResponse{}
}

func handleDeescalate(log *SymLogEntry) SymLambdaResponse {
	target := ParsedSrn{Srn: log.Fields.Target.Srn}

	// Call your business logic, passing, for example, the username and the target
	// doSomething(log.actor.username, target.srn.slug);
	fmt.Printf("Target: %+v", target)

	return SymLambdaResponse{}
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, log *SymLogEntry) SymLambdaResponse {
	if log.Event.Type == "escalate" {
		return handleEscalate(log)
	} else if log.Event.Type == "deescalate" {
		return handleDeescalate(log)
	} else {
		return SymLambdaResponse{
			Body:   make(map[string]interface{}),
			Errors: []string{"Unknown event type"},
		}
	}
}

func main() {
	lambda.Start(Handler)
}
