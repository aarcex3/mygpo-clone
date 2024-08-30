from litestar.contrib.sqlalchemy.base import UUIDAuditBase
from sqlalchemy.orm import Mapped, mapped_column, relationship

from src.podcast_tag.models import PodcastTag


class Tag(UUIDAuditBase):
    """Tag in DB model"""

    name: Mapped[str] = mapped_column(unique=True)
    code: Mapped[str] = mapped_column(unique=True, index=True)
    usage: Mapped[int]
    podcasts: Mapped[list["Podcast"]] = relationship(  # type: ignore
        "Podcast", secondary=PodcastTag, back_populates="tags", lazy="selectin"
    )
