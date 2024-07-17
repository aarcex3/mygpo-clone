"""Defintion for the user table"""

from typing import Optional
from sqlmodel import SQLModel, Field, Relationship
from pydantic import EmailStr


class User(SQLModel, table=True):
    """User in DB model"""

    id: Optional[int] = Field(default=None, primary_key=True, index=True)
    username: str = Field(index=True, unique=True)
    password: str
    email: EmailStr
    devices: list["Device"] = Relationship(back_populates="owner")  # type: ignore
    podcast_lists: list["UserList"] = Relationship(back_populates="owner")  # type: ignore
