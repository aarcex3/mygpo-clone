"""Utilities for the auth routes"""

import bcrypt


def hash_password(password: str) -> str:
    """Hash the user password when registering"""
    password_bytes = password.encode("utf-8")
    salt = bcrypt.gensalt()
    hashed_password = bcrypt.hashpw(password_bytes, salt)
    return hashed_password.decode("utf-8")


def check_password(input_password: str, hashed_password: str) -> bool:
    """
    Check if the given password matches the one in the db
    """
    try:
        input_password_bytes = input_password.encode("utf-8")
        hashed_password_bytes = hashed_password.encode("utf-8")
        return bcrypt.checkpw(input_password_bytes, hashed_password_bytes)
    except (ValueError, TypeError) as e:
        print(f"Error checking password: {e}")
        return False
