"""
Main entry point for the app
"""

import uvicorn
from fastapi import FastAPI

from src.database import create_db_and_tables
from src.dependecies import SECURITY
from src.routers import (
    authentication,
    device,
    device_sync,
    directory,
    favorites,
    lists,
    subscriptions,
)


def create_app() -> FastAPI:
    """Create instance of app"""
    _app = FastAPI(title="Gpodder.net clone", version="0.1.0")

    routers = [
        authentication.router,
        device.router,
        device_sync.router,
        directory.router,
        favorites.router,
        lists.router,
        subscriptions.router,
    ]

    for router in routers:
        _app.include_router(router, prefix="/api/v2")

    return _app


if __name__ == "__main__":
    app = create_app()
    SECURITY.handle_errors(app)
    create_db_and_tables()
    uvicorn.run("__main__:app", host="0.0.0.0", port=8000)
