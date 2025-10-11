# Arox

This `.env` file is used to configure the PostgreSQL database connection for the **Arox** project.

## Environment Variables

| Variable       | Description                          | Example Value        |
|----------------|--------------------------------------|----------------------|
| `DB_USER`      | PostgreSQL username                  | `postgres`           |
| `DB_PASSWORD`  | Password for the specified user      | `postgres`           |
| `DB_NAME`      | Name of the database                 | `arox-gateway`       |
| `DB_SSLMODE`   | SSL connection mode                  | `disable` (default)  |
| `DB_PORT`      | Port on which PostgreSQL is running  | `5432`               |
| `HOST`         | Host where the database is located   | `localhost`          |
| `PRODUCTS_PORT` | Port to connect with arox-products   | `8001`               |
| `JWT_SECRET`    | 128 bit secret key                  | `your generated key` |

## Example Connection String

```text
postgres://postgres:postgres@localhost:5432/arox-gateway?sslmode=disable
