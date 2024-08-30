from collections.abc import AsyncGenerator
from contextlib import asynccontextmanager

from advanced_alchemy.base import UUIDAuditBase
from advanced_alchemy.extensions.litestar import (
    AsyncSessionConfig,
    SQLAlchemyAsyncConfig,
)
from advanced_alchemy.extensions.litestar.plugins.init.config.asyncio import (
    autocommit_before_send_handler,
)
from litestar import Litestar
from litestar.datastructures import State
from litestar.exceptions import ClientException
from litestar.status_codes import HTTP_409_CONFLICT
from sqlalchemy.exc import IntegrityError
from sqlalchemy.ext.asyncio import AsyncSession, async_sessionmaker, create_async_engine

from src.config import get_settings

settings = get_settings()


sqlalchemy_config = SQLAlchemyAsyncConfig(
    connection_string=settings.DB_URL,
    create_all=True,
    before_send_handler="autocommit",
    session_config=AsyncSessionConfig(expire_on_commit=False),
)


@asynccontextmanager
async def db_connection(app: Litestar) -> AsyncGenerator[None, None]:
    """Create the db connection"""
    if "engine" not in app.state:
        app.state.engine = create_async_engine(settings.DB_URL, echo=True)

    async with app.state.engine.begin() as conn:
        await conn.run_sync(UUIDAuditBase.metadata.create_all)

    try:
        yield
    finally:
        await app.state.engine.dispose()


sessionmaker = async_sessionmaker(expire_on_commit=False)


async def provide_session(state: State) -> AsyncGenerator[AsyncSession, None]:
    """Get the db session"""
    if not hasattr(state, "engine") or state.engine is None:
        raise RuntimeError("Database engine not initialized")

    async with sessionmaker(bind=state.engine) as session:
        try:
            async with session.begin():
                yield session
        except IntegrityError as exc:
            raise ClientException(
                status_code=HTTP_409_CONFLICT,
                detail=str(exc),
            ) from exc
