"""Database module"""

from sqlmodel import Session, create_engine

from src.dependecies import SETTINGS


def get_session():
    """Get database session"""
    engine = create_engine(f"sqlite:///{SETTINGS.DB_URL}", echo=True)

    with Session(engine) as session:
        yield session
    engine.dispose()
