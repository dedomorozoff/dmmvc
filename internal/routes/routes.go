package routes

import (
	"dmmvc/internal/controllers"
	"dmmvc/internal/websocket"
	"dmmvc/internal/middleware"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// SetupRouter настраивает все маршруты приложения
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Инициализация WebSocket Hub
	hub := websocket.NewHub()
	go hub.Run()

	// Настройка доверенных прокси
	r.SetTrustedProxies([]string{"127.0.0.1", "::1"})

	// Middleware
	r.Use(middleware.RequestLogger())

	// Настройка сессий
	secret := os.Getenv("SESSION_SECRET")
	if secret == "" {
		secret = "secret"
	}
	store := cookie.NewStore([]byte(secret))
	r.Use(sessions.Sessions("dmmvc_session", store))

	// Статические файлы
	r.Static("/static", "./static")

	// Загрузка шаблонов
	r.LoadHTMLGlob("templates/**/*.html")

	// Публичные маршруты
	r.GET("/", controllers.HomePage)
	r.GET("/login", controllers.LoginPage)
	r.POST("/login", controllers.LoginPost)
	r.GET("/logout", controllers.Logout)

	// WebSocket маршрут
	r.GET("/ws", controllers.WebSocketHandler(hub))

	// Защищенные маршруты
	authorized := r.Group("/")
	authorized.Use(middleware.AuthRequired())
	authorized.Use(middleware.InjectUserData())
	{
		authorized.GET("/dashboard", controllers.DashboardPage)
		authorized.GET("/profile", controllers.ProfilePage)
		authorized.GET("/websocket", controllers.WebSocketDemo)
		
		// Пример CRUD маршрутов для пользователей (только для админа)
		admin := authorized.Group("/admin")
		{
			admin.GET("/users", controllers.UserList)
			admin.GET("/users/create", controllers.UserCreate)
			admin.POST("/users", controllers.UserStore)
			admin.GET("/users/:id/edit", controllers.UserEdit)
			admin.POST("/users/:id/update", controllers.UserUpdate)
			admin.POST("/users/:id/delete", controllers.UserDelete)
		}
	}

	return r
}
