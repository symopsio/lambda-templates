from pathlib import Path

from devtools import debug

from .models import SymLogEntry, SymLambdaResponse


def handle(event, context) -> dict:
    log = SymLogEntry.parse_obj(event)
    print("Got event:")
    debug(log)

    if log.event.type == "escalate":
        response = handle_escalate(log)
    elif log.event.type == "deescalate":
        response = handle_deescalate(log)
    else:
        response = SymLambdaResponse(body={"ok": False}, errors=["Unknown event type"])

    return response.dict()


def handle_escalate(log: SymLogEntry) -> SymLambdaResponse:
    if log.run.flow.slug != "my-lambda-flow":
        return SymLambdaResponse(body={"ok": False}, errors=["Unknown flow"])

    # Call your business logic, passing, for example, the username and the target
    # do_something(log.actor.username, log.fields.target.srn.slug)

    return SymLambdaResponse(body={"ok": True})


def handle_deescalate(log: SymLogEntry)-> SymLambdaResponse:
    # Call your business logic, passing, for example, the username and the target
    # do_something(log.actor.username, log.fields.target.srn.slug)

    ...

    return SymLambdaResponse(body={"ok": True})


if __name__ == "__main__":
    path = Path(__file__).parent.parent.parent / "payload.json"
    log = SymLogEntry.parse_file(path)
    handle(log, None)
