"""Defintion for the link table for podcast and tag"""

import uuid
from typing import Optional

from sqlmodel import Field, SQLModel


class PodcastTag(SQLModel, table=True):
    """Link table for tags and podcasts"""

    tag_id: Optional[uuid.UUID] = Field(foreign_key="tag.id", primary_key=True)
    podcast_id: Optional[uuid.UUID] = Field(foreign_key="podcast.id", primary_key=True)
