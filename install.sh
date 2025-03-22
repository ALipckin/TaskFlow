docker network create task-network

(cd ./Backend/TaskRestApiService && docker compose up -d --build) &
(cd ./Backend/AuthService && docker compose up -d --build) &
(cd ./Backend/TaskStorageService && docker compose up -d --build) &
(cd ./Backend/NotifyService && docker compose up -d --build) &
(cd ./Frontend && docker compose up -d --build) &

wait