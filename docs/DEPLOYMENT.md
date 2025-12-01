**English** | [Русский](DEPLOYMENT.ru.md)

# DMMVC Deployment

## Local Development

### Requirements
- Go 1.20 or higher
- Git (optional)

### Installation

1. **Navigate to project directory**
```bash
cd c:\cygwin64\home\alexl\dmmvc
```

2. **Install dependencies**
```bash
go mod tidy
```

3. **Run server**
```bash
go run cmd/server/main.go
```

4. **Open browser**
```
http://localhost:8080
```

## Production Deployment

### 1. Build Binary

```bash
# Windows
go build -o dmmvc.exe cmd/server/main.go

# Linux/Mac
go build -o dmmvc cmd/server/main.go
```

### 2. Configure .env for Production

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

### 3. Run

```bash
./dmmvc
```

## Docker Deployment

### Create Dockerfile

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

### Create docker-compose.yml

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

### Run with Docker

```bash
docker-compose up -d
```

## Systemd Service (Linux)

### Create file /etc/systemd/system/dmmvc.service

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

### Manage Service

```bash
# Start
sudo systemctl start dmmvc

# Stop
sudo systemctl stop dmmvc

# Restart
sudo systemctl restart dmmvc

# Enable on boot
sudo systemctl enable dmmvc

# Status
sudo systemctl status dmmvc

# Logs
sudo journalctl -u dmmvc -f
```

## Nginx Reverse Proxy

### Nginx Configuration

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

### SSL with Let's Encrypt

```bash
# Install certbot
sudo apt install certbot python3-certbot-nginx

# Get certificate
sudo certbot --nginx -d yourdomain.com

# Auto-renewal
sudo certbot renew --dry-run
```

## Monitoring and Logging

### 1. Logging

Logs are saved to the file specified in `LOG_FILE`:

```bash
# View logs
tail -f /var/log/dmmvc/app.log

# Log rotation (logrotate)
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

### 2. Monitoring

Use monitoring tools:
- Prometheus + Grafana
- New Relic
- Datadog

## Backup

### Database

```bash
# MySQL backup
mysqldump -u root -p dmmvc > backup_$(date +%Y%m%d).sql

# Restore
mysql -u root -p dmmvc < backup_20240101.sql
```

### SQLite Backup

```bash
# Copy DB file
cp dmmvc.db dmmvc_backup_$(date +%Y%m%d).db
```

## Security

### 1. Firewall

```bash
# UFW (Ubuntu)
sudo ufw allow 22/tcp
sudo ufw allow 80/tcp
sudo ufw allow 443/tcp
sudo ufw enable
```

### 2. Updates

```bash
# Update system
sudo apt update && sudo apt upgrade

# Update Go
# Download new version from golang.org
```

### 3. Application Security

- Use strong `SESSION_SECRET`
- Change admin password
- Use HTTPS
- Regularly update dependencies
- Limit database access

## Performance

### 1. Go Optimization

```bash
# Build with optimization
go build -ldflags="-s -w" -o dmmvc cmd/server/main.go
```

### 2. Caching

Add Redis for caching:

```go
import "github.com/go-redis/redis/v8"

var rdb = redis.NewClient(&redis.Options{
    Addr: "localhost:6379",
})
```

### 3. CDN

Use CDN for static files:
- Cloudflare
- AWS CloudFront
- Fastly

## Scaling

### Horizontal Scaling

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

Use:
- Nginx
- HAProxy
- AWS ELB

## Troubleshooting

### Issue: Server won't start

```bash
# Check logs
tail -f dmmvc.log

# Check port
netstat -tulpn | grep 8080

# Check permissions
ls -la dmmvc
chmod +x dmmvc
```

### Issue: Database connection error

```bash
# Check MySQL
sudo systemctl status mysql

# Check connection
mysql -u root -p

# Check .env file
cat .env
```

### Issue: 502 Bad Gateway (Nginx)

```bash
# Check if app is running
ps aux | grep dmmvc

# Check Nginx logs
tail -f /var/log/nginx/error.log
```

## Deployment Checklist

- [ ] Build application
- [ ] Configure .env for production
- [ ] Configure database
- [ ] Configure web server (Nginx)
- [ ] Configure SSL certificate
- [ ] Configure firewall
- [ ] Configure logging
- [ ] Configure backup
- [ ] Change admin password
- [ ] Test all functions
- [ ] Configure monitoring

---

**Done!** Your application is ready for production!
