# Auth

## Installation

### Build

`docker compose up -d --build`

### Migrations

`docker compose exec ms-auth bash`

`sh install-migrateCLI.sh`

`make migrate_up`

## Feature:

-   Postgres
-   Redis
-   GRpc
-   Swagger
