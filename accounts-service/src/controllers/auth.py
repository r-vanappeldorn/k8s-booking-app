import re 
from fastapi import APIRouter, Depends
from fastapi.responses import JSONResponse
from src.dto.auth.sign_up_request import SignUpRequest
from src.dto.common.bad_request import BadRequest
from src.services.deps import get_user_repository
from src.repositories.user_repository import UserRepository
from src.models.user import User
from src.utils.auth import create_access_token, JWT_VERIFY_EMAIL

router = APIRouter()

EMAIL_REGEX = re.compile(r"^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
MIN_USERNAME_LENGTH = 9
MIN_PASSWORD_LENGTH = 12

@router.post("/signup")
def signUp(request: SignUpRequest, user_repository: UserRepository = Depends(get_user_repository)):
    if not EMAIL_REGEX.match(request.email):
        raise BadRequest("INVALID_EMAIL", "Email is invalid")
    
    user = user_repository.get_user_by_email(request.email)
    if isinstance(user, User):
        raise BadRequest("INVALID_EMAIL", "Email already used")

    if len(request.username) < MIN_USERNAME_LENGTH:
        raise BadRequest("INVALID_USERNAME", f"Username must be atleast {MIN_USERNAME_LENGTH} characters long")
    
    user = user_repository.get_user_by_username(request.username)
    if isinstance(user, User):
        raise BadRequest("INVALID_USERNAME", "Username already taken")
    
    if len(request.password) < MIN_PASSWORD_LENGTH:
        raise BadRequest("INVALID_PASSWORD", f"Password must be atleast {MIN_PASSWORD_LENGTH} characters long")
    
    user = user_repository.create_user(request.username, request.email, request.password)

    token = create_access_token({
        "sub": user.user_id,
        "purpose": JWT_VERIFY_EMAIL
    })

    response = JSONResponse(content={
        "status": "ok",
        "token": token,
        "token_type": "bearer"
    })
    response.headers["Authorization"] = f"Bearer {token}"
    
    return response