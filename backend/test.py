import requests
import json
from pymongo import MongoClient
from bson.objectid import ObjectId

BASE_URL = "http://localhost:8080/api"

# MongoDB credentials
MONGO_USERNAME = "bugsfounder"
MONGO_PASSWORD = "kubari"
MONGO_HOST = "localhost"
MONGO_PORT = 27017
MONGO_DB = "bugsFounderDB"


def connect_to_mongo():
    uri = f"mongodb://{MONGO_USERNAME}:{MONGO_PASSWORD}@{MONGO_HOST}:{MONGO_PORT}/{MONGO_DB}"
    client = MongoClient(uri)
    return client[MONGO_DB]


def test_create_blog():
    url = f"{BASE_URL}/blog"
    payload = {
        "title": "Sample Blog Title",
        "content": "This is the content of the blog.",
        "author": "Author Name",
        "tags": ["tag1", "tag2"],
        "is_hidden": False,
    }
    response = requests.post(url, json=payload)
    print("Create Blog:", response.status_code, response.json())


def test_get_all_blogs():
    url = f"{BASE_URL}/blogs"
    response = requests.get(url)
    print("Get All Blogs:", response.status_code, response.json())


def test_get_blog_by_id(blog_id):
    url = f"{BASE_URL}/blogs/{blog_id}"
    response = requests.get(url)
    print("Get Blog By ID:", response.status_code, response.json())


def test_update_blog(blog_id):
    url = f"{BASE_URL}/update_blog/{blog_id}"
    payload = {
        "title": "Updated Blog Title",
        "content": "Updated content of the blog.",
        "author": "Updated Author Name",
        "tags": ["updated_tag1", "updated_tag2"],
        "is_hidden": False,
    }
    response = requests.post(url, json=payload)
    print("Update Blog:", response.status_code, response.json())


def test_delete_blog(blog_id):
    url = f"{BASE_URL}/delete_blog/{blog_id}"
    response = requests.post(url)
    print("Delete Blog:", response.status_code, response.json())


def test_hide_blog(blog_id):
    url = f"{BASE_URL}/hide_blog/{blog_id}"
    response = requests.post(url)
    print("Hide Blog:", response.status_code, response.json())


def test_unhide_blog(blog_id):
    url = f"{BASE_URL}/unhide_blog/{blog_id}"
    response = requests.post(url)
    print("Unhide Blog:", response.status_code, response.json())


def test_create_tutorial():
    url = f"{BASE_URL}/tutorial"
    payload = {
        "title": "Sample Tutorial Title",
        "description": "This is the description of the tutorial.",
        "author": "Author Name",
        "tags": ["tag1", "tag2"],
        "is_hidden": False,
        "sub_tutorial": [
            {
                "title": "Sub Tutorial Title",
                "content": "Content of the sub tutorial.",
                "author": "Sub Author Name",
                "tags": ["sub_tag1", "sub_tag2"],
                "is_hidden": False,
            }
        ],
    }
    response = requests.post(url, json=payload)
    print("Create Tutorial:", response.status_code, response.json())


def test_get_all_tutorials():
    url = f"{BASE_URL}/tutorials"
    response = requests.get(url)
    print("Get All Tutorials:", response.status_code, response.json())


def test_get_tutorial_by_id(tutorial_id):
    url = f"{BASE_URL}/tutorials/{tutorial_id}"
    response = requests.get(url)
    print("Get Tutorial By ID:", response.status_code, response.json())


def test_update_tutorial(tutorial_id):
    url = f"{BASE_URL}/update_tutorial/{tutorial_id}"
    payload = {
        "title": "Updated Tutorial Title",
        "description": "Updated description of the tutorial.",
        "author": "Updated Author Name",
        "tags": ["updated_tag1", "updated_tag2"],
        "is_hidden": False,
        "sub_tutorial": [
            {
                "title": "Updated Sub Tutorial Title",
                "content": "Updated content of the sub tutorial.",
                "author": "Updated Sub Author Name",
                "tags": ["updated_sub_tag1", "updated_sub_tag2"],
                "is_hidden": False,
            }
        ],
    }
    response = requests.post(url, json=payload)
    print("Update Tutorial:", response.status_code, response.json())


