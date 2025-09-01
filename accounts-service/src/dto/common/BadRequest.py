from typing import Any, Optional, Dict
from fastapi import status, HTTPException

class BadRequest(HTTPException):
    def __init__(self, detail: str = "Something went wrong", headers: Optional[Dict[str, str]] = None):

        super().__init__(
            status_code=status.HTTP_400_BAD_REQUEST,
            detail=detail,
            headers=headers
            )