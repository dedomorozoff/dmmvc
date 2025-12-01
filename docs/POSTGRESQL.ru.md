[English](POSTGRESQL.md) | **Русский** | [Docs](README.md)

# Поддержка PostgreSQL в DMMVC

DMMVC теперь поддерживает PostgreSQL в качестве базы данных наряду с SQLite и MySQL.

## Быстрый старт

### 1. Установите PostgreSQL

**Windows:**
Скачайте с [postgresql.org](https://www.postgresql.org/download/windows/)

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

### 2. Создайте базу данных

```bash
# Подключитесь к PostgreSQL
psql -U postgres

# Создайте базу данных
CREATE DATABASE dmmvc;

# Создайте пользователя (опционально)
CREATE USER dmmvc_user WITH PASSWORD 'your_password';
GRANT ALL PRIVILEGES ON DATABASE dmmvc TO dmmvc_user;

# Выход
\q
```

### 3. Настройте DMMVC

Отредактируйте файл `.env`:

```env
DB_TYPE=postgres
DB_DSN=host=localhost user=postgres password=postgres dbname=dmmvc port=5432 sslmode=disable
```

Или используйте формат URL:

```env
DB_TYPE=postgres
DB_DSN=postgres://postgres:postgres@localhost:5432/dmmvc?sslmode=disable
```

### 4. Запустите приложение

```bash
go run cmd/server/main.go
```

DMMVC автоматически создаст таблицы и заполнит начальные данные.

## Параметры конфигурации

### Формат DSN

**Стандартный формат:**
```
host=localhost user=myuser password=mypass dbname=mydb port=5432 sslmode=disable TimeZone=UTC
```

**Формат URL:**
```
postgres://user:password@host:port/database?sslmode=disable
```

### Основные параметры

| Параметр | Описание | По умолчанию | Пример |
|----------|----------|--------------|--------|
| `host` | Хост БД | localhost | localhost, 192.168.1.100 |
| `port` | Порт БД | 5432 | 5432 |
| `user` | Пользователь | postgres | myuser |
| `password` | Пароль | - | mypassword |
| `dbname` | Имя БД | - | dmmvc |
| `sslmode` | Режим SSL | disable | disable, require, verify-full |
| `TimeZone` | Часовой пояс | UTC | UTC, Europe/Moscow |

### Режимы SSL

- `disable` - Без SSL (разработка)
- `require` - SSL обязателен, но без проверки
- `verify-ca` - Проверка CA сертификата
- `verify-full` - Полная проверка SSL (production)

## Конфигурация для production

Для production используйте безопасные настройки:

```env
DB_TYPE=postgres
DB_DSN=host=db.example.com user=dmmvc_prod password=strong_password dbname=dmmvc_prod port=5432 sslmode=verify-full TimeZone=UTC
```

### Настройки пула соединений

Вы можете настроить пул соединений в `internal/database/db.go`:

```go
sqlDB, err := DB.DB()
if err != nil {
    log.Fatal(err)
}

// Максимальное количество простаивающих соединений
sqlDB.SetMaxIdleConns(10)

// Максимальное количество открытых соединений
sqlDB.SetMaxOpenConns(100)

// Максимальное время жизни соединения
sqlDB.SetConnMaxLifetime(time.Hour)
```

## Настройка с Docker

См. `docker/docker-compose.postgres.yml` для полной настройки.

### Пример конфигурации

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

Запуск:
```bash
docker-compose up -d
```

## Миграция с SQLite/MySQL

### 1. Экспорт данных (если нужно)

**Из SQLite:**
```bash
# С помощью GORM данные будут автоматически мигрированы
# Просто измените DB_TYPE и запустите приложение
```

**Из MySQL:**
```bash
# Экспорт данных MySQL
mysqldump -u user -p dmmvc > backup.sql

# Импорт в PostgreSQL (может потребовать корректировок)
psql -U postgres -d dmmvc < backup.sql
```

### 2. Обновите конфигурацию

Измените `.env`:
```env
DB_TYPE=postgres
DB_DSN=host=localhost user=postgres password=postgres dbname=dmmvc port=5432 sslmode=disable
```

### 3. Запустите миграции

```bash
go run cmd/server/main.go
```

GORM автоматически создаст таблицы в PostgreSQL.

## Специфичные возможности PostgreSQL

### Полнотекстовый поиск

```go
// Пример: Поиск в постах
var posts []models.Post
DB.Where("to_tsvector('russian', title || ' ' || content) @@ plainto_tsquery('russian', ?)", "поисковый запрос").Find(&posts)
```

### Поддержка JSON

```go
type Product struct {
    gorm.Model
    Name       string
    Attributes datatypes.JSON `gorm:"type:jsonb"` // PostgreSQL JSONB
}

// Запрос JSON
DB.Where("attributes->>'color' = ?", "red").Find(&products)
```

### Массивы

```go
import "github.com/lib/pq"

type Article struct {
    gorm.Model
    Title string
    Tags  pq.StringArray `gorm:"type:text[]"`
}
```

## Решение проблем

### Отказ в соединении

```
Error: connection refused
```

**Решение:**
- Проверьте, запущен ли PostgreSQL: `sudo systemctl status postgresql`
- Запустите PostgreSQL: `sudo systemctl start postgresql`
- Проверьте порт: `sudo netstat -plnt | grep 5432`

### Ошибка аутентификации

```
Error: password authentication failed
```

**Решение:**
- Проверьте имя пользователя и пароль в `.env`
- Сбросьте пароль: `ALTER USER postgres PASSWORD 'newpassword';`
- Проверьте метод аутентификации в `pg_hba.conf`

### База данных не существует

```
Error: database "dmmvc" does not exist
```

**Решение:**
```bash
psql -U postgres
CREATE DATABASE dmmvc;
```

### Ошибка SSL

```
Error: SSL is not enabled on the server
```

**Решение:**
Добавьте `sslmode=disable` в DSN для разработки:
```
DB_DSN=host=localhost user=postgres password=postgres dbname=dmmvc port=5432 sslmode=disable
```

## Советы по производительности

1. **Используйте индексы:**
```go
type User struct {
    gorm.Model
    Email string `gorm:"uniqueIndex"`
    Name  string `gorm:"index"`
}
```

2. **Используйте пул соединений:**
```go
sqlDB.SetMaxOpenConns(25)
sqlDB.SetMaxIdleConns(5)
```

3. **Используйте подготовленные запросы:**
GORM использует подготовленные запросы по умолчанию.

4. **Пакетные операции:**
```go
DB.CreateInBatches(users, 100)
```

## Полезные команды

```bash
# Подключиться к базе данных
psql -U postgres -d dmmvc

# Список баз данных
\l

# Список таблиц
\dt

# Описание таблицы
\d users

# Показать размер таблицы
SELECT pg_size_pretty(pg_total_relation_size('users'));

# Показать активные соединения
SELECT * FROM pg_stat_activity;

# Резервная копия базы данных
pg_dump -U postgres dmmvc > backup.sql

# Восстановление базы данных
psql -U postgres dmmvc < backup.sql
```

## Ресурсы

- [Документация PostgreSQL](https://www.postgresql.org/docs/)
- [GORM PostgreSQL драйвер](https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL)
- [pgx драйвер](https://github.com/jackc/pgx)

## Поддержка

Если у вас возникли проблемы с поддержкой PostgreSQL:
1. Проверьте эту документацию
2. Просмотрите логи PostgreSQL: `sudo tail -f /var/log/postgresql/postgresql-*.log`
3. Откройте issue на GitHub с деталями ошибки