def test_delete_tutorial(tutorial_id):
    url = f"{BASE_URL}/delete_tutorial/{tutorial_id}"
    response = requests.post(url)
    print("Delete Tutorial:", response.status_code, response.json())


def test_hide_tutorial(tutorial_id):
    url = f"{BASE_URL}/hide_tutorial/{tutorial_id}"
    response = requests.post(url)
    print("Hide Tutorial:", response.status_code, response.json())


def test_unhide_tutorial(tutorial_id):
    url = f"{BASE_URL}/unhide_tutorial/{tutorial_id}"
    response = requests.post(url)
    print("Unhide Tutorial:", response.status_code, response.json())


def test_signup():
    url = f"{BASE_URL}/signup"
    payload = {
        "username": "new_user",
        "email": "new_user@example.com",
        "password": "securepassword",
    }
    response = requests.post(url, json=payload)
    print("Signup:", response.status_code, response.json())


def test_login():
    url = f"{BASE_URL}/login"
    payload = {"username": "existing_user", "password": "securepassword"}
    response = requests.post(url, json=payload)
    print("Login:", response.status_code, response.json())


def test_logout():
    url = f"{BASE_URL}/logout"
    response = requests.get(url)
    print("Logout:", response.status_code, response.json())


def test_create_user():
    url = f"{BASE_URL}/create_user"
    payload = {
        "username": "new_user",
        "email": "new_user@example.com",
        "password": "securepassword",
        "is_active": True,
        "is_admin": False,
    }
    response = requests.post(url, json=payload)
    print("Create User:", response.status_code, response.json())


def test_get_user_by_username(username):
    url = f"{BASE_URL}/user/{username}"
    response = requests.get(url)
    print("Get User By Username:", response.status_code, response.json())


def test_get_user_by_email(email):
    url = f"{BASE_URL}/user_email/{email}"
    response = requests.get(url)
    print("Get User By Email:", response.status_code, response.json())


def test_update_user(user_id):
    url = f"{BASE_URL}/update_user/{user_id}"
    payload = {
        "username": "updated_user",
        "email": "updated_user@example.com",
        "password": "newsecurepassword",
        "is_active": True,
        "is_admin": False,
    }
    response = requests.post(url, json=payload)
    print("Update User:", response.status_code, response.json())


def test_delete_user(user_id):
    url = f"{BASE_URL}/delete_user/{user_id}"
    response = requests.post(url)
    print("Delete User:", response.status_code, response.json())


if __name__ == "__main__":
    # Connect to the MongoDB server
    collection = connect_to_mongo()

    # Query to find the document
    document = collection.find_one({"blogid": 1})

    if document:
        print("Document ID:", document["_id"])
    else:
        print("Document not found")

    # Blog tests
    test_create_blog()
    test_get_all_blogs()
    test_get_blog_by_id("sample_blog_id")
    test_update_blog("sample_blog_id")
    test_delete_blog("sample_blog_id")
    test_hide_blog("sample_blog_id")
    test_unhide_blog("sample_blog_id")

    # Tutorial tests
    test_create_tutorial()
    test_get_all_tutorials()
    test_get_tutorial_by_id("sample_tutorial_id")
    test_update_tutorial("sample_tutorial_id")
    test_delete_tutorial("sample_tutorial_id")
    test_hide_tutorial("sample_tutorial_id")
    test_unhide_tutorial("sample_tutorial_id")

    # User tests
    test_signup()
    test_login()
    test_logout()
    test_create_user()
    test_get_user_by_username("new_user")
    test_get_user_by_email("new_user@example.com")
    test_update_user("sample_user_id")
    test_delete_user("sample_user_id")
