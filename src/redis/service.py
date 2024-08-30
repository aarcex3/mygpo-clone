# from collections.abc import AsyncGenerator
# from typing import Any

# import aioredis

# from src.config import get_settings

# settings = get_settings()


# async def init_redis() -> aioredis.Redis:
#     """Initialize Redis pool"""
#     return await aioredis.from_url(settings.REDIS_URL)


# async def provide_redis_service() -> AsyncGenerator[aioredis.Redis, Any]:
#     """Dependency function to get Redis client"""
#     redis = await init_redis()
#     try:
#         yield redis
#     finally:
#         await redis.close()
