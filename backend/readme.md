example:
blog_manager --> contains methods related to blogs
api_handler -> apis

app_manager -> endpoint methods used in api's handler methods

blog_Manager->app_manager->api_handper

# Backend - Go

### API's

#### PUBLIC

| API                                                  | METHOD | BODY | DESCRIPTION                |
| ---------------------------------------------------- | ------ | ---- | -------------------------- |
| <b> BLOG </b>                                        |        |      |                            |
| /api/public/blog                                     | POST   | JSON | Add New Blog               |
| /api/public/blog                                     | GET    | -    | Get all Blogs              |
| /api/public/blog/:blog_url                           | PUT    | JSON | Update one blog by u rl    |
| /api/public/blog/:blog_url                           | GET    | -    | Get one blog by url        |
| /api/public/blog/:blog_url                           | DELETE | -    | Delete one blog by url     |
| <b> TUTORIAL </b>                                    |        |      |                            |
| /api/public/tutorial                                 | POST   | JSON | Add New Tutorial           |
| /api/public/tutorial                                 | GET    | JSON | Get all Tutorials          |
| /api/public/tutorial/:tutorial_url                   | POST   | JSON | Add New SubTutorial        |
| /api/public/tutorial/:tutorial_url                   | GET    | -    | Get one Tutorial           |
| /api/public/tutorial/:tutorial_url/:sub_tutorial_url | GET    |      | Get Sub tutorial by url    |
| /api/public/tutorial/:tutorial_url/:sub_tutorial_url | PUT    | JSON | Update sub tutorial by url |
| /api/public/tutorial/:tutorial_url                   | PUT    | JSON | update tutorial            |
| /api/public/tutorial/:tutorial_url/:sub_tutorial_url | DELETE |      | delete a sub tutorial      |
| /api/public/tutorial/:tutorial_url                   | DELETE |      | delete tutorial            |
| <b> USER </b>                                        |        |      |                            |
| /api/public/user/signup                              | POST   | JSON | Create New User            |
| /api/public/user/getAllUser                          | GET    |      |                            |
| /api/public/user/login                               | POST   | JSON |                            |
| /api/public/user/logout                              | POST   | JSON |                            |
| /api/public/user/update/:username_or_email           | PUT    | JSON |                            |
| /api/public/user/delete/:username_or_email           | DELETE |      |                            |
| <b>SEARCH</b>                                        |        |      |                            |
| /api/public/search/blog_title/:query                 | GET    |      |                            |
| /api/public/search/tutorial_title/:query             | GET    |      |                            |
| /api/public/search/:title/:sub_title                 | GET    |      |                            |
| /api/public/search/:query                            | GET    |      |                            |

### TESTING API's

1.  Create One Blog - POST

    url:

    ```
    http://localhost:8080/api/public/blog
    ```

    body:

    ```json
    {
      "title": "Understanding Concurrency in Go now",
      "content": "Concurrency in Go is achieved using goroutines...",
      "author": "Alice Johnson",
      "tags": ["go", "concurrency", "programming"],
      "url": "understanding-concurrency-in-go-3"
    }
    ```

    response:

    ```json
    {
      "inserted_id": "66ab4640941c253beb929a33"
    }
    ```

2.  Create One Tutorial - POST

    url:

    ```
    http://localhost:8080/api/public/tutorial
    ```

    body:

    ```json
    {
      "title": "Building RESTful APIs with Go",
      "description": "This tutorial covers the basics of building RESTful APIs using Go and Gin framework...",
      "author": "John Doe",
      "tags": ["go", "api", "rest", "gin"],
      "sub_tutorials": [
        {
          "title": "Introduction to RESTful APIs",
          "content": "REST stands for Representational State Transfer...",
          "author": "John Doe",
          "tags": ["api", "rest"],
          "url": "introduction-to-restful-apis"
        },
        {
          "title": "Setting up a Go Project",
          "content": "In this section, we will set up a new Go project...",
          "author": "John Doe",
          "tags": ["go", "project setup"],
          "url": "setting-up-a-go-project"
        }
      ],
      "url": "building-restful-apis-with-go"
    }
    ```

    response:

    ```json
    {
      "inserted_id": "66ab4643941c253beb929a34"
    }
    ```

