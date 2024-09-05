from typing import Sequence

from advanced_alchemy.filters import LimitOffset, OrderBy
from litestar import MediaType, get
from litestar.controller import Controller
from litestar.repository.filters import LimitOffset, OrderBy
from sqlalchemy import distinct, select

from src.directory import urls
from src.podcast.dto import PodcastDTO
from src.podcast.models import Podcast
from src.podcast.service import PodcastService
from src.podcast_tag.models import PodcastTag
from src.tag.dto import TagDTO
from src.tag.models import Tag
from src.tag.service import TagService


class DirectoryController(Controller):
    """Directory Controller"""

    @get(path=urls.RETRIEVE_TOP_TAGS, media_type=MediaType.JSON, return_dto=TagDTO)
    async def top_tags(self, tag_service: TagService, count: int = 5) -> Sequence[Tag]:

        return await tag_service.list(
            LimitOffset(limit=count, offset=0), OrderBy("usage", "desc")
        )

    @get(
        path=urls.RETRIEVE_PODCASTS_FOR_TAG,
        media_type=MediaType.JSON,
        return_dto=PodcastDTO,
    )
    async def podcasts_for_tag(
        self,
        podcast_service: PodcastService,
        tag_code: str,
        count: int = 3,
    ) -> Sequence[Podcast]:

        return await podcast_service.list(
            PodcastTag.c.tag_id.in_(select(Tag.id).where(Tag.code == tag_code)),
            LimitOffset(limit=count, offset=0),
            OrderBy("subscribers_count", "desc"),
        )

    # @get(path=urls.RETRIEVE_PODCAST_DATA, media_type=MediaType.JSON)
    # async def podcast_data(self, directory_service: DirectoryService, podcast_url: str):
    #     pass

    # @get(path=urls.RETRIEVE_EPISODE_DATA, media_type=MediaType.JSON)
    # async def episode_data(self, directory_service: DirectoryService, episode_url: str):
    #     pass

    # @get(path=urls.PODCAST_TOPLIST, media_type=MediaType.JSON)
    # async def podcast_toplist(self, directory_service: DirectoryService, count: int):
    #     pass

    # @get(path=urls.PODCAST_SEARCH, media_type=MediaType.JSON)
    # async def podcast_search(self, directory_service: DirectoryService, query: str):
    #     pass
