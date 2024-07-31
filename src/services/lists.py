from fastapi import HTTPException, Response, status
from fastapi.responses import JSONResponse
from sqlmodel import Session, select

from src.models.user_list import UserList
from src.schemas.lists import ListOut, UpdateList


async def create_podcasts_list(
    title: str, username: str, owner_id: int, session: Session
) -> Response | HTTPException:
    name: str = title.replace(" ", "-").lower()
    url: str = f"https://gpodder-clone.net/user/{username}/lists/{name}"
    user_list: UserList = UserList(title=title, name=name, owner_id=owner_id, web=url)
    try:
        session.add(user_list)
        session.commit()
        content = f"List created, url at {url}"
        return Response(content=content, status_code=status.HTTP_201_CREATED)
    except:
        raise HTTPException(
            status_code=status.HTTP_409_CONFLICT,
            detail=f"List with name {name} already exists",
        )


async def get_user_lists(user_id: int, session: Session) -> list[ListOut]:
    user_lists = session.exec(
        select(UserList).where(UserList.owner_id == user_id)
    ).all()
    return [ListOut(**user_list.model_dump(exclude={"id"})) for user_list in user_lists]


async def get_user_list(
    listname: str, user_id: int, session: Session
) -> ListOut | HTTPException:
    user_list = session.exec(
        select(UserList)
        .where(UserList.name == listname)
        .where(UserList.owner_id == user_id)
    ).one()
    if user_list:
        return ListOut(**user_list.model_dump(exclude={"id"}))
    else:
        raise HTTPException(
            detail="List not found", status_code=status.HTTP_404_NOT_FOUND
        )


async def update_user_list(
    listname: str, username: str, user_id: int, new_list: UpdateList, session: Session
):
    user_list = session.exec(
        select(UserList)
        .where(UserList.name == listname)
        .where(UserList.owner_id == user_id)
    ).one()
    if user_list:
        user_list.title = new_list.title
        user_list.name = user_list.title.replace(" ", "-").lower()
        user_list.web = (
            f"https://gpodder-clone.net/user/{username}/lists/{user_list.name}"
        )
        try:
            session.add(user_list)
            session.commit()
            session.refresh(user_list)
            return JSONResponse(
                content=user_list.model_dump(exclude={"owner_id"}),
                status_code=status.HTTP_200_OK,
            )

        except Exception as ex:
            raise HTTPException(
                detail=ex, status_code=status.HTTP_500_INTERNAL_SERVER_ERROR
            ) from ex
    else:
        raise HTTPException(
            detail="List not found", status_code=status.HTTP_404_NOT_FOUND
        )


async def delete_user_list(
    listname: str, user_id: int, session: Session
) -> ListOut | HTTPException:
    user_list = session.exec(
        select(UserList)
        .where(UserList.name == listname)
        .where(UserList.owner_id == user_id)
    ).one()
    if user_list:
        session.delete(user_list)
        session.commit()
        return Response(
            content=f"List {listname} deleted", status_code=status.HTTP_200_OK
        )
    else:
        raise HTTPException(
            detail="List not found", status_code=status.HTTP_404_NOT_FOUND
        )
