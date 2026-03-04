# Dependencies & Development Tools

This file lists the runtime dependencies and developer tools used by the SRTBackend project.

## Go module
- Primary module: `github.com/akpatri/srt` — see `go.mod` for exact module dependencies and versions.

## Runtime services (required to run the app locally)
- PostgreSQL — database for GORM (export `DATABASE_URL` or use `.env`).
- Redis — caching / transient storage.
- NATS — event messaging for realtime pipeline.

Recommended quick start (Docker):
```bash
# start Postgres, Redis, NATS
docker run -d --name srt-postgres -e POSTGRES_PASSWORD=pass -e POSTGRES_DB=srt -p 5432:5432 postgres:15
docker run -d --name srt-redis -p 6379:6379 redis:7
docker run -d --name srt-nats -p 4222:4222 nats:2
```

## Developer tools (install with `go install`)
- golangci-lint (linters):
  ```bash
  go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
  ```
- mockgen (mocks):
  ```bash
  go install github.com/golang/mock/mockgen@latest
  ```
- goimports (format + organize imports):
  ```bash
  go install golang.org/x/tools/cmd/goimports@latest
  ```
- gofumpt (strict formatting):
  ```bash
  go install mvdan.cc/gofumpt@latest
  ```

After installing tools, run linters/formatters from the module directory (`SRTBackend`):
```bash
cd SRTBackend
gofumpt -w .
goimports -w .
golangci-lint run ./...
```

## Common commands
- Build: `cd SRTBackend && go build ./...`
- Vet: `cd SRTBackend && go vet ./...`
- Tidy modules: `cd SRTBackend && go mod tidy`

If you want, I can add a `docker-compose.yml` and Makefile targets to automate the local setup.
