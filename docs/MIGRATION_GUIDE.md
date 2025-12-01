[English](MIGRATION_GUIDE.md) | [📚 Docs](README.md)

# Database Migration Guide

Guide for migrating between different database systems in DMMVC.

## Quick Reference

| From | To | Difficulty | Data Loss Risk |
|------|-----|-----------|----------------|
| SQLite → PostgreSQL | Easy | Low |
| SQLite → MySQL | Easy | Low |
| MySQL → PostgreSQL | Medium | Low |
| PostgreSQL → MySQL | Medium | Medium |
| Any → SQLite | Easy | Low |

## Migration Methods

### Method 1: GORM Auto-Migration (Recommended for Development)

This method recreates tables in the new database. **Data will be lost.**

1. **Backup current database**
2. **Change database configuration** in `.env`:
   ```env
   DB_TYPE=postgres
   DB_DSN=host=localhost user=postgres password=postgres dbname=dmmvc port=5432 sslmode=disable
   ```
3. **Run application** - GORM will create tables automatically
4. **Manually migrate data** if needed

### Method 2: Export/Import (Preserves Data)

#### SQLite → PostgreSQL

**Step 1: Export SQLite data**
```bash
# Install sqlite3 command-line tool
sqlite3 dmmvc.db .dump > dump.sql
```

**Step 2: Convert SQL syntax**

SQLite and PostgreSQL have different syntax. Common changes:
- `AUTOINCREMENT` → `SERIAL` or `BIGSERIAL`
- `INTEGER PRIMARY KEY` → `SERIAL PRIMARY KEY`
- `DATETIME` → `TIMESTAMP`
- Remove SQLite-specific pragmas

**Step 3: Import to PostgreSQL**
```bash
# Create database
psql -U postgres -c "CREATE DATABASE dmmvc;"

# Import (may need manual fixes)
psql -U postgres -d dmmvc < dump.sql
```

**Step 4: Update configuration**
```env
DB_TYPE=postgres
DB_DSN=host=localhost user=postgres password=postgres dbname=dmmvc port=5432 sslmode=disable
```

#### MySQL → PostgreSQL

**Step 1: Export MySQL data**
```bash
mysqldump -u user -p --compatible=postgresql dmmvc > dump.sql
```

**Step 2: Convert SQL syntax**

