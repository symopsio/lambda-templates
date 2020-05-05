import { Handler, Context } from "aws-lambda";
import type { sym as t } from "@symops/types";
import { sym as protos } from "@symops/types";

/*
 * Create an Approval object from the incoming event and then escalate the
 * target user
 *
 * Return a response object with an error message if the action fails.
 */
export const handleApprove: Handler = async (
  event: any,
  _context: Context
): Promise<t.messages.IApprovalResponse> => {
  let approval;
  try {
    approval = protos.messages.Approval.create(event);
  } catch (error) {
    console.log(`Unable to deserialize approval: ${error}`);
    return { ok: false, error: error };
  }

  // Get the target user's slack identity (their email address)
  const identity = getIdentity(
    approval.request?.target?.user,
    protos.enums.Service.SLACK
  );
  console.log(`Target user slack id: ${identity}`);
  console.log(`Target resource service: ${approval.request?.target?.resource?.service}`);
  console.log(`Target resource id: ${approval.request?.target?.resource?.id}`);
  console.log(`Reason: ${approval.request?.meta?.reason}`);

  return { ok: true };
};

/*
 * Create an Approval object from the incoming event and then expire the
 * target user
 *
 * Return a response object with an error message if the action fails.
 */
export const handleExpire: Handler = async (
  event: any,
  _context: Context
): Promise<t.messages.IExpirationResponse> => {
  let expiration;
  try {
    expiration = protos.messages.Expiration.create(event);
  } catch (error) {
    console.log(`Unable to deserialize expiration: ${error}`);
    return { ok: false, error: error };
  }

  // Get the user's slack identity (their email address)
  const identity = getIdentity(
    expiration.target?.user,
    protos.enums.Service.SLACK
  );
  console.log(`Target user slack id: ${identity}`);
  console.log(`Target resource service: ${expiration.target?.resource?.service}`);
  console.log(`Target resource id: ${expiration.target?.resource?.id}`);

  return { ok: true };
};

/*
 * Helper function to get the user's id in a given target service
 * This will be in a shared Sym library eventually
 */
const getIdentity = (
  user: protos.models.User,
  service: protos.enums.Service
): string | null => {
  for (const identity of user.identities) {
    if (identity.service === service) {
      return identity.id;
    }
  }
};
