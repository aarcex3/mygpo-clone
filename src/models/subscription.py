"""Defintion for the link table for devices and podcasts"""

import uuid
from typing import Optional

from sqlmodel import Field, SQLModel


class Subscription(SQLModel, table=True):
    """Link table for podcasts and devices"""

    device_id: Optional[uuid.UUID] = Field(foreign_key="device.id", primary_key=True)
    podcast_id: Optional[uuid.UUID] = Field(foreign_key="podcast.id", primary_key=True)
