"""
Utilities and crud methods for Auth endpoint
"""

from typing import Union

from fastapi import HTTPException, Response, status
from sqlmodel import Session, select

from src.models.user import User
from src.schemas.authentication import RegistrationSchema
from src.utils.authentication import hash_password


async def create_user(
    form: RegistrationSchema, session: Session
) -> Union[Response, HTTPException]:
    """
    Register a new user
    """
    new_user = User(
        username=form.username, password=hash_password(form.password), email=form.email
    )
    session.add(new_user)
    try:
        session.commit()
        return Response(status_code=status.HTTP_201_CREATED, content="User created")
    except Exception as ex:
        raise HTTPException(
            status_code=status.HTTP_409_CONFLICT, detail="Username already exists"
        ) from ex


async def find_user(username: str, session: Session) -> Union[User, None]:
    """
    Find the user for the given username
    """
    user = session.exec(select(User).where(User.username == username)).first()
    return user
