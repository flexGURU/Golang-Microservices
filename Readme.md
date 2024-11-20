# Proudct API


### RESTful Go based JSON API built using the Gorilla framework. The API allows CRUD based operations on a product list.

## Content
### MicroServices
- use of Http Handlers for dealing with http reponses and requests
- Gorilla Mux for creating serve Mux for dealing with routing and Middleware
- Swagger Open API for API documentation -- ## [./swagger.yaml](./swagger.yaml)
- Running tests for the Product struct Validation [./data/products_test.go](./data/products_test.go)


## Golang Basics
- Go Channels for concurrency 
- use of Contexts to deal with Server Shutdown [./working/main.go](./working/main.go)
- use of interfaces for the Product Struct Type
- creating struct type functions for the Product type for Add and List product


## Configuartion of Make files:
[./makefile](./makefile)