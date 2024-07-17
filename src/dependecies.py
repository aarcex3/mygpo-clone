"""
App dependecies
"""

import os

from pydantic_settings import BaseSettings


class Settings(BaseSettings):
    """App settings"""

    DB_URL: str = str(os.getenv("DB_URL"))
    SECRET_KEY: str = str(os.getenv("SECRET_KEY", "secret"))


SETTINGS = Settings()
