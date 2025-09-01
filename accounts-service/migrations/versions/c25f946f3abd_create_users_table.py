"""create users table

Revision ID: c25f946f3abd
Revises: 
Create Date: 2025-09-01 20:35:24.781036

"""
from typing import Sequence, Union

from alembic import op
import sqlalchemy as sa
from sqlalchemy.sql import func


# revision identifiers, used by Alembic.
revision: str = 'c25f946f3abd'
down_revision: Union[str, Sequence[str], None] = None
branch_labels: Union[str, Sequence[str], None] = None
depends_on: Union[str, Sequence[str], None] = None


def upgrade() -> None:
    """Upgrade schema."""
    op.create_table(
        "user",
        sa.Column("user_id", sa.Integer, primary_key=True),
        sa.Column("email", sa.Text, unique=True, nullable=False),
        sa.Column("username", sa.Text, unique=True, nullable=False),
        sa.Column("password_hash", sa.Text, unique=True, nullable=False),
        sa.Column('created_at', sa.DateTime(), nullable=False, server_default=func.now())
    )
    pass


def downgrade() -> None:
    """Downgrade schema."""
    op.drop_table("user")
    pass
