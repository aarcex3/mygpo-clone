from typing import Final

from litestar.contrib.sqlalchemy.base import UUIDAuditBase
from sqlalchemy import Column, ForeignKey, Table

PodcastTag: Final[Table] = Table(
    "podcast_tag",
    UUIDAuditBase.metadata,
    Column(
        "podcast_id", ForeignKey("podcast.id", ondelete="CASCADE"), primary_key=True
    ),
    Column("tag_id", ForeignKey("tag.id", ondelete="CASCADE"), primary_key=True),
)
