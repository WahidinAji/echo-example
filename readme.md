# Simple Echo With Mysql

- notes
```text
->  You can add password key-value in .env file and don't forget to import that
    into database/database.go
->  Don't forget to copy .env.example to .env
```
- example
```dotenv
DB_PASS=yourpassword
```
```go
dbPass := os.Getenv("DB_PASS")
if dbPass == "" {
	log.Fatal("DB_PASS ENVIRONMENT CANNOT BE NULL")
}
// change dsn line,
dsn := fmt.Sprintf("%s:%s@/%s?parseTime=true",user,dbPass,dbName)
```

## Posts
`post.go`
```text
define model from tables -> struct
```
`repository.go`
```text
define query sql of database -> injection by dependencies
```
`controller.go`
```text
define request and response from user, then inject by dependencies
```
```text
they'll looks like 
struct -> posts -> dependency *sql.DB 
                -> repository
                -> controller <- dependencies.repositories
                -> route/main <- dependencies.controllers

but, actually i'm just a little confused in this case. tell me when i'm incorrect
```

* Url
```http request
http://localhost:1323/posts
```
* Get Method -> Get All
```json
[
  {
    "id": 1,
    "title": "Post One",
    "desc": "desc"
  },
  {
    "id": 2,
    "title": "Post Two",
    "desc": "desc"
  },
  {
    "id": 3,
    "title": "post Three",
    "desc": "desc"
  }
]
```
* Url
```http request
http://localhost:1323/posts/{postId}
```
* Get Method -> Get Id,
```json
[
  {
    "id": 2,
    "title": "Post Two",
    "desc": "desc"
  }
]
```

## Books
```text
in this case, i'm just implement repository and then inject that
method to controller and then continues to route.
it's looks like

struct book -> repository -> controller -> route
```
* Url
```http request
http://localhost:1323/books
```
* Get Method -> Get All
```json
[
  {
    "id": 1,
    "author": "Aji",
    "title": "This is aji wanted"
  },
  {
    "id": 2,
    "author": "Wahidin",
    "title": "This is wahidin wanted"
  }
]
```

## Customers -> Get All
`handle/customer.go`
```text
actually i dunno what i've did in this case lol :D
```
* Url
```http request
http://localhost:1323/customers
```
* Get Method
```json
[
  {
    "id": 2,
    "name": "Aji",
    "email": "aji@mail.com"
  },
  {
    "id": 1,
    "name": "Wahidin",
    "email": "wahidin@mail.com"
  }
]
```
* Post Method 

`body request`
```json
{
    "name":"Wahidin",
    "email":"wahidin@mail.com"
}
```

`body response`
```json
{
  "id": 1,
  "name": "Wahidin",
  "email": "wahidin@mail.com"
}
```