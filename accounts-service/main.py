import uvicorn
from fastapi import FastAPI
from src.controllers.health import router as health_router

def init_app() -> FastAPI:
    app = FastAPI(title="Accounts service", version="1.0.0")
    app.include_router(health_router, prefix="/health")

    return app

app = init_app()

if __name__ == "__main__":
    uvicorn.run("main:app", host="0.0.0.0", port=8000, reload=True)