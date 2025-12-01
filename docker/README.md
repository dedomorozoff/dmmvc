# Docker Configuration

Docker and Docker Compose files for DMMVC.

## Files

- `docker-compose.postgres.yml` - Docker Compose with PostgreSQL, pgAdmin, and DMMVC app
- `init-db.sql` - PostgreSQL initialization script

## Quick Start

### Start with PostgreSQL

```bash
# From project root
docker-compose -f docker/docker-compose.postgres.yml up -d

# View logs
docker-compose -f docker/docker-compose.postgres.yml logs -f

# Stop
docker-compose -f docker/docker-compose.postgres.yml down
```

### Access Services

- **Application**: http://localhost:8080
- **pgAdmin**: http://localhost:5050
  - Email: `admin@dmmvc.local`
  - Password: `admin`

## Documentation

See [docs/DOCKER.md](../docs/DOCKER.md) for complete Docker documentation.

## Configuration

Edit `docker-compose.postgres.yml` to customize:
- Database credentials
- Ports
- Volumes
- Environment variables

## Production

For production deployment:
1. Change default passwords
2. Enable SSL for PostgreSQL
3. Use secrets management
4. Configure proper backup strategy

See [docs/DEPLOYMENT.md](../docs/DEPLOYMENT.md) for production setup.
