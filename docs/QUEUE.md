**English** | [Русский](QUEUE.ru.md)

# Task Queue

DMMVC includes built-in support for background task processing using Redis-backed queues.

## Features

- **Asynchronous processing** - Execute tasks in the background
- **Delayed tasks** - Schedule tasks for future execution
- **Priority queues** - Critical, default, and low priority queues
- **Retry mechanism** - Automatic retry on failure
- **Task monitoring** - Track task status and progress

## Quick Start

### 1. Install Redis

Redis is required for task queue functionality.

**Docker:**
```bash
docker run -d -p 6379:6379 redis:alpine
```

**See [CACHE.md](CACHE.md) for other installation methods**

### 2. Configure Task Queue

Edit `.env` file:

```env
# Redis connection
REDIS_URL=localhost:6379
REDIS_PASSWORD=

# Enable worker to process tasks
QUEUE_WORKER_ENABLED=true
QUEUE_CONCURRENCY=10
```

### 3. Start Server

```bash
make run
```

The worker will automatically start and begin processing tasks.

## Usage

### Creating Tasks

```go
import (
    "dmmvc/internal/queue"
    "time"
)

// Create email task
task, err := queue.NewEmailDeliveryTask(
    "user@example.com",
    "Welcome!",
    "Welcome to our service!",
)

// Enqueue immediately
queue.EnqueueTask(task)

// Enqueue with delay
queue.EnqueueTaskIn(task, 5*time.Minute)

// Enqueue at specific time
queue.EnqueueTaskAt(task, time.Now().Add(1*time.Hour))
```

### Priority Queues

```go
import "github.com/hibiken/asynq"

// Critical priority
queue.EnqueueTask(task, asynq.Queue("critical"))

// Default priority
queue.EnqueueTask(task, asynq.Queue("default"))

// Low priority
queue.EnqueueTask(task, asynq.Queue("low"))
```

### Task Options

```go
// Max retry attempts
queue.EnqueueTask(task, asynq.MaxRetry(3))

// Task timeout
queue.EnqueueTask(task, asynq.Timeout(5*time.Minute))

// Unique task (prevent duplicates)
queue.EnqueueTask(task, asynq.Unique(24*time.Hour))

// Combine options
queue.EnqueueTask(task,
    asynq.Queue("critical"),
    asynq.MaxRetry(5),
    asynq.Timeout(10*time.Minute),
)
```

## Built-in Tasks

### Email Delivery

```go
task, _ := queue.NewEmailDeliveryTask(
    "user@example.com",
    "Subject",
    "Email body",
)
queue.EnqueueTask(task)
```

### Image Resize

```go
task, _ := queue.NewImageResizeTask(
    "/uploads/image.jpg",
    "/uploads/thumb.jpg",
    300,
    200,
)
queue.EnqueueTask(task)
```

## Creating Custom Tasks

### 1. Define Task Type

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

### 2. Create Handler

```go
func HandleDataExportTask(ctx context.Context, t *asynq.Task) error {
    var p DataExportPayload
    if err := json.Unmarshal(t.Payload(), &p); err != nil {
        return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
    }

    logrus.Infof("Exporting data for user %d in %s format", p.UserID, p.Format)
    
    // Your export logic here
    
    return nil
}
```

### 3. Register Handler

```go
// internal/queue/worker.go
mux := asynq.NewServeMux()
mux.HandleFunc(TypeEmailDelivery, HandleEmailDeliveryTask)
mux.HandleFunc(TypeImageResize, HandleImageResizeTask)
mux.HandleFunc(TypeDataExport, HandleDataExportTask) // Add this
```

## API Endpoints

DMMVC includes example queue endpoints:

- `POST /api/queue/email` - Enqueue email task
- `POST /api/queue/email/delayed` - Enqueue delayed email task
- `POST /api/queue/image` - Enqueue image processing task
- `GET /api/queue/stats` - Queue statistics

### Example Request

