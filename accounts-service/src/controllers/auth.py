from fastapi import APIRouter, status, HTTPException
from dto.auth.SignUpRequest import SignUpRequest
from dto.common.BadRequest import BadRequest
import re 

router = APIRouter()

EMAIL_REGEX = re.compile(r"^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

@router.posr("/signup")
def signUp(request: SignUpRequest):

    if not EMAIL_REGEX.match(request.email):
        raise BadRequest("Email is invalid")
    
    # check email unique

    # check username unique

    # check 

    return {"status": "ok"}