"""
Main entry point for the app
"""

from fastapi import FastAPI
import uvicorn

app = FastAPI(title="Gpodder.net clone", version="0.1.0")


if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8000)
