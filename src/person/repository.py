from __future__ import annotations

from typing import TYPE_CHECKING

from advanced_alchemy.repository import SQLAlchemyAsyncRepository

from src.person.models import Person

if TYPE_CHECKING:

    from sqlalchemy.ext.asyncio import AsyncSession


class PersonRepository(
    SQLAlchemyAsyncRepository[Person]
):  # pylint: disable=duplicate-bases
    """User Repository"""

    model_type = Person


async def provide_person_repo(session: AsyncSession) -> PersonRepository:
    """Provide an instance of the user repository

    Args:
        db (AsyncSession): Database session

    Returns:
        PersonRepository: PersonRepository instance
    """
    return PersonRepository(session=session)
