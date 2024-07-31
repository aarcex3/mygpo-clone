from authx import TokenPayload
from fastapi import APIRouter, Depends
from sqlmodel import Session

from src.database import get_session
from src.dependecies import SECURITY
from src.schemas.subscriptions import SubscriptionChange
from src.services.subscriptions import (
    get_device_subscriptions,
    get_user_subscriptions,
    make_subscriptions_changes,
)
from src.utils.responses import format_podcasts_response

"""
Subscriptions routes
"""

router = APIRouter(prefix="/subscriptions", tags=["Subscriptions"])


@router.get(
    "/{username}/{device_id}.{search_format}",
    dependencies=[Depends(SECURITY.access_token_required)],
)
async def device_subscriptions(
    username: str,
    device_id: str,
    search_format: str,
    payload: TokenPayload = Depends(SECURITY.access_token_required),
    session: Session = Depends(get_session),
):
    """
    Get the subscriptions for a specific device belonging to a user.

    Args:
        username (str): The username of the user.
        device_id (str): The ID of the device.
        format (str): The format for the response (e.g., JSON).

    Returns:
        A list of subscriptions for the specified device.
    """
    podcasts = await get_device_subscriptions(
        user_id=payload.sub, device_id=device_id, session=session
    )
    return format_podcasts_response(podcasts, search_format)


@router.get(
    "/{username}.{search_format}",
    dependencies=[Depends(SECURITY.access_token_required)],
)
async def user_subscriptions(
    search_format: str,
    payload: TokenPayload = Depends(SECURITY.access_token_required),
    session: Session = Depends(get_session),
):
    """
    Get all subscriptions for a given user.

    Args:
        username (str): The username of the user.
        format (str): The format for the response (e.g., JSON).

    Returns:
        A list of subscriptions for the user.
    """
    podcasts = await get_user_subscriptions(user_id=int(payload.sub), session=session)
    return format_podcasts_response(podcasts, search_format)


@router.put("/{username}/{device_id}.{search_format}")
async def device_subscriptions(username: str, device_id: str, search_format: str):
    """
    Update the subscriptions for a specific device belonging to a user.

    Args:
        username (str): The username of the user.
        device_id (str): The ID of the device.
        format (str): The format for the response (e.g., JSON).

    Returns:
        The updated list of subscriptions for the device.
    """
    pass


@router.post(
    "/{username}/{device_id}",
    dependencies=[Depends(SECURITY.access_token_required)],
)
async def subscription_changes(
    device_id: str,
    changes: SubscriptionChange,
    payload: TokenPayload = Depends(SECURITY.access_token_required),
    session: Session = Depends(get_session),
):
    """
    Create changes to the subscriptions for a specific device.

    Args:
        username (str): The username of the user.
        device_id (str): The ID of the device.

    Returns:
        Confirmation of the subscription changes.
    """
    result = await make_subscriptions_changes(
        device_id=device_id, changes=changes, user_id=int(payload.sub), session=session
    )
    return result


@router.get("/{username}/{device_id}")
async def subscription_changes(username: str, device_id: str):
    """
    Get changes to the subscriptions for a specific device.

    Args:
        username (str): The username of the user.
        device_id (str): The ID of the device.

    Returns:
        The subscription changes for the device.
    """
    pass
