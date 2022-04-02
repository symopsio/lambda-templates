/**
 * Starter model for Sym events, just including the username and event type
 * properties that are required by our example Lambda implementation.
 *
 * For more details on the event object format, refer to our reporting docs:
 * https://docs.symops.com/docs/reporting
 */
export interface SymEvent {
  actor: SymActor;
  event: SymEventDetails;
}

/**
 * Metadata on the user that initiated the Sym event
 */
export interface SymActor {
  username: string;
}

/**
 * Details on the kind of event this is.
 */
export interface SymEventDetails {
  type: string;
}

/**
 * Sym expects your Lambda implementation to return an object a body that can be
 * empty or can include messages your Sym workflow handler processes. The object
 * can also return error messages that will be surfaced to the user and Sym
 * implementer.
 */
export interface SymResult {
  body: object;
  errors: string[];
}
