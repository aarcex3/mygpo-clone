"""
Directory Routes
"""

from fastapi import APIRouter, Depends
from sqlmodel import Session

from src.database import get_session
from src.services.directory import (
    get_episode_data,
    get_podcast_data,
    get_podcast_tags,
    get_podcasts_by_query,
    get_top_podcasts,
    get_top_tags,
)
from src.utils.responses import format_query_response

router = APIRouter(tags=["Directory"])


@router.get("/tags/{count}")
async def top_tags(count: int = 3, session: Session = Depends(get_session)):
    """
    Get the top tags based on usage count.

    Args:
        count (int): The number of top tags to retrieve.

    Returns:
        A list of the top tags.
    """
    return await get_top_tags(count=count, session=session)


@router.get("/tags/{code}/{count}")
async def podcasts_tags(code: str, count: int, session: Session = Depends(get_session)):
    """
    Get the top podcasts associated with a specific tag.

    Args:
        tag (str): The tag to filter podcasts by.
        count (int): The number of podcasts to retrieve.

    Returns:
        A list of podcasts associated with the tag.
    """
    return await get_podcast_tags(code=code, count=count, session=session)


@router.get("/data/podcast")
async def podcast_data(url: str, session: Session = Depends(get_session)):
    """
    Get data for a specific podcast by its URL.

    Args:
        url (str): The URL of the podcast.

    Returns:
        Data of the specified podcast.
    """
    return await get_podcast_data(url=url, session=session)


@router.get("/data/episode")
async def episode_data(episode_url: str, session: Session = Depends(get_session)):
    """
    Get data for a specific episode of a podcast by its url.

    Args:
        episode_url (str): The URL of the episode.

    Returns:
        Data of the specified episode.
    """
    return await get_episode_data(episode_url=episode_url, session=session)


@router.get("/top/{count}.{search_format}")
async def top_podcasts(
    count: int, search_format: str, session: Session = Depends(get_session)
):
    """
    Get a toplist of podcasts in a specified format.

    Args:
        count (int): The number of podcasts to retrieve.
        search_format (str): The format of the toplist (e.g., XML, JSON).

    Returns:
        A toplist of podcasts.
    """
    podcasts = await get_top_podcasts(count, session)
    return format_query_response(podcasts, search_format)


@router.get("/search.{search_format}")
async def podcast_search(
    query: str, search_format: str, session: Session = Depends(get_session)
):
    """
    Search for podcasts in a specified format.

    Args:
        query (str): The search query
        search_format (str): The format of the search results (e.g., JSON).

    Returns:
        The search results for podcasts.
    """
    podcasts = await get_podcasts_by_query(query, session)
    return format_query_response(podcasts, search_format)
