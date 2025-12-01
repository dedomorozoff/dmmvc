[English](QUEUE.md) | **Русский**

# Очереди задач

DMMVC включает встроенную поддержку фоновой обработки задач с использованием очередей на базе Redis.

## Возможности

- **Асинхронная обработка** - Выполнение задач в фоновом режиме
- **Отложенные задачи** - Планирование задач на будущее
- **Приоритетные очереди** - Критические, обычные и низкоприоритетные очереди
- **Механизм повторов** - Автоматический повтор при ошибке
- **Мониторинг задач** - Отслеживание статуса и прогресса задач

## Быстрый старт

### 1. Установка Redis

Redis необходим для работы очередей задач.

**Docker:**
```bash
docker run -d -p 6379:6379 redis:alpine
```

**См. [CACHE.ru.md](CACHE.ru.md) для других методов установки**

### 2. Настройка очереди задач

Отредактируйте файл `.env`:

```env
# Redis подключение
REDIS_URL=localhost:6379
REDIS_PASSWORD=

# Включить worker для обработки задач
QUEUE_WORKER_ENABLED=true
QUEUE_CONCURRENCY=10
```

### 3. Запуск сервера

```bash
make run
```

Worker автоматически запустится и начнет обработку задач.

## Использование

### Создание задач

```go
import (
    "dmmvc/internal/queue"
    "time"
)

// Создать задачу отправки email
task, err := queue.NewEmailDeliveryTask(
    "user@example.com",
    "Добро пожаловать!",
    "Добро пожаловать в наш сервис!",
)

// Добавить в очередь немедленно
queue.EnqueueTask(task)

// Добавить с задержкой
queue.EnqueueTaskIn(task, 5*time.Minute)

// Добавить на определенное время
queue.EnqueueTaskAt(task, time.Now().Add(1*time.Hour))
```

### Приоритетные очереди

```go
import "github.com/hibiken/asynq"

// Критический приоритет
queue.EnqueueTask(task, asynq.Queue("critical"))

// Обычный приоритет
queue.EnqueueTask(task, asynq.Queue("default"))

// Низкий приоритет
queue.EnqueueTask(task, asynq.Queue("low"))
```

### Опции задач

```go
// Максимальное количество повторов
queue.EnqueueTask(task, asynq.MaxRetry(3))

// Таймаут задачи
queue.EnqueueTask(task, asynq.Timeout(5*time.Minute))

// Уникальная задача (предотвращение дубликатов)
queue.EnqueueTask(task, asynq.Unique(24*time.Hour))

// Комбинация опций
queue.EnqueueTask(task,
    asynq.Queue("critical"),
    asynq.MaxRetry(5),
    asynq.Timeout(10*time.Minute),
)
```

## Встроенные задачи

### Отправка Email

```go
task, _ := queue.NewEmailDeliveryTask(
    "user@example.com",
    "Тема",
    "Текст письма",
)
queue.EnqueueTask(task)
```

### Изменение размера изображения

```go
task, _ := queue.NewImageResizeTask(
    "/uploads/image.jpg",
    "/uploads/thumb.jpg",
    300,
    200,
)
queue.EnqueueTask(task)
```

## Создание пользовательских задач

### 1. Определение типа задачи

```go
// internal/queue/tasks.go
const (
    TypeDataExport = "data:export"
)

type DataExportPayload struct {
    UserID int    `json:"user_id"`
    Format string `json:"format"`
}

func NewDataExportTask(userID int, format string) (*asynq.Task, error) {
    payload, err := json.Marshal(DataExportPayload{
        UserID: userID,
        Format: format,
    })
    if err != nil {
        return nil, err
    }
    return asynq.NewTask(TypeDataExport, payload), nil
}
```

### 2. Создание обработчика

```go
func HandleDataExportTask(ctx context.Context, t *asynq.Task) error {
    var p DataExportPayload
    if err := json.Unmarshal(t.Payload(), &p); err != nil {
        return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
    }

    logrus.Infof("Экспорт данных для пользователя %d в формате %s", p.UserID, p.Format)
    
    // Ваша логика экспорта здесь
    
    return nil
}
```

### 3. Регистрация обработчика

```go
// internal/queue/worker.go
mux := asynq.NewServeMux()
mux.HandleFunc(TypeEmailDelivery, HandleEmailDeliveryTask)
mux.HandleFunc(TypeImageResize, HandleImageResizeTask)
mux.HandleFunc(TypeDataExport, HandleDataExportTask) // Добавить это
```

## API эндпоинты

DMMVC включает примеры queue эндпоинтов:

- `POST /api/queue/email` - Добавить задачу отправки email
- `POST /api/queue/email/delayed` - Добавить отложенную задачу email
- `POST /api/queue/image` - Добавить задачу обработки изображения
- `GET /api/queue/stats` - Статистика очереди

### Пример запроса

```bash
curl -X POST http://localhost:8080/api/queue/email \
  -H "Content-Type: application/json" \
  -d '{
    "to": "user@example.com",
    "subject": "Привет",
    "body": "Добро пожаловать!"
  }'
```

