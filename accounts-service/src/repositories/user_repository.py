import bcrypt
from typing import Optional
from src.models.user import User
from sqlalchemy.orm import Session
from typing import Iterator
from fastapi import Depends
from sqlalchemy.orm import Session

from src.services.database import get_db

class UserRepository:
    def __init__(self, db: Session):
        self._db = db

    def get_user_by_email(self, email: str) -> Optional[User]:
        return self._db.query(User).filter(User.email == email).first()
    
    def get_user_by_id(self, user_id: int) -> Optional[User]:
        return self._db.query(User).filter(User.user_id ==  user_id).first()

    def get_user_by_username(self, username: str) -> Optional[User]:
        return self._db.query(User).filter(User.username == username).first()

    def create_user(self, username: str, email: str,  password: str) -> User:
        password_hash = bcrypt.hashpw(password.encode(), bcrypt.gensalt())

        user = User(
            username=username,
            email=email,
            password_hash=password_hash
            )
        self._db.add(user)
        self._db.commit()
        self._db.refresh(user)
    
        return user

    def verify_email_by_user_id(self, user_id: int):
        user = self._db.query(User).filter(User.user_id == user_id).first()
        user.is_email_verified = True
        self._db.commit()

    def is_password_valid(self, password: str, password_hash: str) -> bool:
        return bcrypt.checkpw(password.encode(), password_hash.encode())


def get_user_repository(db: Session = Depends(get_db)) -> Iterator[UserRepository]:
    yield UserRepository(db)