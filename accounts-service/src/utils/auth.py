from datetime import datetime, timedelta, timezone

from jose import jwt
import os

SECRET_KEY = os.environ['JWT_SECRET_KEY']

JWT_VERIFY_EMAIL = 'verify_emailuser_id'

def create_access_token(data: dict, minutes: int = 60):
    to_encode = data.copy()
    to_encode["minutes"] = minutes
    expire = datetime.now(timezone.utc) + timedelta(minutes=minutes)
    to_encode.update({"exp": expire})
    encoded_jwt = jwt.encode(to_encode, SECRET_KEY, algorithm="HS256")

    return encoded_jwt

def decode_token(token: str):
    return jwt.decode(token, SECRET_KEY, algorithms="HS256")