#!/bin/bash

docker network inspect task-network >/dev/null 2>&1 || docker network create task-network

(cd ./Backend/TaskRestApiService && docker compose up -d --build)

echo "Waiting start task-rest-api-service..."
until curl -s http://localhost:5437/health >/dev/null; do
  sleep 1
done
echo "task-rest-api-service запущен!"

(cd ./Backend/AuthService && docker compose up -d --build) &
(cd ./Backend/TaskStorageService && docker compose up -d --build) &
(cd ./Backend/NotifyService && docker compose up -d --build) &
(cd ./Frontend && docker compose up -d --build) &

wait