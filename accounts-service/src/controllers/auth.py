import re 
import os
from fastapi import APIRouter, Depends, Request
from fastapi.responses import JSONResponse, RedirectResponse
from urllib.parse import unquote

from src.dto.auth.sign_up_request import SignUpRequest
from src.dto.auth.sign_in_request import SignInRequest
from src.dto.common.bad_request import BadRequest
from src.repositories.user_repository import UserRepository, get_user_repository
from src.models.user import User
from src.utils.auth import create_access_token, decode_token, JWT_VERIFY_EMAIL, JWT_SIGNED_IN
from src.utils.mailer import send_verification_email
from src.middleware.auth_middleware import get_current_user

EMAIL_REGEX = re.compile(r"^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
MIN_USERNAME_LENGTH = 9
MIN_PASSWORD_LENGTH = 12

router = APIRouter(prefix="/auth")

@router.get('/info')
def info(user: User = Depends(get_current_user)):
    return {
        "user_id": user.user_id,
        "username": user.username,
        "email": user.email
    }

@router.post("/sign-in")
def sign_in(sign_in_request: SignInRequest, user_repository: UserRepository = Depends(get_user_repository)):
    user = None
    if sign_in_request.username != None:
        user = user_repository.get_user_by_username(sign_in_request.username)

    if sign_in_request.email != None:
        user = user_repository.get_user_by_email(sign_in_request.email)
    
    if not user:
        raise BadRequest("INVALID_CREDENTIALS", "Email or username is incorrect")
    
    if not user.is_email_verified:
        raise BadRequest("EMAIL_NOT_VERIFIED", "Email is not verified")
    
    if not user_repository.is_password_valid(sign_in_request.password, user.password_hash):
        raise BadRequest("INVALID_PASSWORD", "Invalid password")
    
    token = create_access_token({
        "sub": str(user.user_id),
        "purpose": JWT_SIGNED_IN
    }, 1440)

    response = JSONResponse({
        "status": "ok"
    })
    response.headers.append("Authorization", f"Bearer {token}")

    return response

@router.get('/verify-email')
def verify_email(token: str, user_repository: UserRepository = Depends(get_user_repository)):
    unquoted_token = unquote(token)

    try:
        payload = decode_token(unquoted_token)
    except Exception:
        raise BadRequest('INVALID_TOKEN', "Verify link is invalid or expired")
    
    if payload.get("purpose") != JWT_VERIFY_EMAIL:
        raise BadRequest("INVALID_PURPOSE", "Verify link is invalid or expired")
    
    user_id = int(payload.get("sub"))
    user_repository.verify_email_by_user_id(user_id)

    url = f"{os.environ["BASE_URL"]}/singin"
    response = RedirectResponse(url)
    response.set_cookie("email_verified", "Verification succeded")

    return response

@router.post("/sign-up")
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
        "sub": str(user.user_id),
        "purpose": JWT_VERIFY_EMAIL
    })

    send_verification_email(user.email, user.username, token)

    response = JSONResponse(content={
        "status": "verification_mail_send",
    })
    
    return response