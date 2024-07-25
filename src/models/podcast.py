"""Defintion for the podcast table"""

import uuid
from typing import Optional

from sqlmodel import Field, Relationship, SQLModel

from src.models.podcastlist import PodcastList
from src.models.podcasttag import PodcastTag
from src.models.subscription import Subscription


class Podcast(SQLModel, table=True):
    """Podcast in DB model"""

    id: Optional[uuid.UUID] = Field(default_factory=uuid.uuid1, primary_key=True)
    name: str = Field(unique=True, index=True)
    description: str
    website: str
    xml_url: str
    author: str
    subscribers_count: int
    subscribers: list["Device"] = Relationship(  # type: ignore
        back_populates="podcasts", link_model=Subscription
    )
    tags: list["Tag"] = Relationship(  # type: ignore
        back_populates="podcasts", link_model=PodcastTag
    )
    episodes: list["Episode"] = Relationship(back_populates="podcast")  # type: ignore
    lists: list["UserList"] = Relationship(  # type: ignore
        back_populates="podcasts", link_model=PodcastList
    )
