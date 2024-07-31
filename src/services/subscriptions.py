from sqlmodel import Session, and_, select

from src.models.device import Device
from src.models.podcast import Podcast
from src.models.subscription import Subscription
from src.schemas.podcast import PodcastOut
from src.schemas.subscriptions import SubscriptionChange
from src.utils.subscriptions import add_podcasts, remove_podcasts


async def get_device_subscriptions(
    user_id: int, device_id: int, session: Session
) -> list[PodcastOut]:
    podcasts: list[Podcast] = session.exec(
        select(Podcast)
        .join(Subscription, Subscription.podcast_id == Podcast.id)
        .join(Device, Device.id == Subscription.device_id)
        .where(Device.id == device_id)
        .where(Device.owner_id == user_id)
    ).all()
    return [PodcastOut(**podcast.model_dump(exclude={"id"})) for podcast in podcasts]


async def get_user_subscriptions(user_id: int, session: Session) -> list[PodcastOut]:
    podcasts: list[Podcast] = session.exec(
        select(Podcast)
        .join(Subscription, Subscription.podcast_id == Podcast.id)
        .join(Device, Device.id == Subscription.device_id)
        .where(Device.owner_id == user_id)
    ).all()
    return [PodcastOut(**podcast.model_dump(exclude={"id"})) for podcast in podcasts]


# TODO
# How to add and remove podcasts id from the subscription table? 
async def make_subscriptions_changes(
    device_id: int, user_id: int, changes: SubscriptionChange, session: Session
):
    device: Device = session.exec(
        select(Device).where(and_(Device.id == device_id, Device.owner_id == user_id))
    ).one()
    add_podcasts(new_podcasts=changes.add, session=session)
    remove_podcasts(old_podcasts=changes.remove, session=session)
