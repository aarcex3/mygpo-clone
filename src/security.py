from typing import Any
from uuid import UUID

from litestar.connection import ASGIConnection
from litestar.security.jwt import JWTAuth, Token

from src.auth import urls as auth
from src.config import get_settings
from src.database import sqlalchemy_config
from src.person.guards import provide_person_service
from src.person.models import Person

settings = get_settings()


async def current_user_from_token(
    token: Token, _: ASGIConnection[Any, Any, Any, Any]
) -> Person | None:
    """Retrieve current user from session"""
    async with sqlalchemy_config.get_session() as session:
        service = await anext(provide_person_service(session))
        person = await service.get_one_or_none(Person.id == UUID(token.sub))
        return person if person and person.is_active else None


AUTH = JWTAuth[Person](
    retrieve_user_handler=current_user_from_token,
    token_secret=settings.SECRET_KEY,
    exclude=[
        "/schema",
        "/tags",
        "/search",
        "/toplist",
        "/data",
        auth.REGISTRATION,
        auth.LOGIN,
    ],
)
