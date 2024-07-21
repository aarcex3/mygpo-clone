"""Definition for link table for user podcast list and podcast"""

from typing import Optional

from sqlmodel import Field, SQLModel


class PodcastList(SQLModel, table=True):
    """Link table for user podcast lists and podcast"""

    user_list_id: Optional[int] = Field(foreign_key="userlist.id", primary_key=True)
    podcast_id: Optional[int] = Field(foreign_key="podcast.id", primary_key=True)
