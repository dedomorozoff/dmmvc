[English](DOCKER.md) | [📚 Docs](README.md)

# Docker Deployment Guide

This guide explains how to deploy DMMVC using Docker and Docker Compose.

## Prerequisites

- Docker 20.10+
- Docker Compose 2.0+

## Quick Start

### 1. Using Docker Compose with PostgreSQL

```bash
# Start all services (app + PostgreSQL + pgAdmin)
docker-compose -f docker/docker-compose.postgres.yml up -d

# View logs
docker-compose -f docker/docker-compose.postgres.yml logs -f

# Stop services
docker-compose -f docker/docker-compose.postgres.yml down

# Stop and remove volumes (WARNING: deletes database data)
docker-compose -f docker/docker-compose.postgres.yml down -v
```

### 2. Access Services

- **Application**: http://localhost:8080
- **pgAdmin**: http://localhost:5050
  - Email: `admin@dmmvc.local`
  - Password: `admin`

### 3. Connect to PostgreSQL from pgAdmin

1. Open pgAdmin at http://localhost:5050
2. Right-click "Servers" → "Register" → "Server"
3. General tab:
   - Name: `DMMVC`
4. Connection tab:
   - Host: `postgres`
   - Port: `5432`
   - Database: `dmmvc`
   - Username: `dmmvc`
   - Password: `dmmvc_password`
5. Click "Save"

## Building Custom Image

### Build Application Image

```bash
# Build image
docker build -t dmmvc:latest .

# Run container
docker run -d \
  -p 8080:8080 \
  -e DB_TYPE=sqlite \
  -e DB_DSN=dmmvc.db \
  -e SESSION_SECRET=your-secret \
  --name dmmvc_app \
  dmmvc:latest
```

### Build with Different Database

**SQLite (default):**
```bash
docker run -d \
  -p 8080:8080 \
  -e DB_TYPE=sqlite \
  -e DB_DSN=dmmvc.db \
  -v $(pwd)/data:/app/data \
  dmmvc:latest
```

**MySQL:**
```bash
docker run -d \
  -p 8080:8080 \
  -e DB_TYPE=mysql \
  -e DB_DSN="user:password@tcp(mysql-host:3306)/dmmvc?charset=utf8mb4&parseTime=True&loc=Local" \
  dmmvc:latest
```

**PostgreSQL:**
```bash
docker run -d \
  -p 8080:8080 \
  -e DB_TYPE=postgres \
  -e DB_DSN="host=postgres-host user=dmmvc password=password dbname=dmmvc port=5432 sslmode=disable" \
  dmmvc:latest
```

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | Server port | 8080 |
| `GIN_MODE` | Gin mode (debug/release) | debug |
| `DB_TYPE` | Database type (sqlite/mysql/postgres) | sqlite |
| `DB_DSN` | Database connection string | dmmvc.db |
| `SESSION_SECRET` | Session secret key | - |
| `LOG_LEVEL` | Log level (debug/info/warn/error) | info |
| `LOG_FILE` | Log file path | dmmvc.log |
| `DEBUG` | Debug mode | true |

## Production Deployment

### 1. Update Configuration

Edit `docker-compose.postgres.yml`:

```yaml
environment:
  GIN_MODE: release
  SESSION_SECRET: your-very-strong-secret-key-here
  DEBUG: false
  POSTGRES_PASSWORD: strong-database-password
```

### 2. Use Secrets (Recommended)

Create `.env` file:
```env
POSTGRES_PASSWORD=strong-password
SESSION_SECRET=strong-secret
```

Update docker-compose.yml:
```yaml
environment:
  POSTGRES_PASSWORD: ${POSTGRES_PASSWORD}
  SESSION_SECRET: ${SESSION_SECRET}
```

### 3. Enable SSL for PostgreSQL

```yaml
postgres:
  environment:
    POSTGRES_PASSWORD: ${POSTGRES_PASSWORD}
  volumes:
    - ./certs:/var/lib/postgresql/certs
  command: >
    -c ssl=on
    -c ssl_cert_file=/var/lib/postgresql/certs/server.crt
    -c ssl_key_file=/var/lib/postgresql/certs/server.key
```

