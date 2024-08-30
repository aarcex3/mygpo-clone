import re
from typing import Annotated

from pydantic import BaseModel, EmailStr, Field, SecretStr, field_validator


class RegistrationForm(BaseModel):
    """Registration form"""

    username: str
    password: SecretStr
    email: EmailStr

    @field_validator("password")
    @classmethod
    def validate_password(cls, values):
        password = values.get_secret_value()
        if not cls.is_strong_password(password):
            raise ValueError("Password does not meet the security requirements.")
        return values

    @staticmethod
    def is_strong_password(password: str) -> bool:
        if len(password) < 8:
            return False
        if not re.search(r"[A-Z]", password):
            return False
        if not re.search(r"[a-z]", password):
            return False
        if not re.search(r"[0-9]", password):
            return False
        if not re.search(r"[@#$%^&+=]", password):
            return False
        return True


class LoginForm(BaseModel):
    """Login form"""

    username: str = Field(description="The unique username for the user")
    password: SecretStr = Field(description="Password for the user")
