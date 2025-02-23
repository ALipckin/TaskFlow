installation:

```
docker network create task-network

cd ./TaskRestApiService

docker compose up -d --build

cd ../TaskStorageService/

docker compose up -d --build

```
