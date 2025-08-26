docker compose up --build -d

docker exec -it appproduct bash

docker exec -it -u root appproduct bash

go test ./...

# Criar tabela via sqlite3

sqlite3 sqlite.db

create table products(id string,name string,price float, status string);
