"""Defintion for the tag table """

from typing import Optional
from sqlmodel import SQLModel, Field, Relationship
from src.models.podcasttag import PodcastTag


class Tag(SQLModel, table=True):
    """Tag in DB model"""

    id: Optional[int] = Field(default=None, primary_key=True, index=True)
    name: str = Field(unique=True, index=True)
    usage: int
    podcasts: list["Podcast"] = Relationship(  # type: ignore
        back_populates="tags", link_model=PodcastTag
    )
