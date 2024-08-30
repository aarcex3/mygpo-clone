from advanced_alchemy.extensions.litestar.dto import SQLAlchemyDTO, SQLAlchemyDTOConfig

from src.podcast.models import Podcast


class PodcastDTO(SQLAlchemyDTO[Podcast]):
    config = SQLAlchemyDTOConfig(exclude={"id", "created_at", "updated_at", "tags"})
