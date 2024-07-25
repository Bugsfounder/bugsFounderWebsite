from fastapi import FastAPI, HTTPException, Depends
from pydantic import BaseModel
from typing import List
from api_data_models import (
    BlogPost,
    Tutorial,
    User,
    AuthenticationData,
    AuthResponse,
    Message,
)

app = FastAPI()

# Mock databases
blog_posts = []
tutorials = []
users = []


# Mock authentication
def authenticate_user(auth_data: AuthenticationData):
    for user in users:
        if (
            user.username == auth_data.username
            and user.hashed_password == auth_data.password
        ):
            return {"access_token": "mocked_access_token", "token_type": "bearer"}
    raise HTTPException(status_code=401, detail="Invalid credentials")


@app.post("/auth/login", response_model=AuthResponse)
def login(auth_data: AuthenticationData):
    return authenticate_user(auth_data)


@app.post("/users/", response_model=User)
def create_user(user: User):
    users.append(user)
    return user


@app.get("/users/", response_model=List[User])
def get_users():
    return users


@app.post("/blogposts/", response_model=BlogPost)
def create_blog_post(blog_post: BlogPost):
    blog_posts.append(blog_post)
    return blog_post


@app.get("/blogposts/", response_model=List[BlogPost])
def get_blog_posts():
    return blog_posts


@app.post("/tutorials/", response_model=Tutorial)
def create_tutorial(tutorial: Tutorial):
    tutorials.append(tutorial)
    return tutorial


@app.get("/tutorials/", response_model=List[Tutorial])
def get_tutorials():
    return tutorials


@app.get("/")
def read_root():
    return {"message": "Welcome to the API"}


# Example usage of Message model
@app.get("/status/", response_model=Message)
def get_status():
    return {"message": "Operation successful", "code": 200}


# Run the application with: uvicorn api_gen:app --reload
