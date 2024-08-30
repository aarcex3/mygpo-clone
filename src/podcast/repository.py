from __future__ import annotations

from typing import TYPE_CHECKING

from advanced_alchemy.repository import SQLAlchemyAsyncRepository

from src.podcast.models import Podcast

if TYPE_CHECKING:

    from sqlalchemy.ext.asyncio import AsyncSession


class PodcastRepository(
    SQLAlchemyAsyncRepository[Podcast]
):  # pylint: disable=duplicate-bases
    """Podcast Repository"""

    model_type = Podcast


async def provide_podcast_repo(session: AsyncSession) -> PodcastRepository:
    """Provide an instance of the podcast repository

    Args:
        db (AsyncSession): Database session

    Returns:
        PodcastRepository: PodcastRepository instance
    """
    return PodcastRepository(session=session)
