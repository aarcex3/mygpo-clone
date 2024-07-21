"""
App dependecies
"""

import os

from pydantic_settings import BaseSettings
from authx import AuthX, AuthXConfig


class Settings(BaseSettings):
    """App settings"""

    DB_URL: str = str(os.getenv("DB_URL"))
    SECRET_KEY: str = str(os.getenv("SECRET_KEY", "secret"))


SETTINGS = Settings()


config = AuthXConfig()
config.JWT_ALGORITHM = "HS256"
config.JWT_SECRET_KEY = SETTINGS.SECRET_KEY
config.JWT_TOKEN_LOCATION = ["headers"]

SECURITY = AuthX(config=config)
