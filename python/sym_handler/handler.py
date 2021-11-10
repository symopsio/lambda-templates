from pathlib import Path

from devtools import debug

from .models import SymLogEntry


def handle(event, context):
    log = SymLogEntry.parse_obj(event)
    print("Got event:")
    debug(log)

    if log.event.type == "escalate":
        return handle_escalate(log)
    elif log.event.type == "deescalate":
        return handle_deescalate(log)
    else:
        return {"ok": False, "error": "Unknown event type"}


def handle_escalate(log: SymLogEntry):
    if log.run.flow.slug != "my-lambda-flow":
        return {"ok": False, "error": "Unknown flow"}

    # Call your business logic, passing, for example, the username and the target
    # do_something(log.actor.username, log.fields.target.srn.slug)

    return {"ok": True}


def handle_deescalate(log: SymLogEntry):
    # Call your business logic, passing, for example, the username and the target
    # do_something(log.actor.username, log.fields.target.srn.slug)

    ...

    return {"ok": True}


if __name__ == "__main__":
    path = Path(__file__).parent.parent.parent / "payload.json"
    log = SymLogEntry.parse_file(path)
    handle(log, None)
