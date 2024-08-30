from advanced_alchemy.extensions.litestar.dto import SQLAlchemyDTO, SQLAlchemyDTOConfig

from src.tag.models import Tag


class TagDTO(SQLAlchemyDTO[Tag]):
    config = SQLAlchemyDTOConfig(exclude={"id", "created_at", "updated_at", "podcasts"})
