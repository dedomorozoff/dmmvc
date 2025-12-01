[English](POSTGRESQL.md) | [Русский](POSTGRESQL.ru.md) | [📚 Docs](README.md)

# PostgreSQL Support in DMMVC

DMMVC now supports PostgreSQL as a database backend alongside SQLite and MySQL.

## Quick Start

### 1. Install PostgreSQL

**Windows:**
Download from [postgresql.org](https://www.postgresql.org/download/windows/)

**Linux (Ubuntu/Debian):**
```bash
sudo apt update
sudo apt install postgresql postgresql-contrib
```

**macOS:**
```bash
brew install postgresql
brew services start postgresql
```

### 2. Create Database

```bash
# Connect to PostgreSQL
psql -U postgres

# Create database
CREATE DATABASE dmmvc;

# Create user (optional)
CREATE USER dmmvc_user WITH PASSWORD 'your_password';
GRANT ALL PRIVILEGES ON DATABASE dmmvc TO dmmvc_user;

# Exit
\q
```

### 3. Configure DMMVC

Edit your `.env` file:

```env
DB_TYPE=postgres
DB_DSN=host=localhost user=postgres password=postgres dbname=dmmvc port=5432 sslmode=disable
```

Or use connection URL format:

```env
DB_TYPE=postgres
DB_DSN=postgres://postgres:postgres@localhost:5432/dmmvc?sslmode=disable
```

### 4. Run Application

```bash
go run cmd/server/main.go
```

DMMVC will automatically create tables and seed initial data.

## Configuration Options

### DSN Format

**Standard format:**
```
host=localhost user=myuser password=mypass dbname=mydb port=5432 sslmode=disable TimeZone=UTC
```

**URL format:**
```
postgres://user:password@host:port/database?sslmode=disable
```

### Common Parameters

| Parameter | Description | Default | Example |
|-----------|-------------|---------|---------|
| `host` | Database host | localhost | localhost, 192.168.1.100 |
| `port` | Database port | 5432 | 5432 |
| `user` | Database user | postgres | myuser |
| `password` | User password | - | mypassword |
| `dbname` | Database name | - | dmmvc |
| `sslmode` | SSL mode | disable | disable, require, verify-full |
| `TimeZone` | Timezone | UTC | UTC, America/New_York |

### SSL Modes

- `disable` - No SSL (development)
- `require` - SSL required but no verification
- `verify-ca` - Verify CA certificate
- `verify-full` - Full SSL verification (production)

## Production Configuration

For production, use secure settings:

```env
DB_TYPE=postgres
DB_DSN=host=db.example.com user=dmmvc_prod password=strong_password dbname=dmmvc_prod port=5432 sslmode=verify-full TimeZone=UTC
```

### Connection Pool Settings

You can configure connection pool in `internal/database/db.go`:

```go
sqlDB, err := DB.DB()
if err != nil {
    log.Fatal(err)
}

// SetMaxIdleConns sets the maximum number of connections in the idle connection pool
sqlDB.SetMaxIdleConns(10)

// SetMaxOpenConns sets the maximum number of open connections to the database
sqlDB.SetMaxOpenConns(100)

// SetConnMaxLifetime sets the maximum amount of time a connection may be reused
sqlDB.SetConnMaxLifetime(time.Hour)
```

## Docker Setup

See `docker/docker-compose.postgres.yml` for complete setup.

### Example Configuration

```yaml
version: '3.8'

services:
  postgres:
    image: postgres:16-alpine
    environment:
      POSTGRES_USER: dmmvc
      POSTGRES_PASSWORD: dmmvc_password
      POSTGRES_DB: dmmvc
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U dmmvc"]
      interval: 10s
      timeout: 5s
      retries: 5

  app:
    build: .
    ports:
      - "8080:8080"
    environment:
      DB_TYPE: postgres
      DB_DSN: host=postgres user=dmmvc password=dmmvc_password dbname=dmmvc port=5432 sslmode=disable
      SESSION_SECRET: your-secret-key
    depends_on:
      postgres:
        condition: service_healthy

volumes:
  postgres_data:
```

Run with:
```bash
docker-compose up -d
```

## Migration from SQLite/MySQL

### 1. Export Data (if needed)

**From SQLite:**
```bash
# Using GORM, data will be automatically migrated
# Just change DB_TYPE and run the app
```

**From MySQL:**
```bash
# Export MySQL data
mysqldump -u user -p dmmvc > backup.sql

# Import to PostgreSQL (may need adjustments)
psql -U postgres -d dmmvc < backup.sql
```

### 2. Update Configuration

Change `.env`:
```env
DB_TYPE=postgres
DB_DSN=host=localhost user=postgres password=postgres dbname=dmmvc port=5432 sslmode=disable
```

### 3. Run Migrations

```bash
go run cmd/server/main.go
```

GORM will automatically create tables in PostgreSQL.

## PostgreSQL-Specific Features

### Full-Text Search

```go
// Example: Search in posts
var posts []models.Post
DB.Where("to_tsvector('english', title || ' ' || content) @@ plainto_tsquery('english', ?)", "search term").Find(&posts)
```

### JSON Support

```go
type Product struct {
    gorm.Model
    Name       string
    Attributes datatypes.JSON `gorm:"type:jsonb"` // PostgreSQL JSONB
}

// Query JSON
DB.Where("attributes->>'color' = ?", "red").Find(&products)
```

### Arrays

```go
import "github.com/lib/pq"

type Article struct {
    gorm.Model
    Title string
    Tags  pq.StringArray `gorm:"type:text[]"`
}
```

## Troubleshooting

### Connection Refused

```
Error: connection refused
```

**Solution:**
- Check if PostgreSQL is running: `sudo systemctl status postgresql`
- Start PostgreSQL: `sudo systemctl start postgresql`
- Check port: `sudo netstat -plnt | grep 5432`

### Authentication Failed

```
Error: password authentication failed
```

**Solution:**
- Check username and password in `.env`
- Reset password: `ALTER USER postgres PASSWORD 'newpassword';`
- Check `pg_hba.conf` authentication method

### Database Does Not Exist

```
Error: database "dmmvc" does not exist
```

**Solution:**
```bash
psql -U postgres
CREATE DATABASE dmmvc;
```

### SSL Error

```
Error: SSL is not enabled on the server
```

**Solution:**
Add `sslmode=disable` to DSN for development:
```
DB_DSN=host=localhost user=postgres password=postgres dbname=dmmvc port=5432 sslmode=disable
```

## Performance Tips

1. **Use Indexes:**
```go
type User struct {
    gorm.Model
    Email string `gorm:"uniqueIndex"`
    Name  string `gorm:"index"`
}
```

2. **Use Connection Pooling:**
```go
sqlDB.SetMaxOpenConns(25)
sqlDB.SetMaxIdleConns(5)
```

3. **Use Prepared Statements:**
GORM uses prepared statements by default.

4. **Batch Operations:**
```go
DB.CreateInBatches(users, 100)
```

## Useful Commands

```bash
# Connect to database
psql -U postgres -d dmmvc

# List databases
\l

# List tables
\dt

# Describe table
\d users

# Show table size
SELECT pg_size_pretty(pg_total_relation_size('users'));

# Show active connections
SELECT * FROM pg_stat_activity;

# Backup database
pg_dump -U postgres dmmvc > backup.sql

# Restore database
psql -U postgres dmmvc < backup.sql
```

## Resources

- [PostgreSQL Documentation](https://www.postgresql.org/docs/)
- [GORM PostgreSQL Driver](https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL)
- [pgx Driver](https://github.com/jackc/pgx)

## Support

If you encounter issues with PostgreSQL support, please:
1. Check this documentation
2. Review PostgreSQL logs: `sudo tail -f /var/log/postgresql/postgresql-*.log`
3. Open an issue on GitHub with error details
