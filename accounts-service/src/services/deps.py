from typing import Iterator
from fastapi import Depends
from sqlalchemy.orm import Session
from src.services.database import get_db
from src.repositories.user_repository import UserRepository

def get_user_repository(db: Session = Depends(get_db)) -> Iterator[UserRepository]:
    yield UserRepository(db)