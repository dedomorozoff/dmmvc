# 🚀 Развертывание DMMVC

## Локальная разработка

### Требования
- Go 1.20 или выше
- Git (опционально)

### Установка

1. **Перейдите в директорию проекта**
```bash
cd c:\cygwin64\home\alexl\dmmvc
```

2. **Установите зависимости**
```bash
go mod tidy
```

3. **Запустите сервер**
```bash
go run cmd/server/main.go
```

4. **Откройте браузер**
```
http://localhost:8080
```

## Production развертывание

### 1. Сборка бинарника

```bash
# Windows
go build -o dmmvc.exe cmd/server/main.go

# Linux/Mac
go build -o dmmvc cmd/server/main.go
```

### 2. Настройка .env для production

```env
PORT=8080
GIN_MODE=release

DB_TYPE=mysql
DB_DSN=user:password@tcp(localhost:3306)/dmmvc?charset=utf8mb4&parseTime=True&loc=Local

SESSION_SECRET=your-very-strong-secret-key-here

LOG_LEVEL=warn
LOG_FILE=/var/log/dmmvc/app.log

DEBUG=false
```

### 3. Запуск

```bash
./dmmvc
```

## Docker развертывание

### Создайте Dockerfile

```dockerfile
FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o dmmvc cmd/server/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/dmmvc .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static
COPY --from=builder /app/.env .

EXPOSE 8080

CMD ["./dmmvc"]
```

### Создайте docker-compose.yml

```yaml
version: '3.8'

services:
  app:
    build: .
    ports:
      - "8080:8080"
    environment:
      - DB_TYPE=mysql
      - DB_DSN=root:password@tcp(db:3306)/dmmvc?charset=utf8mb4&parseTime=True&loc=Local
    depends_on:
      - db
    volumes:
      - ./logs:/var/log/dmmvc

  db:
    image: mysql:8.0
    environment:
      MYSQL_ROOT_PASSWORD: password
      MYSQL_DATABASE: dmmvc
    volumes:
      - mysql_data:/var/lib/mysql
    ports:
      - "3306:3306"

volumes:
  mysql_data:
```

### Запуск с Docker

```bash
docker-compose up -d
```

## Systemd сервис (Linux)

### Создайте файл /etc/systemd/system/dmmvc.service

```ini
[Unit]
Description=DMMVC Web Application
After=network.target

[Service]
Type=simple
User=www-data
WorkingDirectory=/opt/dmmvc
ExecStart=/opt/dmmvc/dmmvc
Restart=on-failure
RestartSec=5s

[Install]
WantedBy=multi-user.target
```

### Управление сервисом

```bash
# Запуск
sudo systemctl start dmmvc

# Остановка
sudo systemctl stop dmmvc

# Перезапуск
sudo systemctl restart dmmvc

# Автозапуск
sudo systemctl enable dmmvc

# Статус
sudo systemctl status dmmvc

# Логи
sudo journalctl -u dmmvc -f
```

## Nginx reverse proxy

### Конфигурация Nginx

```nginx
server {
    listen 80;
    server_name yourdomain.com;

    location / {
        proxy_pass http://localhost:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }

    location /static/ {
        alias /opt/dmmvc/static/;
        expires 30d;
        add_header Cache-Control "public, immutable";
    }
}
```

### SSL с Let's Encrypt

```bash
# Установка certbot
sudo apt install certbot python3-certbot-nginx

# Получение сертификата
sudo certbot --nginx -d yourdomain.com

# Автообновление
sudo certbot renew --dry-run
```

## Мониторинг и логирование

### 1. Логирование

Логи сохраняются в файл, указанный в `LOG_FILE`:

```bash
# Просмотр логов
tail -f /var/log/dmmvc/app.log

# Ротация логов (logrotate)
sudo nano /etc/logrotate.d/dmmvc
```

```
/var/log/dmmvc/*.log {
    daily
    rotate 7
    compress
    delaycompress
    notifempty
    create 0640 www-data www-data
    sharedscripts
}
```

### 2. Мониторинг

Используйте инструменты мониторинга:
- Prometheus + Grafana
- New Relic
- Datadog

## Резервное копирование

### База данных

```bash
# MySQL backup
mysqldump -u root -p dmmvc > backup_$(date +%Y%m%d).sql

# Восстановление
mysql -u root -p dmmvc < backup_20240101.sql
```

### SQLite backup

```bash
# Копирование файла БД
cp dmmvc.db dmmvc_backup_$(date +%Y%m%d).db
```

## Безопасность

### 1. Firewall

```bash
# UFW (Ubuntu)
sudo ufw allow 22/tcp
sudo ufw allow 80/tcp
sudo ufw allow 443/tcp
sudo ufw enable
```

### 2. Обновления

```bash
# Обновление системы
sudo apt update && sudo apt upgrade

# Обновление Go
# Скачайте новую версию с golang.org
```

### 3. Безопасность приложения

- ✅ Используйте сильный `SESSION_SECRET`
- ✅ Смените пароль администратора
- ✅ Используйте HTTPS
- ✅ Регулярно обновляйте зависимости
- ✅ Ограничьте доступ к БД

## Производительность

### 1. Оптимизация Go

```bash
# Сборка с оптимизацией
go build -ldflags="-s -w" -o dmmvc cmd/server/main.go
```

### 2. Кеширование

Добавьте Redis для кеширования:

```go
import "github.com/go-redis/redis/v8"

var rdb = redis.NewClient(&redis.Options{
    Addr: "localhost:6379",
})
```

### 3. CDN

Используйте CDN для статических файлов:
- Cloudflare
- AWS CloudFront
- Fastly

## Масштабирование

### Горизонтальное масштабирование

```yaml
# docker-compose.yml
services:
  app:
    build: .
    deploy:
      replicas: 3
    
  nginx:
    image: nginx
    depends_on:
      - app
```

### Load Balancer

Используйте:
- Nginx
- HAProxy
- AWS ELB

## Troubleshooting

### Проблема: Сервер не запускается

```bash
# Проверьте логи
tail -f dmmvc.log

# Проверьте порт
netstat -tulpn | grep 8080

# Проверьте права
ls -la dmmvc
chmod +x dmmvc
```

### Проблема: Ошибка подключения к БД

```bash
# Проверьте MySQL
sudo systemctl status mysql

# Проверьте подключение
mysql -u root -p

# Проверьте .env файл
cat .env
```

### Проблема: 502 Bad Gateway (Nginx)

```bash
# Проверьте, запущено ли приложение
ps aux | grep dmmvc

# Проверьте логи Nginx
tail -f /var/log/nginx/error.log
```

## Чеклист развертывания

- [ ] Сборка приложения
- [ ] Настройка .env для production
- [ ] Настройка базы данных
- [ ] Настройка веб-сервера (Nginx)
- [ ] Настройка SSL сертификата
- [ ] Настройка firewall
- [ ] Настройка логирования
- [ ] Настройка резервного копирования
- [ ] Смена пароля администратора
- [ ] Тестирование всех функций
- [ ] Настройка мониторинга

---

**Готово!** Ваше приложение готово к production! 🚀
