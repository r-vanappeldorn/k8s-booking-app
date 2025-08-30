from fastapi import FastAPI, APIRouter
from src.controllers.health import router as health_router

def init_app() -> FastAPI:
    app = FastAPI(title="Accounts service", version="1.0.0")
    api_router = APIRouter(prefix="/api/accounts")

    api_router.include_router(health_router)

    app.include_router(api_router)

    return app