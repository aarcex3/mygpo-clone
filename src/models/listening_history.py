"""Definition for listening history table"""

import uuid
from datetime import datetime
from typing import Optional

from sqlmodel import Field, SQLModel


class ListeningHistory(SQLModel, table=True):
    """LH in DB model"""

    id: Optional[int] = Field(default=None, index=True, primary_key=True)
    user_id: Optional[int] = Field(foreign_key="user.id")
    episode_id: Optional[int] = Field(foreign_key="episode.id")
    progress: float
    date: datetime
