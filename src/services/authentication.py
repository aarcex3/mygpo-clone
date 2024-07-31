"""
Utilities and crud methods for Auth endpoint
"""

from typing import Union

from fastapi import HTTPException, Response, status
from fastapi.security import HTTPBasicCredentials
from sqlmodel import Session, select

from src.dependecies import SECURITY
from src.models import User
from src.schemas.authentication import RegistrationSchema
from src.utils.authentication import check_password, hash_password
from src.utils.device import create_user_device


async def create_user(
    form: RegistrationSchema, session: Session
) -> Union[Response, HTTPException]:
    """
    Register a new user
    """
    new_user = User(
        username=form.username, password=hash_password(form.password), email=form.email
    )
    device = create_user_device(new_user)
    new_user.devices.append(device)
    session.add(device)
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


def authenticate_user(user: User, password: str):
    """
    Authenticate user and generate access token
    """
    if user and check_password(password, user.password):
        token = SECURITY.create_access_token(uid=str(user.id), username=user.username)
        headers = {"Authorization": f"Bearer {token}"}
        return Response(
            status_code=status.HTTP_200_OK,
            headers=headers,
            content="Succesfully logged in",
        )
    raise HTTPException(
        status_code=status.HTTP_401_UNAUTHORIZED, detail="Bad credentials"
    )
