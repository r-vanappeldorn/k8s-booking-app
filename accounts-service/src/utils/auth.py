from datetime import datetime, timedelta, timezone

from jose import jwt
import os

SECRET_KEY = os.environ['JWT_SECRET_KEY']

JWT_VERIFY_EMAIL = 'verify_emailuser_id'

def create_access_token(data: dict):
    to_encode = data.copy()
    expire = datetime.now(timezone.utc) + timedelta(minutes=60)
    to_encode.update({"exp": expire})
    encoded_jwt = jwt.encode(to_encode, SECRET_KEY, algorithm="HS256")

    return encoded_jwt