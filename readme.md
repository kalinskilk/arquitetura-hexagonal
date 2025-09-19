# Arquitetura Hexagonal com Golang

Este projeto é um exemplo de **Arquitetura Hexagonal** implementada em **Golang**, com suporte a CLI, HTTP e banco SQLite.

---

## Pré-requisitos

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Inicializando o projeto com Docker

Para construir e subir os containers:

```bash
docker compose up --build -d
```

---

Para acessar o container da aplicação:

```bash
docker exec -it appproduct bash
```

Para acessar o container como root:

```bash
docker exec -it -u root appproduct bash
```

## Testes

Executar todos os testes:

```bash
go test ./...
```

## CLI

### Criar um produto via CLI

```bash
go run main.go cli -a=create -n="product cli" -p=25.0
```

### Obter um produto via CLI

```bash
go run main.go cli -a=get --id=uuid
```

## HTTP

```bash
go run main.go http
```

### Criar produto

```bash
curl -X POST http://localhost:9000/product \
  -H "Content-Type: application/json" \
  -d '{ "name": "product cli", "price": 25.0 }'
```

### Buscar produto

```bash
curl http://localhost:9000/product/uuid
```

### Habilitar produto

```bash
curl http://localhost:9000/product/uuid/enable
```

### Desabilitar produto

```bash
curl http://localhost:9000/product/uuid/disable
```

## Arquitetura de Pastas

A aplicação segue o padrão de Arquitetura Hexagonal, separando claramente as regras de negócio (core) das implementações externas (adapters).

```bash
/adapters
  /cli
  /db
  /dto
  /web

/application
  product_service.go
  product.go
  /mocks

/cmd
  cli

```

### Pasta /adapters

Contém todas as implementações externas, ou seja, como a aplicação se conecta com o “mundo externo”.

### /adapters/cli

Implementa a interface de linha de comando (CLI).
Permite criar, buscar e manipular produtos via terminal.

### /adapters/db

Implementa a persistência de dados.
Aqui ficam os repositórios concretos (ex.: SQLite) que sabem salvar e buscar produtos.

### /adapters/dto

Contém os objetos de transporte de dados (Data Transfer Objects).
São usados para receber e enviar dados entre camadas ou expor via APIs.

### /adapters/web

Implementa a interface HTTP (REST API).
Permite expor a aplicação via endpoints como POST /product, GET /product/{id}, etc.

## Pasta /application

Contém a regra de negócio da aplicação, independente de tecnologia.
É o “núcleo” da Arquitetura Hexagonal.

### product.go

Entidade principal Product (com atributos id, name, price, status).
Aqui ficam as regras que garantem a consistência do produto.

### product_service.go

Casos de uso (application services) relacionados a produtos.
Ex.: criar produto, habilitar/desabilitar, consultar produto.

### /mocks

Implementações falsas (mockadas) para facilitar testes unitários, sem depender de banco ou infraestrutura externa.

## pasta /cmd

Ponto de entrada da aplicação.
Cada subpasta aqui define como a aplicação será executada.

### /cmd/cli

Ponto de entrada da aplicação em modo CLI.
Usa os adapters e a camada de application para rodar comandos como:

```bash
go run main.go cli -a=create -n="product cli" -p=25.0
```
