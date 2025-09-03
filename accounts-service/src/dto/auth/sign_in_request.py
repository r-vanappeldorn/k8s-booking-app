from pydantic import BaseModel
from typing import Optional

class SignInRequest(BaseModel):
    username: Optional[str] = None
    email: Optional[str] = None
    password: str