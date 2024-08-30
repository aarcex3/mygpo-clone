from typing import Annotated, Any, Dict, Sequence

from litestar import Request, Response, post, status_codes
from litestar.controller import Controller
from litestar.enums import MediaType, RequestEncodingType
from litestar.params import Body
from litestar.security.jwt import Token

from src.auth import urls
from src.auth.schemas import LoginForm, RegistrationForm
from src.auth.service import AuthService
from src.person.models import Person


class AuthController(Controller):
    """Authentication controller"""

    tags: Sequence[str] = ["auth", "authentication"]

    @post(
        urls.REGISTRATION,
        status_code=status_codes.HTTP_201_CREATED,
        media_type=MediaType.JSON,
    )
    async def register(
        self,
        data: Annotated[
            RegistrationForm, Body(media_type=RequestEncodingType.URL_ENCODED)
        ],
        auth_service: AuthService,
    ) -> Dict:
        """Registration endpoint"""
        user = await auth_service.register(form=data)
        if user:
            return {"message": "User succesfully created"}
        else:
            return {"message": "User  was not created"}

    @post(
        urls.LOGIN,
        status_code=status_codes.HTTP_200_OK,
        media_type=MediaType.JSON,
    )
    async def login(
        self,
        data: Annotated[LoginForm, Body(media_type=RequestEncodingType.URL_ENCODED)],
        auth_service: AuthService,
    ) -> Response[Any]:
        """Login endpoint"""
        return await auth_service.login(data=data)

    @post(
        urls.LOGOUT,
        status_code=status_codes.HTTP_200_OK,
        media_type=MediaType.JSON,
    )
    async def logout(
        self,
        request: Request[Person, Token, Any],
        auth_service: AuthService,
    ) -> Response[Any]:
        """Logout endpoint"""
        return await auth_service.logout(token=request.auth)
