from pydantic import BaseModel


class UpdateDeviceSchema(BaseModel):
    """
    Schema for updating a device
    """

    caption: str
    device_type: str
