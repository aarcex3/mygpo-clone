"""
CRUD methods for the directory endpoints
"""

from fastapi import HTTPException, status
from sqlmodel import Session, desc, or_, select

from src.models.episode import Episode
from src.models.podcast import Podcast
from src.models.podcasttag import PodcastTag
from src.models.tag import Tag
from src.schemas.podcast import PodcastOut


async def get_top_tags(count: int, session: Session) -> list[Tag]:
    """
    Get the top n tags based on usage, where n = count
    """
    top_tags = session.exec(select(Tag).order_by(desc(Tag.usage)).limit(count)).all()
    return [tag.model_dump(exclude={"id"}) for tag in top_tags]


async def get_podcast_tags(code: str, count: int, session: Session) -> list[Podcast]:
    """
    Get the top n (count) podcasts associated with the tag (code).
    """
    # Define the query to get podcasts associated with a specific tag
    statement = (
        select(Podcast)
        .join(PodcastTag, Podcast.id == PodcastTag.podcast_id)
        .join(Tag, Tag.id == PodcastTag.tag_id)
        .where(Tag.code == code)
        .limit(count)
    )

    # Execute the query
    result = session.exec(statement)

    # Fetch all results
    podcasts = result.all()

    # Return a list of Podcast models without the id field
    return [podcast.model_dump(exclude={"id"}) for podcast in podcasts]


async def get_podcast_data(url: str, session: Session) -> Podcast:
    """
    Retrieve data for podcasts
    """
    podcast = session.exec(select(Podcast).where(Podcast.xml_url == url)).first()
    if podcast:
        return podcast.model_dump(exclude={"id"})
    else:
        raise HTTPException(
            status_code=status.HTTP_404_NOT_FOUND, detail="Podcast not found"
        )


async def get_episode_data(episode_url: str, session: Session) -> Episode:
    """Retrieve episdod data"""
    episode = session.exec(
        select(Episode).where(Episode.audio_url == episode_url)
    ).first()
    if episode:
        return episode.model_dump(exclude={"id", "podcast_id"})
    else:
        raise HTTPException(
            status_code=status.HTTP_404_NOT_FOUND, detail="Episode not found"
        )


async def get_podcasts_by_query(query: str, session: Session) -> list[PodcastOut]:
    """Search podcasts by query"""
    podcasts = session.exec(
        select(Podcast).filter(
            or_(Podcast.author.contains(query), Podcast.description.contains(query))
        )
    ).all()
    return [PodcastOut(**podcast.model_dump(exclude={"id"})) for podcast in podcasts]
