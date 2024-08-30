"""
Main entry point for the app
"""

from litestar import Litestar
from litestar.contrib.sqlalchemy.plugins import SQLAlchemyPlugin
from litestar.di import Provide

from src.auth.controller import AuthController
from src.auth.service import AuthService
from src.database import db_connection, provide_session, sqlalchemy_config

# from src.redis.service import provide_redis_service
from src.directory.controller import DirectoryController
from src.person.repository import provide_person_repo
from src.person.service import PersonService
from src.podcast.repository import provide_podcast_repo
from src.podcast.service import provide_podcast_service
from src.security import AUTH
from src.tag.repository import provide_tag_repo
from src.tag.service import TagService

app = Litestar(
    debug=True,
    route_handlers=[AuthController, DirectoryController],
    dependencies={
        "auth_service": Provide(AuthService, sync_to_thread=False),
        "person_service": Provide(PersonService, sync_to_thread=False),
        "person_repo": Provide(provide_person_repo),
        "podcast_service": Provide(provide_podcast_service),
        "podcast_repo": Provide(provide_podcast_repo),
        "tag_service": Provide(TagService, sync_to_thread=False),
        "tag_repo": Provide(provide_tag_repo),
        # "redis_service": Provide(provide_redis_service),
        "session": Provide(provide_session),
    },
    plugins=[SQLAlchemyPlugin(sqlalchemy_config)],
    lifespan=[db_connection],
    on_app_init=[AUTH.on_app_init],
)


if __name__ == "__main__":
    import uvicorn

    uvicorn.run("src.main:app", host="0.0.0.0", port=8000)
