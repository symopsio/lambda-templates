export type SchemaVersion = number;
/**
 * A Sym Resource Name (SRN) is a unique identifier for a Sym Resource.
 */
export type SRN = string;

export type UUID = string;

export type Status = string;

export type Type = string;

/**
 * This model is sent by Sym both to Log Destinations for [reporting](https://docs.symops.com/docs/reporting), as well as for the [HTTP](https://docs.symops.com/docs/http) and [AWS Lambda](https://docs.symops.com/docs/aws-lambda) Strategies.
 */
export interface SymLogEntry {
  id: UUID;
  meta: LogEntryMeta;
  state: LogEntryState;
  run: LogEntryRun;
  event: LogEntryEvent;
  actor: LogEntryActor;
  fields: LogEntryApprovalFields;
  type: Type;
}
export interface LogEntryMeta {
  schema_version?: SchemaVersion;
}

export interface LogEntryError {
  message: string;
  code: string;
}

export interface LogEntryState {
  status: Status;
  errors: LogEntryError[];
}
export interface LogEntryRun {
  srn: SRN;
  parent?: SRN;
  flow: SRN;
}
export interface LogEntryEvent {
  srn: SRN;
  type: string;
  template: SRN;
  timestamp: string;
  channel: string;
}
export interface LogEntryActor {
  user: SRN;
  name: string;
  username?: string;
  identity: LogEntryIdentity;
}

/**
 * This object will also contain any fields contained in the specific Identity (e.g. the User ID of a Slack User).
 */
export interface LogEntryIdentity {
  service: string;
  external_id: string;
  [k: string]: unknown;
}

/**
 * This object will also contain all of the user-supplied fields.
 */
export interface LogEntryApprovalFields {
  reason: string;
  duration: number;
  target: LogEntryTarget;
  [k: string]: unknown;
}

export interface LogEntryTarget {
  name: string;
  srn: SRN;
  type: string;
  label: string;
  settings: Settings;
}

export interface Settings {
  [k: string]: string;
}
export class ParsedSRN {
  // SRNs have the following structure:
  // <ORG>:<MODEL>:<SLUG>:<VERSION>[:<IDENTIFIER>]

  srn: SRN;

  constructor(srn: SRN) {
    this.srn = srn;
  }

  getComponents(): string[] {
    return this.srn.split(":");
  }

  public get org(): string {
    return this.getComponents()[0];
  }

  public get model(): string {
    return this.getComponents()[1];
  }

  public get slug(): string {
    return this.getComponents()[2];
  }

  public get version(): string {
    return this.getComponents()[3];
  }

  public get identifier(): string {
    return this.getComponents()[4];
  }
}
