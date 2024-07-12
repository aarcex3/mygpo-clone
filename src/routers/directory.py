from fastapi import APIRouter
from fastapi.responses import ORJSONResponse

"""
Directory Routes
"""

router = APIRouter(tags=["Directory"])


@router.get("/tags/{count}.json", response_model=ORJSONResponse)
async def top_tags(count: int):
    """
    Get the top tags based on usage count.

    Args:
        count (int): The number of top tags to retrieve.

    Returns:
        ORJSONResponse: A list of the top tags.
    """
    pass


@router.get("/tags/{tag}/{count}.json", response_model=ORJSONResponse)
async def podcasts_tags(tag: str, count: int):
    """
    Get the top podcasts associated with a specific tag.

    Args:
        tag (str): The tag to filter podcasts by.
        count (int): The number of podcasts to retrieve.

    Returns:
        ORJSONResponse: A list of podcasts associated with the tag.
    """
    pass


@router.get("/data/podcast.json", response_model=ORJSONResponse)
async def podcast_data(url: str):
    """
    Get data for a specific podcast by its URL.

    Args:
        url (str): The URL of the podcast.

    Returns:
        ORJSONResponse: Data of the specified podcast.
    """
    pass


@router.get("/data/episode.json", response_model=ORJSONResponse)
async def episode_data(podcast_url: str, episode_url: str):
    """
    Get data for a specific episode of a podcast.

    Args:
        podcast_url (str): The URL of the podcast.
        episode_url (str): The URL of the episode.

    Returns:
        ORJSONResponse: Data of the specified episode.
    """
    pass


@router.get("/toplist/{count}.{format}", response_model=ORJSONResponse)
async def podcasts_toplist(count: int, format: str):
    """
    Get a toplist of podcasts in a specified format.

    Args:
        count (int): The number of podcasts to retrieve.
        format (str): The format of the toplist (e.g., JSON).

    Returns:
        ORJSONResponse: A toplist of podcasts.
    """
    pass


@router.get("/search.{format}", response_model=ORJSONResponse)
async def podcast_search(query: str, format: str):
    """
    Search for podcasts in a specified format.

    Args:
        query (str): The search query
        format (str): The format of the search results (e.g., JSON).

    Returns:
        ORJSONResponse: The search results for podcasts.
    """
    pass
