from __future__ import annotations

from typing import TYPE_CHECKING

from advanced_alchemy.repository import SQLAlchemyAsyncRepository

from src.tag.models import Tag

if TYPE_CHECKING:

    from sqlalchemy.ext.asyncio import AsyncSession


class TagRepository(SQLAlchemyAsyncRepository[Tag]):  # pylint: disable=duplicate-bases
    """Tag Repository"""

    model_type = Tag


async def provide_tag_repo(session: AsyncSession) -> TagRepository:
    """Provide an instance of the tag repository

    Args:
        db (AsyncSession): Database session

    Returns:
        TagRepository: TagRepository instance
    """
    return TagRepository(session=session)
