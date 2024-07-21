"""
Authentication Routes
"""

from authx import AuthXDependency
from fastapi import APIRouter, Depends, Form, HTTPException, Response, status
from fastapi.security import HTTPBasic, HTTPBasicCredentials
from sqlmodel import Session

from src.crud.authentication import create_user, find_user
from src.database import get_session
from src.dependecies import SECURITY
from src.schemas.authentication import RegistrationSchema
from src.utils.authentication import check_password

router = APIRouter(prefix="/auth", tags=["Auth"])


@router.post("/register")
async def register(
    form: RegistrationSchema = Depends(), session: Session = Depends(get_session)
):
    """
    User registration endpoint.

    This endpoint allows a new user to register by providing a username, password,
    and email. The registration details are validated and stored in the database.

    Parameters:
    - form (RegistrationSchema): The registration form data including username, password, and email.
    - session (Session): The database session dependency.

    Returns:
    - JSON response indicating the success or failure of the registration attempt.
    """
    return await create_user(form=form, session=session)


@router.post("/login")
async def login(
    credentials: HTTPBasicCredentials = Depends(HTTPBasic()),
    session: Session = Depends(get_session),
):
    """
    User login endpoint.

    This endpoint allows a user to log in using HTTP Basic Authentication. The credentials
    are provided in the request's Authorization header.

    Parameters:
    - credentials: HTTPBasicCredentials: The username and password provided by the user.

    Returns:
    - Response object with the access token or HTTPException if the credentials aren't valid
    """
    user = await find_user(credentials.username, session)
    if user and check_password(credentials.password, user.password):
        token = SECURITY.create_access_token(uid=str(user.id))
        headers = {"Authorization": f"Bearer {token}"}
        return Response(
            status_code=status.HTTP_200_OK,
            headers=headers,
            content="Succesfully logged in",
        )
    raise HTTPException(
        status_code=status.HTTP_401_UNAUTHORIZED, detail="Bad credentials"
    )


@router.post("/logout", dependencies=[Depends(SECURITY.access_token_required)])
async def logout(
    deps: AuthXDependency = Depends(SECURITY.get_dependency),
):
    """
    User logout endpoint.

    This endpoint allows a user to log out.

    Parameters:
    - access_token (str): The authorization token set in the login endpoint

    Returns:
    - JSON response indicating success or failure of logout attempt.
    """

    deps.unset_access_cookies()
    return {"message": "Logout successful"}
