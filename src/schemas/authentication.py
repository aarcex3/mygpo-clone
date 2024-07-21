"""
Schemas for the authentication routes
"""

from dataclasses import dataclass

from fastapi import Form
from pydantic import EmailStr


@dataclass
class RegistrationSchema:
    """
    Schema for user registration
    """

    username: str = Form(...)
    password: str = Form(...)
    email: EmailStr = Form(...)
