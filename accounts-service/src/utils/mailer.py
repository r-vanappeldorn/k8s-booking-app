import os
import smtplib
from email.message import EmailMessage
from urllib.parse import quote

from jinja2 import Environment, FileSystemLoader, select_autoescape

from src.utils.auth import decode_token

SMTP_HOST = os.getenv("SMTP_HOST")
SMTP_PORT = int(os.getenv("SMTP_PORT", ""))
SMTP_USER = os.getenv("SMTP_USER", "")
SMTP_PASSWORD = os.getenv("SMTP_PASSWORD", "")
FROM_EMAIL = os.getenv("FROM_EMAIL")
BASE_URL = os.getenv("BASE_URL")

template_provider = Environment(
    loader=FileSystemLoader("src/templates/mail/"),
    autoescape=select_autoescape(["html", "xml"]),
    enable_async=False,
)


def send_verification_email(to: str, username: str, token: str):
    encoded_token = quote(token, safe="")
    link = f"{BASE_URL}/api/accounts/auth/verify-email?token={encoded_token}"

    msg = EmailMessage()
    msg["Subject"] = "Verify your email"
    msg["From"] = FROM_EMAIL
    msg["To"] = to

    payload = decode_token(token)
    minutes = payload.get("minutes")

    template = template_provider.get_template("verify_email.html")
    content = template.render(link=link, minutes=minutes, username=username)

    msg.set_content(f"Hi {username}, verify your e-mail through: {link}")
    msg.add_alternative(content, subtype="html")

    if not SMTP_HOST or not SMTP_PORT:
        return

    with smtplib.SMTP(SMTP_HOST, SMTP_PORT) as s:
        if SMTP_USER and SMTP_PASSWORD:
            s.starttls()
            s.login(SMTP_USER, SMTP_PASSWORD)
        s.send_message(msg)
