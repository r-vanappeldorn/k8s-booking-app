from fastapi import FastAPI, APIRouter
from src.controllers.health import router as health_controller
from src.controllers.auth import router as auth_controller

def init_app() -> FastAPI:
    app = FastAPI(title="Accounts service", version="1.0.0")
    api_router = APIRouter(prefix="/api/accounts")

    api_router.include_router(health_controller)
    api_router.include_router(auth_controller)

    app.include_router(api_router)
    print("hello")

    for r in app.router.routes:
        print(getattr(r, "methods", None), getattr(r, "path", None))

    return app