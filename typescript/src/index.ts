import * as path from "path";
import * as fs from "fs";
import { Context } from "aws-lambda";

import type { SymEvent, SymResult } from "./models";

/**
 * Processes an escalation or deescalation event from the Sym platform.
 * @param event
 * @param context
 * @returns body with optional messages and errors
 */
// eslint-disable-next-line @typescript-eslint/no-unused-vars
export const handler = async (event: SymEvent, _context?: Context): Promise<SymResult> => {
  console.log("Got Event: " + JSON.stringify(event, null, 2))

  try {
    const userName = resolveUser(event);
    const message = updateUser(userName, event);
    return { body: { "message": message }, errors: [] }
  } catch (error) {
    return { body: {}, errors: [getErrorMessage(error)] }
  }
};

/**
 * Takes the incoming user from the event and resolves to the right user id for
 * the system you're escalating the user in.
 * @param event
 * @returns username
 */
function resolveUser(event: SymEvent): string {
  return event.actor.username;
}

/**
 * Placeholder to handle updating the given username based on the event type
 * @param username
 * @param event
 * @returns message to send back to Sym
 */
function updateUser(username: string, event: SymEvent): string {
  const eventType = event.event.type;
  switch(eventType) {
    case "escalate":
      return `Escalating user: ${username}`;
    case "deescalate":
      return `Deescalating user: ${username}`;
    default:
      throw new Error(`Unsupported event type: ${eventType}`);
  }
}

/**
 * Utility function to safely get an error message
 * @param error
 * @returns message
 */
function getErrorMessage(error: unknown) {
  if (error instanceof Error) return error.message
  return String(error)
}

/**
 * Finds the right JSON file for local testing
 * @param arg
 * @return SymEvent
 */
function resolveLocalJson(arg: string): SymEvent {
  let file;
  switch(arg) {
    case "-d":
      file = "deescalate.json"
      break;
    case "-e":
      file = "escalate.json"
      break;
    default:
      throw new Error(`Specify either -e or -d, you supplied: "${arg}"`);
  }
  const fullPath = path.join(__dirname, "..", "..", "test", file);
  const data = fs.readFileSync(fullPath);
  return JSON.parse(data.toString());
}

/**
 * Lets you test your function locally, with an escalate or deescalate payload
 * in the ../../test directory.
 *
 * $ npx ts-node src/index.ts [-e | -d]
 * or
 * $ npm run-script local -- [-e | -d]
 */
if (require.main === module) {
  try {
    const arg = process.argv[2]
    const event = resolveLocalJson(arg);
    handler(event).then((res) => {
      console.log("Got Result: " + JSON.stringify(res, null, 2));
    });
  } catch (error) {
    console.log(getErrorMessage(error))
  }
}
