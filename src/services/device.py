"""
Utilities and crud methods for Device endpoint
"""

from uuid import UUID

from fastapi import HTTPException, Response, status
from sqlmodel import Session, select

from src.models.device import Device
from src.schemas.device import UpdateDeviceSchema


async def get_user_devices(user_id: int, session: Session) -> list[Device]:
    """
    Get all user devices
    """
    devices = session.exec(select(Device).where(Device.owner_id == user_id)).all()
    return [device.model_dump(exclude={"owner_id"}) for device in devices]


async def update_device_info(
    device_id: int, data: UpdateDeviceSchema, owner_id: int, session: Session
):
    """
    Update device info (caption, type)
    """
    device = session.exec(
        select(Device).where(Device.id == device_id).where(Device.owner_id == owner_id)
    ).one()

    device.caption, device.device_type = data.caption, data.device_type
    session.add(device)
    try:
        session.commit()
        return Response(status_code=status.HTTP_200_OK, content="Device updated")
    except Exception as ex:
        raise HTTPException(
            status_code=status.HTTP_422_UNPROCESSABLE_ENTITY, detail="Bad data"
        ) from ex
