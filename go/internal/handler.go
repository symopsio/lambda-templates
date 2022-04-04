package internal

import (
	"fmt"
	"log"
)

// HandleSymEvent includes placeholders to resolve the user and handle
// escalation/deescalation.
func HandleSymEvent(event SymEvent) (*SymResult, error) {
	log.Printf("Got actor: %+v\n", event.Actor)
	log.Printf("Got details: %+v\n", event.Details)
	username, err := ResolveUser(event)
	if err != nil {
		return ResultForError(err), nil
	}

	message, err := UpdateUser(username, event)
	if err != nil {
		return ResultForError(err), nil
	}

	log.Printf("Update complete: %s", message)

	return &SymResult{Body: map[string]interface{}{"message": message}, Errors: []string{}}, nil
}

// ResultForError creates a SymResult payload for an error message
func ResultForError(err error) *SymResult {
	s := fmt.Sprintf("%e", err)
	return &SymResult{Body: map[string]interface{}{}, Errors: []string{s}}
}

// ResolveUser is a placeholder that takes the incoming user from the Sym event
// and resolves to the right user id for the system you're escalating the user in.
func ResolveUser(event SymEvent) (string, error) {
	return event.Actor.Username, nil
}

// UpdateUser is a placeholder to handle updating the given username based on
// the event type
func UpdateUser(username string, event SymEvent) (string, error) {
	eventType := event.Details.Type
	switch eventType {
	case "escalate":
		return fmt.Sprintf("Escalating user: %s", username), nil
	case "deescalate":
		return fmt.Sprintf("Deescalating user: %s", username), nil
	default:
		return "", fmt.Errorf("Unsupported event type: %s", eventType)
	}
}
