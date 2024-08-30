from advanced_alchemy.base import UUIDAuditBase
from pydantic import EmailStr
from sqlalchemy.orm import Mapped, mapped_column


class Person(UUIDAuditBase):
    """User in DB model"""

    username: Mapped[str] = mapped_column(unique=True, index=True)
    password: Mapped[str]
    email: Mapped[EmailStr] = mapped_column(unique=True, index=True)
    is_active: Mapped[bool] = mapped_column(default=True, nullable=False)
