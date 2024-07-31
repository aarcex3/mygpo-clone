"""Definition for episode table"""

import uuid
from datetime import datetime
from typing import Optional

from sqlmodel import Field, Relationship, SQLModel

from src.models.podcast import Podcast


class Episode(SQLModel, table=True):
    """Podcast in DB model"""

    id: Optional[int] = Field(decimal_places=None,index=True, primary_key=True)
    title: str
    description: str
    audio_url: str
    duration: int
    podcast: Podcast = Relationship(back_populates="episodes")
    podcast_id: Optional[int] = Field(foreign_key="podcast.id")
    released_on: datetime
