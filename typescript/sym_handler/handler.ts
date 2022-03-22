import path from "path";
import fs from "fs";
import { Context } from "aws-lambda";
import { ParsedSRN } from "./models";
import type { SymLogEntry } from "./models";

export const handle = async (
  log: SymLogEntry,
  _ctx?: Context
): Promise<Object> => {
  console.log("Got Event: ", log);

  if (log.event.type == "escalate") {
    return handleEscalate(log);
  } else if (log.event.type == "deescalate") {
    return handleDeescalate(log);
  } else {
    return { body: { "ok": false}, errors: ["Unknown event type"] };
  }
};

const handleEscalate = async (log: SymLogEntry): Promise<Object> => {
  const flow = new ParsedSRN(log.run.flow);
  const target = new ParsedSRN(log.fields.target.srn);

  if (flow.slug != "my-lambda-flow") {
    return { body: { "ok": false}, errors: ["Unknown flow"] };
  }

  // Call your business logic, passing, for example, the username and the target
  // doSomething(log.actor.username, target.srn.slug);

  return { body: { "ok": true}, errors: [] };
};

const handleDeescalate = async (log: SymLogEntry): Promise<Object> => {
  const target = new ParsedSRN(log.fields.target.srn);

  // Call your business logic, passing, for example, the username and the target
  // doSomething(log.actor.username, target.srn.slug);

  return { body: { "ok": true}, errors: [] };
};

if (require.main === module) {
  const payload = path.join(__dirname, "..", "..", "payload.json");
  const data = fs.readFileSync(payload);
  const log: SymLogEntry = JSON.parse(data.toString());
  handle(log).then((res) => {
    console.log(res);
  });
}
