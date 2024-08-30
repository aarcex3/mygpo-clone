from typing import Any

from advanced_alchemy.service import SQLAlchemyAsyncRepositoryService
from sqlalchemy.ext.asyncio import AsyncSession

from src.tag.models import Tag
from src.tag.repository import TagRepository


class TagService(SQLAlchemyAsyncRepositoryService[Tag]):
    """Tag Service"""

    repository_type = TagRepository

    def __init__(self, session: AsyncSession, **repo_kwargs: Any) -> None:
        self.repository: TagRepository = self.repository_type(
            session=session, **repo_kwargs
        )  # type: ignore
        self.model_type = self.repository.model_type
