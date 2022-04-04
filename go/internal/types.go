package internal

// SymActor has metadata on the user that initiated the Sym Event
type SymActor struct {
	Username string `json:"username"`
}

// SymEventDetails has data on the kind of event this is
type SymEventDetails struct {
	Type string `json:"type"`
}

// Starter model for Sym events, just including the username and event type
// properties that are required by our example Lambda implementation.
// For more details on the event object format, refer to our reporting docs:
// https://docs.symops.com/docs/reporting
type SymEvent struct {
	Actor   *SymActor        `json:"actor"`
	Details *SymEventDetails `json:"event"`
}

// Sym expects your Lambda implementation to return an object a body that can be
// empty or can include messages your Sym workflow handler processes. The object
// can also return error messages that will be surfaced to the user and Sym
// implementer.
type SymResult struct {
	Body   map[string]interface{} `json:"body"`
	Errors []string               `json:"errors"`
}
