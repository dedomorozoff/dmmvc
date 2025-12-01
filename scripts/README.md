# DMMVC Scripts

Utility scripts for DMMVC development and testing.

## test-db-connection.go

Tests database connection with current configuration.

### Usage

```bash
# Test with .env configuration
go run scripts/test-db-connection.go

# Test specific database
DB_TYPE=postgres DB_DSN="host=localhost user=postgres password=postgres dbname=dmmvc port=5432 sslmode=disable" go run scripts/test-db-connection.go

# Test MySQL
DB_TYPE=mysql DB_DSN="user:password@tcp(localhost:3306)/dmmvc?charset=utf8mb4&parseTime=True&loc=Local" go run scripts/test-db-connection.go

# Test SQLite
DB_TYPE=sqlite DB_DSN="test.db" go run scripts/test-db-connection.go
```

### Output Example

```
Testing database connection...
Database Type: postgres
DSN: host=localhost user...

Connecting to PostgreSQL...

✅ Database connection successful!
Database Version: PostgreSQL 16.0 on x86_64-pc-linux-gnu...

Connection Stats:
  Open Connections: 1
  In Use: 0
  Idle: 1

✅ Test completed successfully!
```
