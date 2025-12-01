[English](QUICKSTART_CLI.md) | **Русский** | [Docs](README.md)

# Быстрый старт с DMMVC CLI

Начните работу с DMMVC CLI за 5 минут!

## Шаг 1: Соберите CLI

```bash
make build
```

Это создаст `dmmvc.exe` в корне проекта.

## Шаг 2: Создайте первый CRUD

Давайте создадим простую систему управления постами блога:

```bash
dmmvc make:crud Post
```

Это сгенерирует:
- Модель: `internal/models/post.go`
- Контроллер: `internal/controllers/post_controller.go` (с 7 CRUD методами)
- Шаблоны: `templates/pages/post/*.html` (index, show, create, edit)

## Шаг 3: Настройте модель

Отредактируйте `internal/models/post.go`:

```go
type Post struct {
    gorm.Model
    Title   string `json:"title" gorm:"not null"`
    Content string `json:"content" gorm:"type:text"`
    Author  string `json:"author"`
}
```

## Шаг 4: Добавьте миграцию

Отредактируйте `internal/database/db.go`, добавьте в `InitDB()`:

```go
db.AutoMigrate(&models.Post{})
```

## Шаг 5: Добавьте маршруты

Отредактируйте `internal/routes/routes.go`, добавьте эти маршруты:

```go
// Post routes
authorized.GET("/post", controllers.PostControllerIndex)
authorized.GET("/post/:id", controllers.PostControllerShow)
authorized.GET("/post/create", controllers.PostControllerCreate)
authorized.POST("/post", controllers.PostControllerStore)
authorized.GET("/post/:id/edit", controllers.PostControllerEdit)
authorized.POST("/post/:id", controllers.PostControllerUpdate)
authorized.POST("/post/:id/delete", controllers.PostControllerDelete)
```

## Шаг 6: Настройте шаблоны

Отредактируйте шаблоны в `templates/pages/post/`:

### `create.html` - Добавьте поля формы:

```html
<form method="POST" action="/post">
    <div class="form-group">
        <label>Заголовок</label>
        <input type="text" name="title" required>
    </div>
    <div class="form-group">
        <label>Содержание</label>
        <textarea name="content" rows="10" required></textarea>
    </div>
    <div class="form-group">
        <label>Автор</label>
        <input type="text" name="author" required>
    </div>
    <button type="submit">Создать пост</button>
</form>
```

### `index.html` - Обновите колонки таблицы:

```html
<thead>
    <tr>
        <th>ID</th>
        <th>Заголовок</th>
        <th>Автор</th>
        <th>Создан</th>
        <th>Действия</th>
    </tr>
</thead>
<tbody>
    {{range .items}}
    <tr>
        <td>{{.ID}}</td>
        <td>{{.Title}}</td>
        <td>{{.Author}}</td>
        <td>{{.CreatedAt}}</td>
        <td>
            <a href="/post/{{.ID}}">Просмотр</a>
            <a href="/post/{{.ID}}/edit">Редактировать</a>
            <form method="POST" action="/post/{{.ID}}/delete" style="display:inline;">
                <button type="submit">Удалить</button>
            </form>
        </td>
    </tr>
    {{end}}
</tbody>
```

## Шаг 7: Запустите и протестируйте

```bash
go run cmd/server/main.go
```

Откройте: `http://localhost:8080/post`

## Больше команд CLI

```bash
# Создать простой контроллер
dmmvc make:controller About

# Создать модель с подсказкой по миграции
dmmvc make:model Category --migration

# Создать middleware
dmmvc make:middleware RateLimit

# Создать шаблон страницы
dmmvc make:page contact

# Показать все ресурсы
dmmvc list

# Показать справку
dmmvc --help
```

## Советы

- **Используйте `make:crud`** для быстрого создания каркаса, затем настраивайте
- **Используйте флаг `--resource`** для контроллеров с CRUD методами
- **Используйте команду `list`** чтобы увидеть все ваши ресурсы
- **Всегда проверяйте** сгенерированный код и адаптируйте под свои нужды

## Следующие шаги

- Прочитайте полную документацию: [CLI.ru.md](CLI.ru.md)
- Посмотрите примеры: [EXAMPLES.ru.md](EXAMPLES.ru.md)
- Узнайте об архитектуре: [ARCHITECTURE.ru.md](ARCHITECTURE.ru.md)
- Вернуться к [Оглавлению](README.md)

Удачного кодирования!
