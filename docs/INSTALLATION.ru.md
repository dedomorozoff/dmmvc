# Руководство по установке DMMVC

## Системные требования

- **Go**: версия 1.20 или выше
- **Git**: для клонирования репозитория
- **Windows**: 10/11 или Windows Server

## Способы установки

### 1. Автоматическая установка (рекомендуется)

Самый простой способ установить DMMVC:

```bash
# Клонируйте репозиторий
git clone https://github.com/dedomorozoff/dmmvc
cd dmmvc

# Запустите установочный скрипт
scripts\install.bat
```

Скрипт автоматически:
- Проверит установку Go
- Загрузит зависимости
- Соберет CLI инструмент
- Установит его глобально

### 2. Установка через go install

Если у вас уже установлен Go:

```bash
# Установить последнюю версию
go install github.com/dedomorozoff/dmmvc/cmd/cli@latest

# CLI будет доступен как 'cli'
# Переименуйте в 'dmmvc' если нужно
```

### 3. Установка через Makefile

```bash
# Клонируйте репозиторий
git clone https://github.com/dedomorozoff/dmmvc
cd dmmvc

# Установите зависимости
go mod tidy

# Соберите и установите CLI
make install-go

# Или используйте Windows-специфичную установку
make install
```

### 4. Ручная сборка

```bash
# Клонируйте репозиторий
git clone https://github.com/dedomorozoff/dmmvc
cd dmmvc

# Установите зависимости
go mod download

# Соберите CLI
go build -o dmmvc.exe cmd/cli/main.go

# Соберите сервер
go build -o server.exe cmd/server/main.go

# Добавьте в PATH (опционально)
# Скопируйте dmmvc.exe в %GOPATH%\bin или любую папку в PATH
```

## Проверка установки

После установки проверьте, что CLI работает:

```bash
# Проверить версию Go
go version

# Проверить установку DMMVC CLI
dmmvc --help

# Если команда не найдена, проверьте PATH
echo %PATH%
```

## Настройка PATH

Если команда `dmmvc` не найдена, добавьте Go bin в PATH:

### Временно (текущая сессия)

```cmd
set PATH=%PATH%;%GOPATH%\bin
```

или

```cmd
set PATH=%PATH%;%USERPROFILE%\go\bin
```

### Постоянно

1. Откройте "Система" → "Дополнительные параметры системы"
2. Нажмите "Переменные среды"
3. В "Системные переменные" найдите `Path`
4. Добавьте путь: `%USERPROFILE%\go\bin` или `%GOPATH%\bin`
5. Перезапустите терминал

## Создание первого проекта

### Использование шаблона проекта

```bash
# Создать новый проект
scripts\create-project.bat my-app

# Перейти в проект
cd my-app

# Установить зависимости
go mod tidy

# Запустить сервер
go run cmd/server/main.go
```

### Использование существующего репозитория

```bash
# Клонировать DMMVC как основу
git clone https://github.com/dedomorozoff/dmmvc my-app
cd my-app

# Удалить историю Git (опционально)
rmdir /s /q .git
git init

# Настроить проект
# Отредактируйте go.mod, измените имя модуля
# Отредактируйте .env, настройте параметры

# Установить зависимости
go mod tidy

# Запустить
go run cmd/server/main.go
```

## Использование как библиотеки

Вы можете использовать DMMVC как библиотеку в своем проекте:

```bash
# Создайте новый Go проект
mkdir my-app
cd my-app
go mod init my-app

# Добавьте DMMVC как зависимость
go get github.com/dedomorozoff/dmmvc@latest
```

Затем импортируйте нужные пакеты:

```go
package main

import (
    "github.com/dedomorozoff/dmmvc/internal/database"
    "github.com/dedomorozoff/dmmvc/internal/logger"
    "github.com/dedomorozoff/dmmvc/internal/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    logger.Init()
    database.Init()
    
    r := gin.Default()
    routes.Setup(r)
    
    r.Run(":8080")
}
```

## Docker установка

### Использование готового образа

```bash
# Соберите образ
docker build -t dmmvc .

# Запустите контейнер
docker run -p 8080:8080 dmmvc
```

### Docker Compose

```bash
# Запустите с PostgreSQL
docker-compose -f docker/docker-compose.postgres.yml up
```

## Обновление

### Обновление CLI

```bash
# Через go install
go install github.com/dedomorozoff/dmmvc/cmd/cli@latest

# Или пересоберите локально
cd dmmvc
git pull
make install-go
```

### Обновление зависимостей проекта

```bash
# Обновить все зависимости
go get -u ./...
go mod tidy

# Обновить конкретный пакет
go get -u github.com/gin-gonic/gin
```

## Удаление

```bash
# Удалить CLI
del %GOPATH%\bin\dmmvc.exe

# Или
del %USERPROFILE%\go\bin\dmmvc.exe

# Удалить проект
cd ..
rmdir /s /q dmmvc
```

## Решение проблем

### Go не найден

```bash
# Установите Go с официального сайта
# https://golang.org/dl/

# Проверьте установку
go version
```

### GOPATH не установлен

```bash
# Проверьте GOPATH
echo %GOPATH%

# Если пусто, Go использует значение по умолчанию
echo %USERPROFILE%\go
```

### Ошибка "command not found: dmmvc"

```bash
# Проверьте, что CLI установлен
dir %GOPATH%\bin\dmmvc.exe

# Проверьте PATH
echo %PATH%

# Добавьте в PATH
set PATH=%PATH%;%GOPATH%\bin
```

### Ошибки при сборке

```bash
# Очистите кеш Go
go clean -cache -modcache

# Переустановите зависимости
del go.sum
go mod download
go mod tidy
```

## Дополнительные инструменты

### Рекомендуемые расширения для VS Code

- Go (официальное расширение)
- Go Template Support
- Docker
- GitLens

### Полезные команды

```bash
# Проверить код
go vet ./...

# Форматировать код
go fmt ./...

# Запустить тесты
go test ./...

# Показать зависимости
go mod graph

# Очистить неиспользуемые зависимости
go mod tidy
```

## Следующие шаги

После установки:

1. Прочитайте [Быстрый старт](QUICKSTART.ru.md)
2. Изучите [CLI инструменты](CLI.ru.md)
3. Посмотрите [Примеры](EXAMPLES.ru.md)
4. Настройте [Базу данных](POSTGRESQL.ru.md)

## Поддержка

- **Документация**: [docs/](.)
- **Issues**: https://github.com/dedomorozoff/dmmvc/issues
- **Discussions**: https://github.com/dedomorozoff/dmmvc/discussions
