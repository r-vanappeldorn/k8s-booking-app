from fastapi.testclient import TestClient
from src.services.app import init_app

app = init_app()
client = TestClient(app)

class TestHealth:
    def test_health_route_should_return_200(self):
        response = client.get("/api/accounts/health")

        assert response.status_code == 200
        assert response.json() == {"status": "ok"}