3.  Create One Sub Tutorial - POST

    url:

    ```
    http://localhost:8080/api/public/tutorial/building-restful-apis-with-go
    ```

    body:

    ```json
    {
      "title": "Go Variables",
      "content": "A detailed explanation of variables in Go.",
      "author": "Jane Smith",
      "tags": ["go", "variables"],
      "url": "go-variables-updated",
      "created_at": "2023-01-01T00:00:00Z",
      "updated_at": "2023-01-01T00:00:00Z"
    }
    ```

    response:

    ```json
    {
      "MatchedCount": 1,
      "ModifiedCount": 1,
      "UpsertedCount": 0,
      "UpsertedID": null
    }
    ```

4.  Create One User - POST

    url:

    ```
    http://localhost:8080/api/public/user/signup
    ```

    body:

    ```json
    {
      "email": "jocefyneroot200307@gmail.com",
      "username": "jocefyneroot",
      "password": "66a6719903fbde059f74e01a"
    }
    ```

    response:

    ```json
    {
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImpvY2VmeW5lcm9vdCIsImVtYWlsIjoiam9jZWZ5bmVyb290MjAwMzA3QGdtYWlsLmNvbSIsImV4cCI6MTcyMjUwMDk3OH0.E6zciEmcFOgUWCl_pya1QIZn8iQv2Gpz_tuBjgLjHRo"
    }
    ```

5.  Update a blog - PUT

    url:

    ```
    http://localhost:8080/api/public/blog/understanding-concurrency-in-go
    ```

    body:

    ```json
    {
      "title": "",
      "content": "updated description is here now Concurrency in Go is achieved using goroutines...",
      "author": "",
      "tags": ["go", "concurrency", "programming", "naming"],
      "url": "understanding-concurrency-in-go"
    }
    ```

    response:

    ```json
    {
      "MatchedCount": 0,
      "ModifiedCount": 0,
      "UpsertedCount": 0,
      "UpsertedID": null
    }
    ```

6.  Update a Tutorial - PUT

    url:

    ```
    http://localhost:8080/api/public/tutorial/building-restful-apis-with-go
    ```

    body:

    ```json
    {
      "title": "building-restful-apis-with-go Go Programming",
      "description": "building-restful-apis-with-goAn advanced tutorial on Go programming language.",
      "author": "Jane Smith",
      "tags": ["programming", "go", "advanced"],
      "url": "building-restful-apis-with-go-updated-again"
    }
    ```

    response:

    ```json
    {
      "MatchedCount": 1,
      "ModifiedCount": 1,
      "UpsertedCount": 0,
      "UpsertedID": null
    }
    ```

7.  Update a sub tutorial - PUT

    url

    ```
    http://localhost:8080/api/public/tutorial/building-restful-apis-with-go/go-variables-constants
    ```

    body

    ```json
    {
      "title": "Go Variables and Constants",
      "content": "A detailed explanation of variables and constants in Go.",
      "author": "Jane Smith",
      "tags": ["go", "variables", "constants"],
      "url": "go-variables-constants-updated"
    }
    ```

    response

    ```json
    {
      "message": "Sub-Tutorial not found"
    }
    ```

8.  Update a user - PUT
    url

    ```
    http://localhost:8080/api/public/users/jocefyneroot
    ```

    body

    ```json
    {
      "email": "jocefyneroot200",
      "username": "",
      "password": ""
    }
    ```

    response

    ```json
    {
      "message": "User updated successfully"
    }
    ```

9.  Delete a blog - DELETE
    url

    ```
    http://localhost:8080/api/public/blog/understanding-concurrency-in-go-2
    ```

    body

    ```json

    ```

    response

    ```json
    {
      "message": "blog deleted successfully"
    }
    ```

10. Delete a tutorial - DELETE
    url

    ```
    http://localhost:8080/api/public/tutorial/advanced-go-programming
    ```

    body

    ```json

    ```

    response

    ```json
    {
      "message": "Tutorial deleted successfully"
    }
    ```

11. Delete a subTutorial - DELETE
    url

    ```
    http://localhost:8080/api/public/tutorial/building-restful-apis-with-go-updated-again/setting-up-a-go-project
    ```

    body

    ```json

    ```

    response

    ```json
    {
      "message": "Sub-Tutorial not found"
    }
    ```

12. Delete a user - DELETE
    url

    ```
    http://localhost:8080/api/public/user/delete/bugsfounder2021@gmail.com
    ```

    body

    ```json

    ```

    response

    ```json
    {
      "message": "User deleted successfully"
    }
    ```

