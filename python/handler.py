import json
import sys

from devtools import debug


def handle(event, context) -> dict:
    """
    Escalates or de-escalates depending on the incoming event type.

    For more details on the event object format, refer to our reporting docs:
    https://docs.symops.com/docs/reporting
    """
    print("Got event:")
    debug(event)

    try:
        username = resolve_user(event)
        message = update_user(username, event)
        return {"body": {"message": message}, "errors": []}
    except Exception as e:
        return {"body": {}, "errors": [str(e)]}


def resolve_user(event) -> str:
    """
    Take the incoming user from the event and resolve to the right
    user id for the system you're escalating the user in.
    """
    return event["actor"]["username"]


def update_user(username, event) -> str:
    """
    Placeholder to handle updating the given user based on the event type
    """
    event_type = event["event"]["type"]
    if event_type == "escalate":
        message = f"Escalating user: {username}"
    elif event_type == "deescalate":
        message = f"Deescalating user: {username}"
    else:
        raise RuntimeError(f"Unsupported event type: {event_type}")
    return message


def resolve_local_json(arg) -> str:
    """Find the right test json file based on the arg"""
    if arg == "-d":
        file = "deescalate.json"
    elif arg == "-e":
        file = "escalate.json"
    else:
        raise RuntimeError(f"Specify either -e or -d, you supplied: {arg}")
    return f"../test/{file}"


def run_local() -> dict:
    """
    This lets you test your function code locally, with an escalate or
    deescalate payload (in the ../test) directory.

    $ python handler.py [-e | -d]
    """
    arg = None if len(sys.argv) < 2 else sys.argv[1]
    path = resolve_local_json(arg)
    with open(path, "r") as payload:
        event = json.load(payload)
    return handle(event, {})


if __name__ == "__main__":
    result = run_local()
    debug(result)
