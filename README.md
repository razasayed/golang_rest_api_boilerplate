## About

This project is a boilerplate for building REST api's in Go. It is a REST api for movies and uses a sqlite database.

## Web Framework

This project uses the Gin web framework (https://gin-gonic.com/) for Go which is one of the highly performant frameworks for building REST api's in Go.

## ORM

GORM (https://gorm.io)

## Endpoints

The following endpoints are available:

* GET /movies

* GET /movies/:id

* POST /movies

* PATCH /movies/:id

* DELETE /movies/:id

## Postman collection

There is a postman collection file available in the repo which can be imported in Postman to try out the api's.

## Future Enhancements

* Adding tests

* Implementing JWT authentication using jwt-go (https://github.com/dgrijalva/jwt-go) or any other package.

* PR's are welcome 🙏🏻