13. Search a blogTitle - GET
    url

    ```
    http://localhost:8080/api/public/search/blog_title/Advanced%20Go%20Programming
    ```

    body

    ```json
    none
    ```

    response

    ```json
    {
      "message": 0
    }
    ```

14. Search a TutorialTItle - GET
    url

    ```
    http://localhost:8080/api/public/search/tutorial_title/building-restful-apis-with-go Go Programming
    ```

    body

    ```json
    none
    ```

    response

    ```json
    {
      "message": 2
    }
    ```

15. Search a subTutorialTitle - GET
    url

    ```
    http://localhost:8080/api/public/search/sub_tutorial_title/building-restful-apis-with-go%20Go%20Programming/Introduction%20to%20RESTful APIs
    ```

    body

    ```json
    none
    ```

    response

    ```json
    {
      "message": 2
    }
    ```

16. Search a Search - GET
    url

    ```
    http://localhost:8080/api/public/search/programming
    ```

    body

    ```json
    none
    ```

    response

    ```json
    {
      "result": [
        {
          "blogs": [
            {
              "_id": "66a67473537e7a09a3f98875",
              "author": "Jane Smith",
              "content": "An advanced tutorial on Go programming language.",
              "createdat": "0001-01-01T00:00:00Z",
              "tags": ["programming", "go", "advanced"],
              "title": "Advanced Go Programming",
              "updated_at": "2024-07-30T02:50:54.016Z",
              "updatedat": "0001-01-01T00:00:00Z",
              "url": "advanced-go-programming"
            }
          ]
        },
        {
          "tutorials": [
            {
              "_id": "66a64aeece719ac744cad37f",
              "author": "Jane Smith",
              "content": "building-restful-apis-with-goAn advanced tutorial on Go programming language.",
              "createdat": "2024-07-28T13:43:10.519Z",
              "description": "This tutorial covers the basics of building RESTful APIs using Go and Gin framework...",
              "sub_tutorials": [
                {
                  "author": "John Doe",
                  "content": "REST stands for Representational State Transfer...",
                  "createdat": "0001-01-01T00:00:00Z",
                  "tags": ["api", "rest"],
                  "title": "Introduction to RESTful APIs",
                  "updatedat": "0001-01-01T00:00:00Z",
                  "url": "introduction-to-restful-apis"
                },
                {
                  "author": "John Doe",
                  "content": "In this section, we will set up a new Go project...",
                  "createdat": "0001-01-01T00:00:00Z",
                  "tags": ["go", "project setup"],
                  "title": "Setting up a Go Project",
                  "updatedat": "0001-01-01T00:00:00Z",
                  "url": "setting-up-a-go-project"
                }
              ],
              "tags": ["programming", "go", "advanced"],
              "title": "building-restful-apis-with-go Go Programming",
              "updated_at": "2024-07-30T02:53:36.625Z",
              "updatedat": "2024-07-28T13:43:10.519Z",
              "url": "building-restful-apis-with-go-updated"
            },
            {
              "_id": "66a64af0ce719ac744cad380",
              "author": "Jane Smith",
              "content": "building-restful-apis-with-goAn advanced tutorial on Go programming language.",
              "createdat": "2024-07-28T13:43:12.075Z",
              "description": "This tutorial covers the basics of building RESTful APIs using Go and Gin framework...",
              "sub_tutorials": [
                {
                  "author": "John Doe",
                  "content": "REST stands for Representational State Transfer...",
                  "createdat": "0001-01-01T00:00:00Z",
                  "tags": ["api", "rest"],
                  "title": "Introduction to RESTful APIs",
                  "updatedat": "0001-01-01T00:00:00Z",
                  "url": "introduction-to-restful-apis"
                },
                {
                  "author": "Jane Smith",
                  "content": "A detailed explanation of variables in Go.",
                  "createdat": "2024-07-30T03:10:39.881Z",
                  "tags": ["go", "variables"],
                  "title": "Go Variables",
                  "updatedat": "2024-07-30T03:10:39.881Z",
                  "url": "go-variables-updated"
                }
              ],
              "tags": ["programming", "go", "advanced"],
              "title": "building-restful-apis-with-go Go Programming",
              "updated_at": "2024-07-30T03:11:42.501Z",
              "updatedat": "2024-07-28T13:43:12.075Z",
              "url": "building-restful-apis-with-go-updated-again"
            }
          ]
        },
        {
          "subTutorials": null
        }
      ]
    }
    ```

#### updated user enums

```
1 -> reset password
2 -> reset email
3 -> reset username
```
