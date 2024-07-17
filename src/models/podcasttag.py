"""Defintion for the link table for podcast and tag"""

from typing import Optional
from sqlmodel import SQLModel, Field


class PodcastTag(SQLModel, table=True):
    """Link table for tags and podcasts"""

    tag_id: Optional[int] = Field(foreign_key="tag.id", primary_key=True)
    podcast_id: Optional[int] = Field(foreign_key="podcast.id", primary_key=True)
