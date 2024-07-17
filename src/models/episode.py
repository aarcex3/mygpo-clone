"""Definition for episode table"""

from datetime import datetime
from typing import Optional
from sqlmodel import SQLModel, Field, Relationship
from src.models.podcast import Podcast


class Episode(SQLModel, table=True):
    """Podcast in DB model"""

    id: Optional[int] = Field(default=None, primary_key=True, index=True)
    title: str
    description: str
    audio_url: str
    duration: int
    podcast: Podcast = Relationship(back_populates="episodes")
    podcast_id: Optional[int] = Field(foreign_key="podcast.id")
    released_on: datetime
