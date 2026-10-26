# Copilot instructions

## Build, test, and lint

- Build all Go packages: `go build ./...`
- Run all tests: `go test ./...`
- Run one test by exact name in a package: `go test ./internal/database -run '^TestName$' -count=1` (replace `TestName` with the test function name).
- Run static analysis: `go vet ./...`
- Format changed Go files with `gofmt -w <file.go>`.
- Start the local PostgreSQL service with `docker compose up -d postgres`. Compose reads database settings from `.env`; initialize it from `.env.example` when needed. The Go process reads its settings directly from environment variables, so provide those variables to the process separately.

## Architecture

This is a Go API module (`github.com/Barrun-Kabir-Rifat/letseat`). The executable belongs under `cmd/api`. The internal package boundaries indicate the intended request flow: routes dispatch to handlers, handlers use services, and services use repositories and models. Keep database-specific connection and persistence concerns in the database/repository layer rather than in HTTP handlers.

The current implementation is still a scaffold: `internal/database` contains the PostgreSQL pool connection, while the other internal packages and the API entry point do not yet contain implementations. Treat existing code as authoritative as those layers are filled in.

PostgreSQL is the local data store. `docker-compose.yml` uses PostgreSQL 17, binds the service to localhost on port 5432, and persists its data in a named volume. Database schema changes belong in `migrations/` and should stay aligned with the models and repository queries that use them.

## Repository conventions

- Configure the database connection through `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`, and `DB_SSLMODE`; the example values are in `.env.example`. The API's `SERVER_PORT` is also documented there.
- Use `pgx/v5` and its connection pool for PostgreSQL access. `database.Connect` bounds connection establishment and pinging with a five-second context, wraps errors with operation context, and closes the pool if the ping fails; preserve those cleanup and error-reporting behaviors.
- Keep dependencies flowing through the package boundaries above instead of letting handlers issue database queries directly.
