"""Defintion for the user table"""

import uuid
from typing import Optional

from pydantic import EmailStr
from sqlmodel import Field, Relationship, SQLModel


class User(SQLModel, table=True):
    """User in DB model"""

    id: Optional[uuid.UUID] = Field(default_factory=uuid.uuid1, primary_key=True)
    username: str = Field(index=True, unique=True)
    password: str
    email: EmailStr
    devices: list["Device"] = Relationship(back_populates="owner")  # type: ignore
    podcast_lists: Optional[list["UserList"]] = Relationship(back_populates="owner")  # type: ignore
