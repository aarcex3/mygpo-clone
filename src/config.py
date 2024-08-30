from __future__ import annotations

from functools import lru_cache

from pydantic_settings import BaseSettings


class Settings(BaseSettings):
    """App settings"""

    SECRET_KEY: str = "SECRET_KEY"
    DB_URL: str = "DB_URL"
    REDIS_URL: str = "REDIS_URL"

    class Config:
        """Settings config"""

        env_file = ".env"


@lru_cache(maxsize=1, typed=True)
def get_settings() -> Settings:
    """Get an instance of Settings class"""
    return Settings()
