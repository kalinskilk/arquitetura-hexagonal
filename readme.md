docker compose up --build -d

docker exec -it appproduct bash

go test ./...
