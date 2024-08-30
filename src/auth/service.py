import uuid
from typing import Any, Dict

from litestar import Response, status_codes
from litestar.exceptions import HTTPException
from litestar.security.jwt import Token

from src.auth.schemas import LoginForm, RegistrationForm
from src.auth.utils import check_password
from src.person.models import Person
from src.person.service import PersonService
from src.security import AUTH


class AuthService:
    """Authentication Service"""

    def __init__(self, person_service: PersonService):
        self.person_service: PersonService = person_service

    async def register(self, form: RegistrationForm) -> Dict:
        """Register new person service"""
        person: Person = await self.person_service.create_person_obj(
            username=form.username,
            password=form.password.get_secret_value(),
            email=form.email,
        )
        return await self.person_service.add(person=person)

    async def login(self, data: LoginForm) -> Response[Any]:
        """Log in existing person service"""
        person: Person | None = await self.person_service.find_user_by_username(
            username=data.username
        )
        if person and check_password(data.password.get_secret_value(), person.password):
            return AUTH.login(
                identifier=str(person.id),
                token_extras={"email": person.email},
                response_body={"message": "Logged In"},
                response_status_code=status_codes.HTTP_200_OK,
            )
        raise HTTPException(
            status_code=status_codes.HTTP_401_UNAUTHORIZED, detail="Bad credentials"
        )

    async def logout(self, token: Token) -> Response[Any]:
        """Log out authenticated person service"""
        person: Person | None = await self.person_service.find_user_by_id(
            person_id=uuid.UUID(token.sub)
        )
        if person:
            print("logging out")
            token_jti = token.jti
            expiration = token.exp - token.iat

            return Response(content={"message": "Logged Out"})
        else:
            raise HTTPException(
                status_code=status_codes.HTTP_404_NOT_FOUND, detail="Person not found"
            )

    async def __call__(self, *args: Any, **kwds: Any) -> Any:
        return self
