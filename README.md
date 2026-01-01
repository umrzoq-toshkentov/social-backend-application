# Social

A social application API built with Go.

## Prerequisites

- Go 1.21+
- Docker and Docker Compose
- [golang-migrate](https://github.com/golang-migrate/migrate) CLI
- [direnv](https://direnv.net/) (optional)

## Setup

1. Clone the repository and navigate to the project directory.

2. Start the database:
```bash
docker-compose up -d db
```

3. Load environment variables:
```bash
source .envrc
# or if using direnv
direnv allow
```

4. Run migrations:
```bash
make migrate-up
```

5. Start the API:
```bash
make run
```

The API will be available at `http://localhost:8080`.

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `ADDR` | Server address | `:8080` |
| `DB_ADDR` | PostgreSQL connection string | `postgres://postgres:postgres@localhost:5433/social?sslmode=disable` |
| `DB_MAX_OPEN_CONNS` | Max open database connections | `30` |
| `DB_MAX_IDLE_CONNS` | Max idle database connections | `30` |
| `DB_MAX_IDLE_TIME` | Max idle time for connections | `15m` |

## Makefile Commands

| Command | Description |
|---------|-------------|
| `make run` | Run the API server |
| `make migrate-up` | Apply all pending migrations |
| `make migrate-down` | Rollback the last migration |
| `make migrate-create <name>` | Create a new migration file |

## Project Structure

```
.
├── cmd/
│   ├── api/          # API application entry point
│   └── migrate/      # Database migrations
├── internal/
│   ├── db/           # Database connection
│   ├── env/          # Environment variable helpers
│   └── store/        # Data access layer
├── docker-compose.yml
├── Makefile
└── .envrc
```
