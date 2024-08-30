from re import U
from typing import Any, Dict, Union
from uuid import UUID

from advanced_alchemy.service import SQLAlchemyAsyncRepositoryService
from litestar import status_codes
from litestar.exceptions import HTTPException
from sqlalchemy.ext.asyncio import AsyncSession

from src.auth.utils import hash_password
from src.person.models import Person
from src.person.repository import PersonRepository


class PersonService(SQLAlchemyAsyncRepositoryService[Person]):
    """User service"""

    repository_type = PersonRepository

    def __init__(self, session: AsyncSession, **repo_kwargs: Any) -> None:
        self.repository: PersonRepository = self.repository_type(
            session=session, **repo_kwargs
        )  # type: ignore
        self.model_type = self.repository.model_type

    async def add(self, person: Person) -> Dict[str, Any]:
        """Add new user to the database

        Args:
            user (User): User object

        Raises:
            HTTPException: If a user already exists in the database

        Returns:
            Dict[str, Any]: Message saying that the user was succesfully created
        """
        try:
            await self.repository.add(person)
            return {"message": f"User {person.username} succesfully created"}
        except Exception as ex:
            raise HTTPException(
                status_code=status_codes.HTTP_409_CONFLICT,
                detail=f"Username already exists. Error: {ex}",
            ) from ex

    async def find_user_by_id(self, person_id: UUID) -> Union[Person, None]:
        """Find a user given its id

        Args:
            user_id (UUID): The user id

        Returns:
            User | None: A user object if the user was found, else None
        """
        return await super().get_one_or_none(Person.id == person_id)

    async def find_user_by_username(self, username: str) -> Union[Person, None]:
        """Find a user given its username

        Args:
            username (str): The username used to search the user

        Returns:
            Union[User, None]: A user object if found, else None
        """
        return await super().get_one_or_none(Person.username == username)

    async def create_person_obj(
        self, username: str, password: str, email: str
    ) -> Person:
        """Create user object with the given attributes

        Args:
            username (str): The unique username for the user
            password (str): The password to be hashed
            email (str): The unique email address of the user

        Returns:
            User: User object to add to the db
        """ """"""
        return Person(username=username, password=hash_password(password), email=email)
