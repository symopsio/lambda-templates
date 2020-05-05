import json
from google.protobuf import json_format
from sym.enums.service_pb2 import Service
from sym.messages.approval_pb2 import Approval
from sym.messages.expiration_pb2 import Expiration

# Approval functions get an Approval event
# Return the JSON form of an ApprovalResponse object
def handle_approve(event, context):
    try:
        approval = json_format.ParseDict(event, Approval())
        user_id = get_identity(approval.request.target.user, Service.SLACK)
        print("Target user slack id:", user_id)
        print("Target resource:", approval.request.target.resource)
        print("Target reason:", approval.request.meta.reason)
    except Exception as e:
        print("Unexpected exception:", str(e))
        return {'ok': False, 'error': str(e)}
    return {'ok': True}

# Expiration functions get an Expiration event
# Return the JSON form of an ExpirationResponse object
def handle_expire(event, context):
    try:
        expiration = json_format.ParseDict(event, Expiration())
        user_id = get_identity(expiration.target.user, Service.SLACK)
        print("Target user slack id:", user_id)
        print("Target resource:", expiration.target.resource)
    except Exception as e:
        print("Unexpected exception:", str(e))
        return {'ok': False, 'error': str(e)}
    return {'ok': True}

# Utility function to get a user ID for a specific service
# This will be in a shared Sym library eventually
def get_identity(user, service):
    for identity in user.identities:
        if identity.service == service:
            return identity.id

# Helpful for local testing
def main():
    approval = {
        'request': {
            'meta': {
                'reason': 'Access to Bucket 123'
            },
            'target': {
                'user': {
                    'identities': [{
                        'service': 'SLACK',
                        'id': 'foo@example.com'
                    }]
                },
                'resource': {
                    'id': 'my-custom-resource',
                    'service': 'CUSTOM'
                }
            }
        },
        'meta': {
            'approver': {
                'identities': [{
                    'service': 'SLACK',
                    'id': 'foo@example.com'
                }]
            }
        }
    }
    response = handle_approve(approval, '')
    print("Got Response: ", response)

if __name__ == "__main__":
    main()
