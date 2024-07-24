from src.models.device import Device
from src.models.user import User


def create_user_device(
    user: User, caption: str = "Server", device_type: str = "SERVER"
) -> Device:
    """
    Create a device for the given user
    """
    device = Device(
        owner=user, owner_id=user.id, device_type=device_type, caption=caption
    )
    return device
