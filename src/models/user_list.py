"""
Defition for the user podcasts lists
"""

from typing import Optional
from sqlmodel import SQLModel, Field, Relationship
from src.models.podcastlist import PodcastList


class UserList(SQLModel, table=True):
    """User podcast list in DB model"""

    id: Optional[int] = Field(default=None, index=True, primary_key=True)
    name: str = Field(index=True)
    title: str = Field(index=True)
    web: str
    owner: "User" = Relationship(back_populates="podcast_lists")  # type: ignore
    owner_id: Optional[int] = Field(foreign_key="user.id")
    podcasts: list["Podcast"] = Relationship(  # type: ignore
        back_populates="lists", link_model=PodcastList
    )
