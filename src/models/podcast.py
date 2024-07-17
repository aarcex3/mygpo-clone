"""Defintion for the podcast table"""

from typing import Optional
from sqlmodel import SQLModel, Field, Relationship

from src.models.podcastlist import PodcastList
from src.models.subscription import Subscription
from src.models.podcasttag import PodcastTag


class Podcast(SQLModel, table=True):
    """Podcast in DB model"""

    id: Optional[int] = Field(default=None, primary_key=True, index=True)
    name: str = Field(unique=True, index=True)
    description: str
    website: str
    xml_url: str
    author: str
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
