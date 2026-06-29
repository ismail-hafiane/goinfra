# GoInfra 🚀

![CI/CD](https://github.com/ismail-hafiane/goinfra/actions/workflows/ci.yml/badge.svg)

Production-grade web server built with Go, PostgreSQL, Nginx, Docker & Prometheus/Grafana — deployed across 3 VMware VMs.

## Architecture
                ┌─────────────────────────────┐
                │     infra-server :30          │
                │  ┌─────────┐  ┌──────────┐  │
     HTTP ────▶ │  │  Nginx  │  │Prometheus│  │
                │  └────┬────┘  └──────────┘  │
                │       │       ┌──────────┐  │
                │       │       │ Grafana  │  │
                │  ┌────▼────┐  └──────────┘  │
                └──┤LB :80  ├─────────────────┘
                   └────┬────┘
          ┌─────────────┴─────────────┐
          ▼                           ▼
┌─────────────────────┐   ┌─────────────────────┐

│   app-server :10    │   │   app-server :10    │

│  ┌───────────────┐  │   │  ┌───────────────┐  │

│  │  app1 :8080   │  │   │  │  app2 :8081   │  │

│  └───────┬───────┘  │   │  └───────┬───────┘  │

└──────────┼──────────┘   └──────────┼──────────┘

└─────────────┬────────────┘

▼

┌─────────────────────────┐

│      db-server :20      │

│  ┌───────────────────┐  │

│  │   PostgreSQL      │  │

│  │   :5432           │  │

│  └───────────────────┘  │

└─────────────────────────┘

## Stack

| Component | Technology |
|-----------|-----------|
| Language | Go 1.26 |
| Database | PostgreSQL 16 |
| Reverse Proxy | Nginx Alpine |
| Containerization | Docker + Docker Compose |
| Monitoring | Prometheus + Grafana |
| CI/CD | GitHub Actions |
| Infrastructure | VMware (3 VMs) |

## VM Layout

| VM | IP | Role | Services |
|----|----|------|----------|
| app-server | 192.168.100.10 | Application | Go app1 :8080, app2 :8081 |
| db-server | 192.168.100.20 | Database | PostgreSQL :5432 |
| infra-server | 192.168.100.30 | Infrastructure | Nginx :80, Prometheus :9090, Grafana :3000 |

## Quick Start (Local)

```bash
git clone https://github.com/ismail-hafiane/goinfra.git
cd goinfra
docker compose up --build -d
```

## Deploy on 3 VMs

```bash
# db-server (192.168.100.20)
docker compose -f docker-compose.db.yml up -d

# app-server (192.168.100.10)
docker compose -f docker-compose.app.yml up --build -d

# infra-server (192.168.100.30)
docker compose -f docker-compose.infra.yml up -d
```

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | /health | Health check |
| GET | /api/users | List all users |
| POST | /api/users | Create a user |
| GET | /metrics | Prometheus metrics |

## Services

| Service | Local | VM |
|---------|-------|----|
| API | http://localhost | http://192.168.100.30 |
| Prometheus | http://localhost:9090 | http://192.168.100.30:9090 |
| Grafana | http://localhost:3000 | http://192.168.100.30:3000 |

## Features

- ✅ REST API with Go native `net/http`
- ✅ PostgreSQL with `pgx` connection pooling
- ✅ Nginx load balancing across 2 Go instances
- ✅ Prometheus metrics (requests/sec, latency, errors)
- ✅ Grafana NOC dashboard
- ✅ Docker multi-stage builds (~15MB image)
- ✅ GitHub Actions CI/CD pipeline
- ✅ Production deployment on 3 VMware VMs

## Author

**Ismail Hafiane** — NOC Analyst → Cloud/DevOps Engineer

[![LinkedIn](https://img.shields.io/badge/LinkedIn-ismailhafiane7-blue)](https://linkedin.com/in/ismailhafiane7)
[![GitHub](https://img.shields.io/badge/GitHub-ismail--hafiane-black)](https://github.com/ismail-hafiane)
