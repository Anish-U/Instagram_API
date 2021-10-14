<p align="center">
  <a href="" rel="noopener">
 <img height=200px src="https://raw.githubusercontent.com/ahmdrz/goinsta/v1/resources/goinsta-image.png" alt="Project logo"></a>
</p>

<h1 align="center">Instagram API</h1>

<p align="center"> Simple Instagram Backend REST API using GO and MongoDB</p>

---

### ❗❗ Requirements

To be able show desired features of REST API such as:

* [x] `POST /users/` creates a new user using JSON request body
* [x] `GET /users/{id}` returns details of specific user as JSON
* [x] `POST /posts/` creates a new post using JSON request body
* [x] `GET /posts/{id}` returns specific post as JSON
* [x] `GET /posts/users/{id}` returns list of all posts of specific user as JSON

`The API should be developed only using Go & MongoDB.` <br> 
`No frameworks / third-party libraries to be used.`
<br> 
`Pagination and Unit Tests can be implemented`

---

### ⚒ Data Types

A User object should look like this:
```
{
  "id": "ID",
  "name": "name of the user",
  "email": "email of the user",
  "password": "password of the user",
}
```

A Post object should look like this:
```
{
  "id": "ID",
  "caption": "caption to the post",
  "imageURL": "link to the image",
  "timestamp": "password of the user",
  "userID": "user ID of owner to the post",
}
```

---

## 📸 Screenshots

`POST /users/`

<img src="https://raw.githubusercontent.com/Anish-U/Instagram_API/master/images/createUser_POST.jpg" alt="Project logo">

---

`GET /users/{id}`

<img src="https://raw.githubusercontent.com/Anish-U/Instagram_API/master/images/getUser_GET.jpg" alt="Project logo">

---

`POST /posts/`

<img src="https://raw.githubusercontent.com/Anish-U/Instagram_API/master/images/createPost_POST.jpg" alt="Project logo">

---

`GET /posts/{id}`

<img src="https://raw.githubusercontent.com/Anish-U/Instagram_API/master/images/getPost_GET.jpg" alt="Project logo">

---

`GET /posts/users/{id}`

<img src="https://raw.githubusercontent.com/Anish-U/Instagram_API/master/images/getUserPosts_GET.jpg" alt="Project logo">

---

`Unit Tests => server_test.go`

<img src="https://raw.githubusercontent.com/Anish-U/Instagram_API/master/images/unitTests.jpg" alt="Project logo">

---

## ✍️ Authors

- [Anish-U](https://github.com/anish-u)