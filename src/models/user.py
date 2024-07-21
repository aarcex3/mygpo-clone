"""Defintion for the user table"""

from typing import Optional

from pydantic import EmailStr
from sqlmodel import Field, Relationship, SQLModel


class User(SQLModel, table=True):
    """User in DB model"""

    id: Optional[int] = Field(default=None, primary_key=True, index=True)
    username: str = Field(index=True, unique=True)
    password: str
    email: EmailStr
    devices: Optional[list["Device"]] = Relationship(back_populates="owner")  # type: ignore
    podcast_lists: Optional[list["UserList"]] = Relationship(back_populates="owner")  # type: ignore
