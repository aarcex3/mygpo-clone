"""
Main entry point for the app
"""

from fastapi import FastAPI
import uvicorn
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


app = create_app()

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8000)
