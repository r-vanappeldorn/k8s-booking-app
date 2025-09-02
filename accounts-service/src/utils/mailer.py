import smtplib
from email.message import EmailMessage
from src.utils.auth import decode_token
import os

SMTP_HOST = os.getenv("SMTP_HOST")
SMTP_PORT = int(os.getenv("SMTP_PORT"))
SMTP_USER = os.getenv("SMTP_USER", "")
SMTP_PASSWORD = os.getenv("SMTP_PASSWORD", "")
FROM_EMAIL = os.getenv("FROM_EMAIL")
BASE_URL = os.getenv("BASE_URL")

def send_verification_email(to: str, token: str):
    link = f"{BASE_URL}/api/accounts/verify-email?token={token}"
    msg = EmailMessage()
    msg["Subject"] = "Verify your email"
    msg["From"] = FROM_EMAIL
    msg["To"] = to
    payload = decode_token(token)
    minutes = payload.get("minutes")

    msg.set_content(f"Click to verify: {link}\nthis link expires in {minutes} minutes.")

    with smtplib.SMTP(SMTP_HOST, SMTP_PORT) as s:
        if SMTP_USER and SMTP_PASSWORD:
            s.starttls()
            s.login(SMTP_USER, SMTP_PASSWORD)
        s.send_message(msg)