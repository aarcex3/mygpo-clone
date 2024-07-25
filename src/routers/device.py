"""
Devices routes 
"""

from authx import TokenPayload
from fastapi import APIRouter, Depends
from sqlmodel import Session

from src.database import get_session
from src.dependecies import SECURITY
from src.schemas.device import UpdateDeviceSchema
from src.services.device import get_user_devices, update_device_info

router = APIRouter(prefix="/devices", tags=["Devices"])


@router.post(
    "/{username}/{device_id}", dependencies=[Depends(SECURITY.access_token_required)]
)
async def device_info(
    device_id: str,
    data: UpdateDeviceSchema,
    payload: TokenPayload = Depends(SECURITY.access_token_required),
    session: Session = Depends(get_session),
):
    """
    Update the information of a specific device for a given user.

    Args:
        - username (str): The username of the user.
        - device_id (str): The ID of the device to update.

    Returns:
        - The result of the update operation.
    """
    return await update_device_info(
        device_id=device_id, data=data, owner_id=payload.sub, session=session
    )


@router.post("/{username}", dependencies=[Depends(SECURITY.access_token_required)])
async def user_devices(
    payload: TokenPayload = Depends(SECURITY.access_token_required),
    session: Session = Depends(get_session),
):
    """
    Retrieve the list of devices for a given user.

    Args:
        username (str): The username of the user.

    Returns:
        The list of devices associated with the user.
    """
    return await get_user_devices(user_id=payload.sub, session=session)


@router.get("/{username}/{device_id}")
async def device_update(username: str, device_id: str):
    """
    Get the update status or information of a specific device for a given user.

    Args:
        username (str): The username of the user.
        device_id (str): The ID of the device to get the update status for.

    Returns:
        The update status or information of the device.
    """
    pass
