# Vote-Go API

Vote App is an application that allows users to vote online on draft laws.

## Requirements

- GO 1.13 
(https://golang.org/doc/install)

## Installation

- Clone project anywhere on your computer.

```
git clone https://github.com/muhammad-tounsi/vote-go-vue.git
```
- Install all the dependencies

```
make install
```

- Run this command to build and run the project

```
make run
```

- Run this command to start a builded version of the project :

```
make start
```

or

- Run with local environnement (Local)

```
go mod init
go get ./


docker-compose up -d

go run main.go
``

Postman request


Home page 

GET : http://localhost:8080/

return

"Welcome To This Awesome API"

Create User

POST : http://localhost:8080/users

body : 
    {
        "email": "sky@gmail.com",
        "firstname": "vky",
        "lastname": "Val",
        "access_level": 1,
        "date_of_birth": "1992-02-08T00:00:00Z",
        "password": "password"
    }

return value:

    {
        "uuid": "aef29d3c-de46-4e0f-bad4-138c1975ff8c",
        "id": 4,
        "email": "sky@gmail.com",
        "firstname": "vky",
        "lastname": "Val",
        "access_level": 1,
        "date_of_birth": "1992-02-08T00:00:00Z",
        "password": "$2a$10$c85rda4Pl.N/Z/jQFPFmTO0olcagk4aBoxsrvE9h5KpSyhMy52CuK",
        "created_at": "2019-10-26T13:55:37.7289705+02:00",
        "updated_at": "2019-10-26T13:55:37.6704403+02:00"
    }




Login 

Post : http://localhost:8080/login

{
	"email": "steven@gmail.com",
	"password" : "password"
}

return 

// token de connexion qu'il faudra utiliser à la suite
token : "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE1NzIwOTI5MjUsInVzZXJfaWQiOjF9.FFdb2QPXDudb9lycxocFxcu6RutHVyX60M46AJ0WEYA"



Get all User

GET : http://localhost:8080/users



get user with id
GET : http://localhost:8080/users/{id}
GET : http://localhost:8080/users/1

Update user



il te faut le token lié à l'utilisateur que tu veux update 

Put : http://localhost:8080/users/2

Params

token : ...

Body : 
{
	 "firstname": "Ssssteven",
	 "email":"sky2@gmail.com",
     "lastname": "Victor2222",
     "password" : "Password"
}

delete 

Params

token : ...

Detelete :  http://localhost:8080/users/2


Create vote 

POST: http://localhost:8080/votes
 
params

token du user


Body 

{
	"Title":     "Propreté des trottoirs",
	"AuthorID" : 3,
	"Desc":      "nouveau vote ."
}


get all votes 

get : http://localhost:8080/votes

get votes by id

get : http://localhost:8080/votes/1

Update vote

put : http://localhost:8080/votes/Update/1


params

token du user

body 

{
	"Title":   ,
	"Desc":     
}


Add vote

put : http://localhost:8080/votes/Add/1


params

token du user


delete vote

delete : http://localhost:8080/votes/id







## Developers

- [TOUNSI Muhammad](https://github.com/Muhammad-Tounsi)
- [MOUHSIN Walid](https://github.com/walidmh)
- [ZAIR Aymane](https://github.com/saneto)
- [CHEA Hugo](https://github.com/HugoChea)
