from authx import TokenPayload
from fastapi import APIRouter, Depends, HTTPException, status
from sqlmodel import Session

from src.database import get_session
from src.dependecies import SECURITY
from src.schemas.lists import CreateList, ListOut, UpdateList
from src.services.lists import (
    create_podcasts_list,
    delete_user_list,
    get_user_list,
    get_user_lists,
    update_user_list,
)
from src.utils.responses import format_query_response

"""
Podcasts lists routes
"""

router = APIRouter(prefix="/lists", tags=["Podcasts lists"])


@router.post(
    "/{username}/create", dependencies=[Depends(SECURITY.access_token_required)]
)
async def podcast_list(
    data: CreateList,
    username: str,
    payload: TokenPayload = Depends(SECURITY.access_token_required),
    session: Session = Depends(get_session),
):
    """
    Create a new podcast list for a given user.

    Args:
        title (str): The title of the podcast list to be created.
        username (str): The username of the user creating the list.

    Returns:
        Message confirming the creation of the list
    """
    return await create_podcasts_list(
        title=data.title, username=username, owner_id=int(payload.sub), session=session
    )


@router.get("/{username}", dependencies=[Depends(SECURITY.access_token_required)])
async def user_lists(
    payload: TokenPayload = Depends(SECURITY.access_token_required),
    session: Session = Depends(get_session),
):
    """
    Get all podcast lists for a given user.

    Args:
        username (str): The username of the user whose podcast lists are to be retrieved.

    Returns:
        list[UserList]
    """
    return await get_user_lists(user_id=int(payload.sub), session=session)


@router.get(
    "/{username}/lists/{listname}.{search_format}",
    dependencies=[Depends(SECURITY.access_token_required)],
)
async def user_list(
    listname: str,
    username: str,
    search_format: str,
    payload: TokenPayload = Depends(SECURITY.access_token_required),
    session: Session = Depends(get_session),
):
    """
    Get a specific podcast list for a given user.

    Args:
        listname (str): The name of the podcast list to retrieve.
        username (str): The username of the user.
        format (str): The format for the response (e.g., JSON).

    Returns:
        list[ListOut]
    """
    user_list: ListOut = await get_user_list(
        listname=listname, user_id=int(payload.sub), session=session
    )
    match search_format:
        case "xml":
            return user_list.to_xml()
        case "opml":
            return user_list.to_opml()
        case "json":
            return user_list
        case _:
            raise HTTPException(
                detail="Format not supported",
                status_code=status.HTTP_406_NOT_ACCEPTABLE,
            )


@router.put(
    "/{username}/lists/{listname}",
    dependencies=[Depends(SECURITY.access_token_required)],
)
async def user_list(
    listname: str,
    username: str,
    new_list: UpdateList,
    payload: TokenPayload = Depends(SECURITY.access_token_required),
    session: Session = Depends(get_session),
):
    """
    Update a specific podcast list for a given user.

    Args:
        listname (str): The name of the podcast list to update.
        username (str): The username of the user.
        format (str): The format for the response (e.g., JSON).

    Returns:
        ORJSONResponse: The updated podcast list.
    """
    return await update_user_list(
        listname=listname,
        username=username,
        user_id=int(payload.sub),
        new_list=new_list,
        session=session,
    )


@router.delete(
    "/{username}/lists/{listname}",
    dependencies=[Depends(SECURITY.access_token_required)],
)
async def user_list(
    listname: str,
    username: str,
    payload: TokenPayload = Depends(SECURITY.access_token_required),
    session: Session = Depends(get_session),
):
    """
    Delete a specific podcast list for a given user.

    Args:
        listname (str): The name of the podcast list to delete.
        username (str): The username of the user.
        format (str): The format for the response (e.g., JSON).

    Returns:
        Confirmation of the deletion.
    """
    return await delete_user_list(
        listname=listname, user_id=int(payload.sub), session=session
    )
