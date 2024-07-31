from fastapi import APIRouter, Depends
from sqlmodel import Session

from src.database import get_session

router = APIRouter(prefix="/favorites", tags=["Favorites"])


@router.get("/{username}")
async def favorite_episodes(username: str,session: Session = Depends(get_session)):
    """
    Get the list of favorite episodes for a given user.

    Args:
        username (str): The username of the user whose favorite episodes are to be retrieved.

    Returns:
        ORJSONResponse: A list of the user's favorite episodes.
    """
    pass
