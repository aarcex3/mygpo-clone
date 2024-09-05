from collections.abc import AsyncGenerator
from typing import Any

from advanced_alchemy.service import SQLAlchemyAsyncRepositoryService
from sqlalchemy import select
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy.orm import joinedload, selectinload

from src.podcast.models import Podcast
from src.podcast.repository import PodcastRepository
from src.podcast_tag.models import PodcastTag
from src.tag.models import Tag


class PodcastService(SQLAlchemyAsyncRepositoryService[Podcast]):
    """Podcast Service"""

    repository_type = PodcastRepository

    def __init__(self, session: AsyncSession, **repo_kwargs: Any) -> None:
        self.repository: PodcastRepository = self.repository_type(
            session=session, **repo_kwargs
        )  # type: ignore
        self.model_type = self.repository.model_type


async def provide_podcast_service(
    session: AsyncSession,
) -> AsyncGenerator[PodcastService, Any]:

    async with PodcastService.new(
        session=session,
        load=[selectinload(Podcast.tags)],
    ) as service:
        yield service
