from sqlmodel import Session, select

from src.models.podcast import Podcast
from src.models.subscription import Subscription


def add_podcasts(new_podcasts: list[str], session: Session):
    pass


def remove_podcasts(old_podcasts: list[str], session: Session):
    pass
