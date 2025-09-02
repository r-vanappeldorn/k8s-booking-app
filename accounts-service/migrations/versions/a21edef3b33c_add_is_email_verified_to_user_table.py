"""add is_email_verified to user table

Revision ID: a21edef3b33c
Revises: c25f946f3abd
Create Date: 2025-09-02 14:46:48.541590

"""
from typing import Sequence, Union

from alembic import op
import sqlalchemy as sa


# revision identifiers, used by Alembic.
revision: str = 'a21edef3b33c'
down_revision: Union[str, Sequence[str], None] = 'c25f946f3abd'
branch_labels: Union[str, Sequence[str], None] = None
depends_on: Union[str, Sequence[str], None] = None


def upgrade() -> None:
    """Upgrade schema."""
    op.add_column(
        "user",
        sa.Column("is_email_verified", sa.Boolean, nullable=False, server_default=sa.text("0"))
    )
    pass


def downgrade() -> None:
    """Downgrade schema."""
    op.drop_column("user", "is_email_verified")
    pass
