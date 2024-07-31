import pytest
from fastapi.testclient import TestClient
from sqlmodel import Session, SQLModel, StaticPool, create_engine

from src.database import get_session
from src.main import create_app
from src.utils.authentication import check_password

app = create_app()


@pytest.fixture(name="session")
def session_fixture():
    engine = create_engine(
        "sqlite:///testing.sqlite",
        connect_args={"check_same_thread": False},
        poolclass=StaticPool,
    )
    SQLModel.metadata.create_all(engine)
    with Session(engine) as session:
        yield session


@pytest.fixture(name="client")
def client_fixture(session: Session):
    app.dependency_overrides[get_session] = lambda: session
    client = TestClient(app)
    yield client
    app.dependency_overrides.clear()


@pytest.mark.asyncio
async def test_register_user(client: TestClient):
    register_data = {
        "username": "test_user",
        "email": "test@example.com",
        "password": "test_password",
    }
    response = client.post("/api/v2/auth/register", data=register_data)
    assert response.status_code == 201


@pytest.mark.asyncio
async def test_register_user_with_existing_username(client: TestClient):
    register_data = {
        "username": "test_user",
        "email": "test@example.com",
        "password": "test_password",
    }
    response = client.post("/api/v2/auth/register", data=register_data)
    assert response.status_code == 409
    assert response.json() == {"detail": "Username already exists"}


headers = None


@pytest.mark.asyncio
async def test_login_user(client: TestClient):
    response = client.post("/api/v2/auth/login", auth=("test_user", "test_password"))
    assert response.status_code == 200
    assert "Authorization" in response.headers
    global headers
    headers = {"Authorization": response.headers["Authorization"]}


@pytest.mark.asyncio
async def test_login_user_bad_credentials(client: TestClient):
    username = "test_user_bad"
    password = "test_password_bad"
    response = client.post("/api/v2/auth/login", auth=(username, password))
    assert response.status_code == 401
    assert "Authorization" not in response.headers
    assert response.json() == {"detail": "Bad credentials"}


@pytest.mark.asyncio
async def test_logout(client: TestClient):
    assert headers is not None
    response = client.post("/api/v2/auth/logout", headers=headers)
    assert response.status_code == 200
    assert "Authorization" not in response.headers
    assert response.content.decode("utf-8") == "Logout successful"
