from fastapi import APIRouter
from fastapi.responses import ORJSONResponse

"""
Subscriptions routes
"""

router = APIRouter(prefix="/subscriptions", tags=["Subscriptions"])


@router.get("/{username}/{device_id}.{format}")
async def get_device_subscriptions(username: str, device_id: str, format: str):
    """
    Get the subscriptions for a specific device belonging to a user.

    Args:
        username (str): The username of the user.
        device_id (str): The ID of the device.
        format (str): The format for the response (e.g., JSON).

    Returns:
        ORJSONResponse: A list of subscriptions for the specified device.
    """
    pass


@router.get("/{username}.{format}")
async def get_user_subscriptions(username: str, format: str):
    """
    Get all subscriptions for a given user.

    Args:
        username (str): The username of the user.
        format (str): The format for the response (e.g., JSON).

    Returns:
        ORJSONResponse: A list of subscriptions for the user.
    """
    pass


@router.put("/{username}/{device_id}.{format}")
async def update_device_subscriptions(username: str, device_id: str, format: str):
    """
    Update the subscriptions for a specific device belonging to a user.

    Args:
        username (str): The username of the user.
        device_id (str): The ID of the device.
        format (str): The format for the response (e.g., JSON).

    Returns:
        ORJSONResponse: The updated list of subscriptions for the device.
    """
    pass


@router.post("/{username}/{device_id}.json")
async def create_subscription_changes(username: str, device_id: str):
    """
    Create changes to the subscriptions for a specific device.

    Args:
        username (str): The username of the user.
        device_id (str): The ID of the device.

    Returns:
        ORJSONResponse: Confirmation of the subscription changes.
    """
    pass


@router.get("/{username}/{device_id}.json")
async def get_subscription_changes(username: str, device_id: str):
    """
    Get changes to the subscriptions for a specific device.

    Args:
        username (str): The username of the user.
        device_id (str): The ID of the device.

    Returns:
        ORJSONResponse: The subscription changes for the device.
    """
    pass
