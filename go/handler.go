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

	if flow.Slug() != "my-lambda-flow" {
		return SymLambdaResponse{
			Body:   make(map[string]interface{}),
			Errors: []string{"Unknown flow"},
		}
	}

	// Call your business logic, passing, for example, the username and the target
	// doSomething(log.actor.username, target.srn.slug);
	fmt.Printf("Fields: %+v", log.Fields)

	return SymLambdaResponse{make(map[string]interface{}), []string{}}
}

func handleDeescalate(log *SymLogEntry) SymLambdaResponse {

	// Call your business logic, passing, for example, the username and the target
	// doSomething(log.actor.username, target.srn.slug);
	fmt.Printf("Fields: %+v", log.Fields)

	return SymLambdaResponse{make(map[string]interface{}), []string{}}
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, log *SymLogEntry) (SymLambdaResponse, error) {
	fmt.Printf("Event: %+v", log)
	fmt.Printf("Event TYpe %s", log.Event.Type)
	if log.Event.Type == "escalate" {
		return handleEscalate(log), nil
	} else if log.Event.Type == "deescalate" {
		return handleDeescalate(log), nil
	} else {
		return SymLambdaResponse{
			Body:   make(map[string]interface{}),
			Errors: []string{"Unknown event type"},
		}, nil
	}
}

func main() {
	lambda.Start(Handler)
}
