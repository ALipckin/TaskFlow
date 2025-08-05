#!/bin/bash

docker network inspect task-network >/dev/null 2>&1 || docker network create task-network

git config core.fileMode false
(  cd ./backend-api/ || exit
   ./install.sh
)

echo "Waiting for Kafka to become healthy..."
until [ "$(docker inspect --format='{{.State.Health.Status}}' kafka 2>/dev/null)" == "healthy" ]; do
  sleep 1
done
echo "Kafka is started"

(cd ./Backend/AuthService && docker compose up -d --build) &
(cd ./Backend/TaskStorageService && docker compose up -d --build) &
(cd ./Backend/NotifyService && docker compose up -d --build) &
(cd ./Frontend && docker compose up -d --build) &

wait