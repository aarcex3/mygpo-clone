"""Defintion for the device table """

import enum
import uuid
from typing import Optional

from sqlmodel import Field, Relationship, SQLModel

from src.models.subscription import Subscription
from src.models.user import User


class DeviceType(enum.Enum):
    """Type of devices"""

    DESKTOP = "DESKTOP"
    MOBILE = "MOBILE"
    SERVER = "SERVER"
    OTHER = "OTHER"


class Device(SQLModel, table=True):
    """Device in DB model"""

    id: Optional[uuid.UUID] = Field(default_factory=uuid.uuid1, primary_key=True)
    caption: str = Field(unique=True, index=True)
    device_type: DeviceType
    podcasts: Optional[list["Podcast"]] = Relationship(  # type: ignore
        back_populates="subscribers", link_model=Subscription
    )
    owner_id: uuid.UUID = Field(foreign_key="user.id")
    owner: User = Relationship(back_populates="devices")
