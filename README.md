# Test HTTP Metrics Server

Simple HTTP server written in Go for testing monitoring, health checks and Prometheus metrics collection.
---
## Endpoints

| Endpoint   | Description               |
| ---------- | ------------------------- |
| `/`        | Main page                 |
| `/health`  | Health check endpoint     |
| `/db`      | Simulated database status |
| `/random`  | Random service status     |
| `/metrics` | Prometheus metrics        |

## Run

```bash
go run ./cmd/server/main.go
```

Server starts on:
```text
http://localhost:8090
```

