# GoInfra

Production-grade web server built with Go, PostgreSQL, Nginx, Docker & Prometheus/Grafana.

## Architecture
Internet → Nginx (load balancer)

↓

Go Server x2 (app1 + app2)

↓

PostgreSQL

↓

Prometheus → Grafana

## Stack

- **Go** — HTTP server, REST API
- **PostgreSQL** — Database with connection pooling
- **Nginx** — Reverse proxy + load balancing
- **Docker Compose** — Multi-service orchestration
- **Prometheus** — Metrics scraping
- **Grafana** — Monitoring dashboard
- **GitHub Actions** — CI/CD pipeline

## Run

```bash
git clone https://github.com/ismail-hafiane/goinfra.git
cd goinfra
docker compose up --build -d
```

## Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | /health | Health check |
| GET | /api/users | List users |
| POST | /api/users | Create user |
| GET | /metrics | Prometheus metrics |

## Services

| Service | URL |
|---------|-----|
| API | http://localhost |
| Prometheus | http://localhost:9090 |
| Grafana | http://localhost:3000 |
