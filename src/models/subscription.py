"""Defintion for the link table for devices and podcasts"""

from typing import Optional
from sqlmodel import SQLModel, Field


class Subscription(SQLModel, table=True):
    """Link table for podcasts and devices"""

    device_id: Optional[int] = Field(foreign_key="device.id", primary_key=True)
    podcast_id: Optional[int] = Field(foreign_key="podcast.id", primary_key=True)