```bash
curl -X POST http://localhost:8080/api/queue/email \
  -H "Content-Type: application/json" \
  -d '{
    "to": "user@example.com",
    "subject": "Hello",
    "body": "Welcome!"
  }'
```

## Controller Example

```go
func SendWelcomeEmail(c *gin.Context) {
    userEmail := c.PostForm("email")
    
    // Create task
    task, err := queue.NewEmailDeliveryTask(
        userEmail,
        "Welcome!",
        "Thank you for registering!",
    )
    if err != nil {
        c.JSON(500, gin.H{"error": "Failed to create task"})
        return
    }
    
    // Enqueue task
    if err := queue.EnqueueTask(task); err != nil {
        c.JSON(500, gin.H{"error": "Failed to enqueue task"})
        return
    }
    
    c.JSON(200, gin.H{"message": "Email will be sent shortly"})
}
```

## Periodic Tasks

Schedule tasks to run periodically:

```go
// internal/queue/scheduler.go
import "github.com/hibiken/asynq"

func SetupScheduler() *asynq.Scheduler {
    scheduler := asynq.NewScheduler(
        asynq.RedisClientOpt{Addr: "localhost:6379"},
        nil,
    )
    
    // Run cleanup every day at midnight
    task, _ := asynq.NewTask(TypeCleanup, nil)
    scheduler.Register("@daily", task)
    
    // Run every hour
    scheduler.Register("@hourly", task)
    
    // Custom cron expression
    scheduler.Register("*/30 * * * *", task) // Every 30 minutes
    
    return scheduler
}
```

## Monitoring

### Asynq Web UI

Install and run the web UI:

```bash
go install github.com/hibiken/asynq/tools/asynq@latest
asynq dash
```

Access at: http://localhost:8080

### Metrics

Track task metrics:

```go
// Get queue info
inspector := asynq.NewInspector(asynq.RedisClientOpt{
    Addr: "localhost:6379",
})

// Get pending tasks
pending, _ := inspector.ListPendingTasks("default")

// Get active tasks
active, _ := inspector.ListActiveTasks("default")

// Get completed tasks
completed, _ := inspector.ListCompletedTasks("default")
```

## Best Practices

1. **Keep tasks idempotent** - Tasks should be safe to retry
2. **Use appropriate timeouts** - Prevent tasks from running too long
3. **Handle errors gracefully** - Return proper error types
4. **Use unique tasks** - Prevent duplicate processing
5. **Monitor queue depth** - Watch for queue buildup
6. **Set retry limits** - Don't retry forever
7. **Use priority queues** - Process critical tasks first

## Error Handling

```go
func HandleTask(ctx context.Context, t *asynq.Task) error {
    // Skip retry for certain errors
    if err := validatePayload(t.Payload()); err != nil {
        return fmt.Errorf("invalid payload: %w", asynq.SkipRetry)
    }
    
    // Retry on temporary errors
    if err := processTask(t); err != nil {
        return err // Will retry
    }
    
    return nil
}
```

## Production Configuration

```env
# Production settings
REDIS_URL=redis.example.com:6379
REDIS_PASSWORD=your-secure-password
QUEUE_WORKER_ENABLED=true
QUEUE_CONCURRENCY=20

# Use multiple workers for high load
# Start multiple instances with different queues
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

## Troubleshooting

### Tasks not processing

1. Check if worker is enabled: `QUEUE_WORKER_ENABLED=true`
2. Verify Redis connection
3. Check worker logs for errors

### Tasks failing

1. Check task handler implementation
2. Review error logs
3. Verify payload structure

### Queue buildup

1. Increase concurrency
2. Add more workers
3. Optimize task handlers
4. Check for failing tasks

## Resources

- [Asynq Documentation](https://github.com/hibiken/asynq)
- [Redis Documentation](https://redis.io/documentation)
- [Cron Expression Guide](https://crontab.guru/)

## Next Steps

- Implement email sending
- Add image processing
- Create data export tasks
- Set up monitoring dashboard
- Configure alerting
