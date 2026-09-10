# Go API Structure — Gin and PostgreSQL

A small **Go, Gin, and GORM example** demonstrating separation between routes, handlers, models, and database access. The implemented endpoint lists users from PostgreSQL.

## Requirements and setup

The existing [go.mod](go.mod) declares Go `1.23.2` and module name `backend-3`. Use a compatible Go toolchain and a local PostgreSQL database. Keep the module name consistent with the imports; do not run `go mod init` again for this checkout.

```sh
go mod download
```

Create a `.env` in the repository root with your development settings:

```dotenv
DB_HOST=127.0.0.1
DB_USER=your_database_user
DB_NAME=go_structure
DB_PASSWORD=your_database_password
DB_PORT=5432
```

Create the database before making a request. The database helper requires the `.env` file to load successfully and constructs a connection with `sslmode=disable`, suitable only for the intended local example configuration.

```sh
go run .
```

Gin runs on port 8080 by default, or the `PORT` process environment setting. Query the endpoint:

```sh
curl http://localhost:8080/users
```

## Request lifecycle

`main.go` registers `GET /users`. Its handler initializes the database, runs `AutoMigrate` for `models.User`, queries the users, and returns JSON. Database initialization and migration happen on every request, not once at startup. Configuration or connection failures can terminate the process.

## Files and data model

| Path | Purpose |
| --- | --- |
| [main.go](main.go) | Router and HTTP server. |
| [routes/routes.go](routes/routes.go) | Route registration. |
| [handlers/users.go](handlers/users.go) | User-list handler. |
| [db/db.go](db/db.go) | Environment, PostgreSQL, and migration setup. |
| [models/user.go](models/user.go) | User fields: `ID`, `Email`, and `CreatedAt`. |
| [utils/utils.go](utils/utils.go) | Separate numeric helper experiment. |

The earlier README included a larger hypothetical model; it was not the model implemented here. The current API does not expose creation, updates, deletion, or authentication.

## Checks and improvements

`go build ./...` checks compilation and `go test ./...` runs any package tests. No `_test.go` files are currently tracked, so that command alone would not establish behavior coverage. Database requests were not exercised for this documentation update.

Possible next improvements include sharing a single database connection, separating migrations from requests, and adding handler tests and explicit error responses.
