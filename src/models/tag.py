"""Defintion for the tag table """

import uuid
from typing import Optional

from sqlmodel import Field, Relationship, SQLModel

from src.models.podcasttag import PodcastTag


class Tag(SQLModel, table=True):
    """Tag in DB model"""

    id: Optional[uuid.UUID] = Field(default_factory=uuid.uuid1, primary_key=True)
    name: str = Field(unique=True, index=True)
    usage: int
    podcasts: list["Podcast"] = Relationship(  # type: ignore
        back_populates="tags", link_model=PodcastTag
    )
