# Mongo-Go REST API

mongo-api is a REST API using MongoDB, Redis, Docker, Kubernetes and JWT.

### Features!

- Uses Gin for routing
- Uses a salt to hash the password
- Uses JWT for authentication
- Uses struct separation to hide sensitive information from call responses
- Error handling
- Viper for secret management

### Installation

`mongo-api` requires [Go](https://golang.org/) 1.9+ to run.

Install the dependencies and devDependencies and start the server.

```sh
$ git clone https://github.com/morelmiles/mongo-api.git
$ cd mongo-api
$ go run main.go
```

### Building for source

For production release:

```sh
$ go build
```

### Todos

- Write Tests

### Tech

Project uses a number of open source projects to work properly:

- [Gin] - Implements a request router and dispatcher in Go
- [MongoDB] - document-based, big community, database
- [Redis] - in-memory database using key-value pairs
- [Docker] - Build, Share, and Run Any App, Anywhere
- [Kubernetes] - Automating deployment, scaling, and management of containerized applications

API endpoint - http://localhost:8080
Swagger endpoint - http://localhost:8080/swagger/index.html

[gin]: https://gin-gonic.com/
[mongodb]: https://www.mongodb.com/
[docker]: https://www.docker.com/
[kubernetes]: https://kubernetes.io/
[redis]: https://redis.io/