Use tools like:
- [pgloader](https://pgloader.io/) (recommended)
- Manual conversion

**Using pgloader:**
```bash
# Install pgloader
# Ubuntu: sudo apt install pgloader
# macOS: brew install pgloader

# Create migration file: mysql-to-postgres.load
cat > mysql-to-postgres.load << EOF
LOAD DATABASE
  FROM mysql://user:password@localhost/dmmvc
  INTO postgresql://postgres:password@localhost/dmmvc
  WITH include drop, create tables, create indexes, reset sequences
  SET maintenance_work_mem to '128MB', work_mem to '12MB';
EOF

# Run migration
pgloader mysql-to-postgres.load
```

**Step 3: Update configuration**
```env
DB_TYPE=postgres
DB_DSN=host=localhost user=postgres password=postgres dbname=dmmvc port=5432 sslmode=disable
```

#### PostgreSQL → MySQL

**Step 1: Export PostgreSQL data**
```bash
pg_dump -U postgres dmmvc > dump.sql
```

**Step 2: Convert SQL syntax**

PostgreSQL to MySQL conversion requires changes:
- `SERIAL` → `AUTO_INCREMENT`
- `BOOLEAN` → `TINYINT(1)`
- `TEXT` → `LONGTEXT`
- Remove PostgreSQL-specific features (arrays, JSONB, etc.)

**Step 3: Import to MySQL**
```bash
mysql -u user -p -e "CREATE DATABASE dmmvc;"
mysql -u user -p dmmvc < dump.sql
```

### Method 3: Application-Level Migration (Best for Production)

Create a migration script that reads from old DB and writes to new DB.

**Example: migrate.go**
```go
package main

import (
    "log"
    "dmmvc/internal/models"
    
    "github.com/glebarez/sqlite"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func main() {
    // Connect to source database (SQLite)
    sourceDB, err := gorm.Open(sqlite.Open("dmmvc.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to source DB:", err)
    }
    
    // Connect to destination database (PostgreSQL)
    destDB, err := gorm.Open(postgres.Open(
        "host=localhost user=postgres password=postgres dbname=dmmvc port=5432 sslmode=disable",
    ), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to destination DB:", err)
    }
    
    // Auto-migrate destination
    destDB.AutoMigrate(&models.User{})
    
    // Migrate users
    var users []models.User
    sourceDB.Find(&users)
    
    for _, user := range users {
        // Reset ID to let destination DB auto-generate
        user.ID = 0
        if err := destDB.Create(&user).Error; err != nil {
            log.Printf("Failed to migrate user %s: %v", user.Username, err)
        } else {
            log.Printf("Migrated user: %s", user.Username)
        }
    }
    
    log.Println("Migration completed!")
}
```

Run migration:
```bash
go run migrate.go
```

## Pre-Migration Checklist

- [ ] Backup current database
- [ ] Test migration on development environment first
- [ ] Check application compatibility with new database
- [ ] Verify all data types are supported
- [ ] Plan for downtime (if needed)
- [ ] Prepare rollback plan
- [ ] Update connection strings
- [ ] Test application after migration

## Post-Migration Checklist

- [ ] Verify all tables exist
- [ ] Check row counts match
- [ ] Test all application features
- [ ] Verify indexes are created
- [ ] Check foreign key constraints
- [ ] Test authentication/login
- [ ] Monitor performance
- [ ] Update documentation

## Common Issues

### Character Encoding

**Problem:** Special characters appear corrupted

**Solution:**
```sql
-- PostgreSQL: Set client encoding
SET client_encoding = 'UTF8';

-- MySQL: Use utf8mb4
ALTER DATABASE dmmvc CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

### Auto-Increment IDs

**Problem:** ID sequences don't match

**Solution (PostgreSQL):**
```sql
-- Reset sequence
SELECT setval('users_id_seq', (SELECT MAX(id) FROM users));
```

**Solution (MySQL):**
```sql
-- Reset auto-increment
ALTER TABLE users AUTO_INCREMENT = 1;
```

### Timestamp Formats

**Problem:** Datetime/timestamp format differences

**Solution:** Use GORM's time.Time type, which handles conversions automatically.

### Boolean Values

**Problem:** Boolean representation differs

**Solution:** GORM handles this automatically, but in raw SQL:
- PostgreSQL: `TRUE`/`FALSE`
- MySQL: `1`/`0`
- SQLite: `1`/`0`

## Testing Migration

### 1. Test Connection

```bash
go run scripts/test-db-connection.go
```

### 2. Verify Tables

```sql
-- PostgreSQL
\dt

-- MySQL
SHOW TABLES;

-- SQLite
.tables
```

### 3. Check Data

```sql
SELECT COUNT(*) FROM users;
SELECT * FROM users LIMIT 5;
```

### 4. Test Application

```bash
go run cmd/server/main.go
```

Try:
- Login
- Create/Read/Update/Delete operations
- All application features

## Rollback Plan

If migration fails:

1. **Stop application**
2. **Restore from backup**
   ```bash
   # PostgreSQL
   psql -U postgres -d dmmvc < backup.sql
   
   # MySQL
   mysql -u user -p dmmvc < backup.sql
   
   # SQLite
   cp backup.db dmmvc.db
   ```
3. **Revert configuration** in `.env`
4. **Restart application**

## Performance Considerations

### PostgreSQL
- Use connection pooling
- Create indexes on frequently queried columns
- Use EXPLAIN ANALYZE for query optimization

### MySQL
- Configure InnoDB buffer pool
- Use appropriate indexes
- Monitor slow query log

### SQLite
- Enable WAL mode for better concurrency
- Use PRAGMA optimize
- Keep database file on fast storage

## Tools

- **pgloader** - MySQL/SQLite to PostgreSQL migration
- **DBeaver** - Universal database tool
- **pgAdmin** - PostgreSQL administration
- **MySQL Workbench** - MySQL administration
- **DB Browser for SQLite** - SQLite GUI

## Resources

- [GORM Documentation](https://gorm.io/docs/)
- [PostgreSQL Migration Guide](https://www.postgresql.org/docs/current/migration.html)
- [pgloader Documentation](https://pgloader.readthedocs.io/)
- [MySQL to PostgreSQL Migration](https://wiki.postgresql.org/wiki/Converting_from_other_Databases_to_PostgreSQL)

## Support

If you need help with migration:
1. Check this guide
2. Test on development environment first
3. Open an issue on GitHub with details
