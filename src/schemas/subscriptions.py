"""
Schemas for the subscription routes
"""

from pydantic import BaseModel


class SubscriptionChange(BaseModel):
    """
    Upload subscription changes schema
    """

    add: list[str]
    remove: list[str]
