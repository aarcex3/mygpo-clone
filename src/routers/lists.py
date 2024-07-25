from fastapi import APIRouter
from fastapi.responses import ORJSONResponse

"""
Podcasts lists routes
"""

router = APIRouter(prefix="/lists", tags=["Podcasts lists"])


@router.post("/{username}/create.{format}")
async def create_podcast_list(title: str, username: str, format: str):
    """
    Create a new podcast list for a given user.

    Args:
        title (str): The title of the podcast list to be created.
        username (str): The username of the user creating the list.
        format (str): The format for the response (e.g., JSON).

    Returns:
        ORJSONResponse: The details of the created podcast list.
    """
    pass


@router.get("/{username}")
async def user_podcast_lists(username: str):
    """
    Get all podcast lists for a given user.

    Args:
        username (str): The username of the user whose podcast lists are to be retrieved.

    Returns:
        ORJSONResponse: A list of the user's podcast lists.
    """
    pass


@router.get("/{username}/lists/{listname}.{format}")
async def get_user_podcast_list(listname: str, username: str, format: str):
    """
    Get a specific podcast list for a given user.

    Args:
        listname (str): The name of the podcast list to retrieve.
        username (str): The username of the user.
        format (str): The format for the response (e.g., JSON).

    Returns:
        ORJSONResponse: The specified podcast list.
    """
    pass


@router.put("/{username}/lists/{listname}.{format}")
async def update_user_podcast_list(listname: str, username: str, format: str):
    """
    Update a specific podcast list for a given user.

    Args:
        listname (str): The name of the podcast list to update.
        username (str): The username of the user.
        format (str): The format for the response (e.g., JSON).

    Returns:
        ORJSONResponse: The updated podcast list.
    """
    pass


@router.delete("/{username}/lists/{listname}.{format}")
async def delete_user_podcast_list(listname: str, username: str, format: str):
    """
    Delete a specific podcast list for a given user.

    Args:
        listname (str): The name of the podcast list to delete.
        username (str): The username of the user.
        format (str): The format for the response (e.g., JSON).

    Returns:
        Confirmation of the deletion.
    """
    pass
