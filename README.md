## REST API using Gin framework and GORM

Here is my pet project to learn Golang concepts.
Stack:
* Golang 1.21
* Postgres
* Docker
* Gin framework
* GORM

To start using Docker:
```
docker compose --file ./.docker/docker-compose.dev.yml up -d
```

Ping endpoint should work
```
GET http://localhost/public/ping
```

