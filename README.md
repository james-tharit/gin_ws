# gin_ws
Starter of Gin webservice integrated with Mongodb and DockerCompose

## Start Application with Docker compose
```shell
docker compose up
```

## Start Application in Background
```shell
docker compose up -d
```

## Endpoint

1. GET /albums
```shell
curl http://localhost:8080/albums --header "Content-Type: application/json" --request "GET"
```