Update DSN:
```
DB_DSN=host=postgres user=dmmvc password=password dbname=dmmvc port=5432 sslmode=require
```

## Backup and Restore

### Backup PostgreSQL

```bash
# Backup database
docker exec dmmvc_postgres pg_dump -U dmmvc dmmvc > backup.sql

# Or using docker-compose
docker-compose -f docker-compose.postgres.yml exec postgres pg_dump -U dmmvc dmmvc > backup.sql
```

### Restore PostgreSQL

```bash
# Restore database
docker exec -i dmmvc_postgres psql -U dmmvc dmmvc < backup.sql

# Or using docker-compose
docker-compose -f docker-compose.postgres.yml exec -T postgres psql -U dmmvc dmmvc < backup.sql
```

### Backup Volumes

```bash
# Backup PostgreSQL data volume
docker run --rm \
  -v dmmvc_postgres_data:/data \
  -v $(pwd):/backup \
  alpine tar czf /backup/postgres-backup.tar.gz -C /data .

# Restore PostgreSQL data volume
docker run --rm \
  -v dmmvc_postgres_data:/data \
  -v $(pwd):/backup \
  alpine tar xzf /backup/postgres-backup.tar.gz -C /data
```

## Monitoring

### View Logs

```bash
# All services
docker-compose -f docker-compose.postgres.yml logs -f

# Specific service
docker-compose -f docker-compose.postgres.yml logs -f app
docker-compose -f docker-compose.postgres.yml logs -f postgres

# Last 100 lines
docker-compose -f docker-compose.postgres.yml logs --tail=100 app
```

### Check Container Status

```bash
# List containers
docker-compose -f docker-compose.postgres.yml ps

# Container stats
docker stats dmmvc_app dmmvc_postgres
```

### Execute Commands in Container

```bash
# Access app container shell
docker exec -it dmmvc_app sh

# Access PostgreSQL
docker exec -it dmmvc_postgres psql -U dmmvc dmmvc

# Run SQL query
docker exec dmmvc_postgres psql -U dmmvc dmmvc -c "SELECT * FROM users;"
```

## Troubleshooting

### Container Won't Start

```bash
# Check logs
docker-compose -f docker-compose.postgres.yml logs app

# Check if port is already in use
netstat -an | grep 8080

# Remove and recreate containers
docker-compose -f docker-compose.postgres.yml down
docker-compose -f docker-compose.postgres.yml up -d
```

### Database Connection Issues

```bash
# Check if PostgreSQL is ready
docker exec dmmvc_postgres pg_isready -U dmmvc

# Check PostgreSQL logs
docker-compose -f docker-compose.postgres.yml logs postgres

# Test connection from app container
docker exec dmmvc_app ping postgres
```

### Reset Everything

```bash
# Stop and remove everything (including volumes)
docker-compose -f docker-compose.postgres.yml down -v

# Remove images
docker rmi dmmvc:latest

# Start fresh
docker-compose -f docker-compose.postgres.yml up -d --build
```

## Multi-Stage Build Optimization

The Dockerfile uses multi-stage builds to create smaller images:

- **Builder stage**: Compiles Go application (~500MB)
- **Runtime stage**: Only contains binary and assets (~20MB)

## Health Checks

The PostgreSQL service includes a health check:

```yaml
healthcheck:
  test: ["CMD-SHELL", "pg_isready -U dmmvc"]
  interval: 10s
  timeout: 5s
  retries: 5
```

The app waits for PostgreSQL to be healthy before starting.

## Scaling

To run multiple app instances:

```bash
docker-compose -f docker-compose.postgres.yml up -d --scale app=3
```

Note: You'll need a load balancer (nginx, traefik) to distribute traffic.

## Resources

- [Docker Documentation](https://docs.docker.com/)
- [Docker Compose Documentation](https://docs.docker.com/compose/)
- [PostgreSQL Docker Image](https://hub.docker.com/_/postgres)
- [pgAdmin Docker Image](https://hub.docker.com/r/dpage/pgadmin4)
