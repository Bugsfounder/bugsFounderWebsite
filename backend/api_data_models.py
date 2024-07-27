from pydantic import BaseModel
from typing import List, Optional

"""
blog
tutorial
user
authentication
authResponse
message
"""


class BlogPost(BaseModel):
    title: str
    content: str
    author_id: str
    tags: List[str] = []

    class Config:
        schema_extra = {
            "example": {
                "title": "Introduction to GO",
                "content": "Go, also known as Golang, is a statically typed, compiled programming language designed at Google...",
                "author_id": "user123",
                "tags": ["programming", "go", "introduction"],
            }
        }


class Tutorial(BaseModel):
    title: str
    content: str
    author_id: str
    tags: List[str] = []

    class Config:
        schema_extra = {
            "example": {
                "title": "React Basics",
                "content": "React is a JavaScript library for building user interfaces...",
                "author_id": "user123",
                "tags": ["web development", "react", "javascript"],
            }
        }


class User(BaseModel):
    username: str
    email: str
    hashed_password: str
    is_active: bool = True
    is_admin: bool = False

    class Config:
        schema_extra = {
            "example": {
                "username": "john_doe",
                "email": "john@example.com",
                "hashed_password": "hashedpassword123",
                "is_active": True,
                "is_admin": False,
            }
        }


class AuthenticationData(BaseModel):
    username: str
    password: str

    class Config:
        schema_extra = {
            "example": {"username": "john_doe", "password": "securepassword"}
        }


class AuthResponse(BaseModel):
    access_token: str
    token_type: str

    class Config:
        schema_extra = {
            "example": {"access_token": "your_access_token", "token_type": "bearer"}
        }


class Message(BaseModel):
    message: str
    code: int

    class Config:
        schema_extra = {"example": {"message": "Operation successful", "code": 200}}
