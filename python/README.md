# Sym Python Lambda Skeleton

Refer to the top-level [README](../README.md) for general information on this skeleton.

## Schema

The schema of Sym's requests to you can be found [here](https://sym.stoplight.io/docs/sym-reporting), and is represented as a Pydantic model in [sym_handler/models.py](sym_handler/models.py).


## Local Testing

[sym_handler/handler.py](sym_handler/handler.py) includes a main method with an example escalation event.

This lets you run the lambda locally using:

```
$ pip install -r requirements.txt
$ python -m sym_handler.handler
```
## Need Help?

Please [reach out](https://docs.symops.com/docs/support) with any questions on this example or help getting started.
