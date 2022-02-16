from datetime import datetime
from typing import Dict, List, Optional, Union

from pydantic import BaseModel
from sym.sdk import SRN


class LogEntryRun(BaseModel):
    srn: SRN
    parent: Optional[SRN]
    flow: SRN


class LogEntryIdentity(BaseModel):
    """This object will also contain any fields contained in the specific Identity (e.g. the User ID of a Slack User)."""

    class Config:
        extra = "allow"

    service: str
    external_id: str


class LogEntryActor(BaseModel):
    user: SRN
    name: str
    username: Optional[str]
    identity: LogEntryIdentity


class LogEntryTarget(BaseModel):
    name: str
    srn: SRN
    type: str
    label: str
    settings: Dict[str, str]


class LogEntryMeta(BaseModel):
    schema_version: int = 4


class LogEntryError(BaseModel):
    message: str
    code: str


class LogEntryState(BaseModel):
    status: str
    errors: List[LogEntryError]


class LogEntryEvent(BaseModel):
    srn: SRN
    type: str
    template: SRN
    timestamp: datetime


class LogEntryFields(BaseModel):
    """This object will also contain all of the user-supplied fields."""

    class Config:
        extra = "allow"


class LogEntryApprovalFields(LogEntryFields):
    reason: str
    duration: int
    target: LogEntryTarget


class SymLogEntry(BaseModel):
    id: str
    meta: LogEntryMeta
    state: LogEntryState
    run: LogEntryRun
    event: LogEntryEvent
    actor: LogEntryActor
    fields: Union[LogEntryApprovalFields, LogEntryFields]
    type: str