## Пример контроллера

```go
func SendWelcomeEmail(c *gin.Context) {
    userEmail := c.PostForm("email")
    
    // Создать задачу
    task, err := queue.NewEmailDeliveryTask(
        userEmail,
        "Добро пожаловать!",
        "Спасибо за регистрацию!",
    )
    if err != nil {
        c.JSON(500, gin.H{"error": "Не удалось создать задачу"})
        return
    }
    
    // Добавить в очередь
    if err := queue.EnqueueTask(task); err != nil {
        c.JSON(500, gin.H{"error": "Не удалось добавить задачу"})
        return
    }
    
    c.JSON(200, gin.H{"message": "Email будет отправлен в ближайшее время"})
}
```

## Периодические задачи

Планирование периодического выполнения задач:

```go
// internal/queue/scheduler.go
import "github.com/hibiken/asynq"

func SetupScheduler() *asynq.Scheduler {
    scheduler := asynq.NewScheduler(
        asynq.RedisClientOpt{Addr: "localhost:6379"},
        nil,
    )
    
    // Запуск очистки каждый день в полночь
    task, _ := asynq.NewTask(TypeCleanup, nil)
    scheduler.Register("@daily", task)
    
    // Запуск каждый час
    scheduler.Register("@hourly", task)
    
    // Пользовательское cron выражение
    scheduler.Register("*/30 * * * *", task) // Каждые 30 минут
    
    return scheduler
}
```

## Мониторинг

### Asynq Web UI

Установка и запуск веб-интерфейса:

```bash
go install github.com/hibiken/asynq/tools/asynq@latest
asynq dash
```

Доступ по адресу: http://localhost:8080

### Метрики

Отслеживание метрик задач:

```go
// Получить информацию об очереди
inspector := asynq.NewInspector(asynq.RedisClientOpt{
    Addr: "localhost:6379",
})

// Получить ожидающие задачи
pending, _ := inspector.ListPendingTasks("default")

// Получить активные задачи
active, _ := inspector.ListActiveTasks("default")

// Получить завершенные задачи
completed, _ := inspector.ListCompletedTasks("default")
```

## Лучшие практики

1. **Делайте задачи идемпотентными** - Задачи должны быть безопасны для повтора
2. **Используйте подходящие таймауты** - Предотвращайте слишком долгое выполнение
3. **Обрабатывайте ошибки корректно** - Возвращайте правильные типы ошибок
4. **Используйте уникальные задачи** - Предотвращайте дублирование обработки
5. **Мониторьте глубину очереди** - Следите за накоплением задач
6. **Устанавливайте лимиты повторов** - Не повторяйте бесконечно
7. **Используйте приоритетные очереди** - Обрабатывайте критические задачи первыми

## Обработка ошибок

```go
func HandleTask(ctx context.Context, t *asynq.Task) error {
    // Пропустить повтор для определенных ошибок
    if err := validatePayload(t.Payload()); err != nil {
        return fmt.Errorf("invalid payload: %w", asynq.SkipRetry)
    }
    
    // Повтор при временных ошибках
    if err := processTask(t); err != nil {
        return err // Будет повтор
    }
    
    return nil
}
```

## Production конфигурация

```env
# Production настройки
REDIS_URL=redis.example.com:6379
REDIS_PASSWORD=your-secure-password
QUEUE_WORKER_ENABLED=true
QUEUE_CONCURRENCY=20

# Используйте несколько workers для высокой нагрузки
# Запустите несколько инстансов с разными очередями
```

## Docker Compose

```yaml
version: '3.8'
services:
  app:
    build: .
    ports:
      - "8080:8080"
    environment:
      - REDIS_URL=redis:6379
      - QUEUE_WORKER_ENABLED=false
    depends_on:
      - redis

  worker:
    build: .
    environment:
      - REDIS_URL=redis:6379
      - QUEUE_WORKER_ENABLED=true
      - QUEUE_CONCURRENCY=10
    depends_on:
      - redis

  redis:
    image: redis:alpine
    ports:
      - "6379:6379"
    volumes:
      - redis_data:/data

volumes:
  redis_data:
```

## Решение проблем

### Задачи не обрабатываются

1. Проверьте, включен ли worker: `QUEUE_WORKER_ENABLED=true`
2. Проверьте подключение к Redis
3. Проверьте логи worker на ошибки

### Задачи падают с ошибкой

1. Проверьте реализацию обработчика задачи
2. Просмотрите логи ошибок
3. Проверьте структуру payload

### Накопление очереди

1. Увеличьте concurrency
2. Добавьте больше workers
3. Оптимизируйте обработчики задач
4. Проверьте наличие падающих задач

## Ресурсы

- [Документация Asynq](https://github.com/hibiken/asynq)
- [Документация Redis](https://redis.io/documentation)
- [Руководство по Cron выражениям](https://crontab.guru/)

## Следующие шаги

- Реализовать отправку email
- Добавить обработку изображений
- Создать задачи экспорта данных
- Настроить мониторинг dashboard
- Настроить оповещения
