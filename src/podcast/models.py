from litestar.contrib.sqlalchemy.base import UUIDAuditBase
from sqlalchemy.ext.associationproxy import AssociationProxy, association_proxy
from sqlalchemy.orm import Mapped, mapped_column, relationship

from src.podcast_tag.models import PodcastTag


class Podcast(UUIDAuditBase):
    """Podcast in DB model"""

    name: Mapped[str] = mapped_column(unique=True)
    description: Mapped[str]
    xml_url: Mapped[str] = mapped_column(unique=True, index=True)
    website: Mapped[str]
    subscribers_count: Mapped[int]
    author: Mapped[str]
    tags: Mapped[list["Tag"]] = relationship(  # type: ignore
        "Tag",
        secondary=PodcastTag,
        back_populates="podcasts",
    )
    tag_list: AssociationProxy[list[str]] = association_proxy("tags", "name")
