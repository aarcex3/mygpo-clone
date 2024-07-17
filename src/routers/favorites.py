from fastapi import APIRouter
from fastapi.responses import ORJSONResponse

router = APIRouter(prefix="/favorites", tags=["Favorites"])


@router.get("/{username}.json")
async def favorite_episodes(username: str):
    """
    Get the list of favorite episodes for a given user.

    Args:
        username (str): The username of the user whose favorite episodes are to be retrieved.

    Returns:
        ORJSONResponse: A list of the user's favorite episodes.
    """
    pass
