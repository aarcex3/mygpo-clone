"""Defintion for the device table """

import enum
from typing import Optional
from sqlmodel import SQLModel, Field, Relationship, Enum

from src.models.subscription import Subscription
from src.models.user import User


class DeviceType(enum.Enum):
    """Type of devices"""

    DESKTOP = "DESKTOP"
    LAPTOP = "LAPTOP"
    MOBILE = "MOBILE"
    SERVER = "SERVER"
    OTHER = "OTHER"


class Device(SQLModel, table=True):
    """Device in DB model"""

    id: Optional[int] = Field(default=None, primary_key=True, index=True)
    caption: str = Field(unique=True, index=True)
    type: Enum[DeviceType]
    podcasts: list["Podcast"] = Relationship(  # type: ignore
        back_populates="subscribers", link_model=Subscription
    )
    owner_id: Optional[int] = Field(foreign_key="user.id")
    owner: User = Relationship(back_populates="devices")
