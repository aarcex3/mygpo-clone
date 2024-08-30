from __future__ import annotations

from typing import TYPE_CHECKING, Any

from src.person.models import Person
from src.person.service import PersonService

if TYPE_CHECKING:
    from collections.abc import AsyncGenerator

    from sqlalchemy.ext.asyncio import AsyncSession


async def provide_person_service(
    session: AsyncSession,
) -> AsyncGenerator[PersonService, Any]:
    """This provides the default Person repository."""
    async with PersonService.new(
        session=session,
        error_messages={
            "duplicate_key": "This user already exists.",
            "integrity": "User operation failed.",
        },
    ) as service:
        yield service
