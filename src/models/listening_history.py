"""Definition for listening history table"""

from datetime import datetime
from typing import Optional
from sqlmodel import SQLModel, Field


class ListeningHistory(SQLModel, table=True):
    """LH in DB model"""

    id: Optional[int] = Field(default=None, primary_key=True, index=True)
    user_id: Optional[int] = Field(foreign_key="user.id")
    episode_id: Optional[int] = Field(foreign_key="episode.id")
    progress: float
    date: datetime
