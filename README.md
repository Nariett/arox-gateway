# Arox

This `.env` file is used to configure the PostgreSQL database connection for the **Arox** project.

## Environment Variables

| Variable    | Description                                 | Example Value        |
|-------------|---------------------------------------------|----------------------|
| `USER`      | PostgreSQL username                         | `postgres`           |
| `PASSWORD`  | Password for the specified user             | `postgres`           |
| `DBNAME`    | Name of the database                        | `arox-db`            |
| `SSLMODE`   | SSL connection mode                         | `disable` (default)  |
| `PORT`      | Port on which PostgreSQL is running         | `8080`               |
| `HOST`      | Host where the database is located          | `localhost`          |

## Example Connection String

```text
postgres://postgres:postgres@localhost:8080/arox-db?sslmode=disable
