package routes

import (
	"dmmvc/internal/controllers"
	"dmmvc/internal/handlers"
	"dmmvc/internal/i18n"
	"dmmvc/internal/websocket"
	"dmmvc/internal/middleware"
	"os"
	"path/filepath"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	
	_ "dmmvc/docs/swagger"
)

// loadTemplates загружает все HTML шаблоны
func loadTemplates(r *gin.Engine) {
	// Собираем все HTML файлы из templates и поддиректорий
	templates := []string{}
	
	// Partials
	partials, _ := filepath.Glob("templates/partials/*.html")
	templates = append(templates, partials...)
	
	// Pages
	pages, _ := filepath.Glob("templates/pages/*.html")
	templates = append(templates, pages...)
	
	// Pages/users
	users, _ := filepath.Glob("templates/pages/users/*.html")
	templates = append(templates, users...)
	
	// Загружаем все найденные шаблоны
	if len(templates) > 0 {
		r.LoadHTMLFiles(templates...)
	}
}

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
	r.Use(i18n.Middleware())

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
	// Используем несколько паттернов для поддержки вложенных директорий
	loadTemplates(r)

	// Публичные маршруты
	r.GET("/", controllers.HomePage)
	r.GET("/login", controllers.LoginPage)
	r.POST("/login", controllers.LoginPost)
	r.GET("/logout", controllers.Logout)

	// WebSocket маршрут
	r.GET("/ws", controllers.WebSocketHandler(hub))

	// Swagger документация (доступна без авторизации для удобства разработки)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// I18n API endpoints
	i18nHandler := handlers.NewI18nHandler()
	r.POST("/api/locale", i18nHandler.SetLocale)
	r.GET("/api/locales", i18nHandler.GetLocales)

	// Защищенные маршруты
	authorized := r.Group("/")
	authorized.Use(middleware.AuthRequired())
	authorized.Use(middleware.InjectUserData())
	{
		authorized.GET("/dashboard", controllers.DashboardPage)
		authorized.GET("/profile", controllers.ProfilePage)
		authorized.GET("/websocket", controllers.WebSocketDemo)
		authorized.GET("/i18n", controllers.I18nDemoPage)
		
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

		// API маршруты (примеры)
		api := authorized.Group("/api")
		{
			api.GET("/users", controllers.APIUserList)
			api.GET("/users/:id", controllers.APIUserGet)
			api.POST("/users", controllers.APIUserCreate)
			api.DELETE("/users/:id", controllers.APIUserDelete)

			// Cache примеры
			api.GET("/users/cached", controllers.CachedUserList)
			api.POST("/cache/clear", controllers.ClearUserCache)
			api.GET("/cache/stats", controllers.CacheStats)

			// Queue примеры
			api.POST("/queue/email", controllers.EnqueueEmailTask)
			api.POST("/queue/email/delayed", controllers.EnqueueDelayedTask)
			api.POST("/queue/image", controllers.EnqueueImageTask)
			api.GET("/queue/stats", controllers.QueueStats)

			// Email примеры
			api.POST("/email/send", controllers.SendEmailDirect)
			api.POST("/email/send/async", controllers.SendEmailAsync)
			api.POST("/email/welcome", controllers.SendWelcomeEmail)
			api.POST("/email/password-reset", controllers.SendPasswordResetEmail)
			api.GET("/email/status", controllers.EmailStatus)

			// Upload примеры
			api.POST("/upload/file", controllers.UploadSingleFile)
			api.POST("/upload/files", controllers.UploadMultipleFiles)
			api.POST("/upload/image", controllers.UploadImage)
			api.GET("/upload/file/:filename", controllers.DownloadFile)
			api.DELETE("/upload/file/:filename", controllers.DeleteUploadedFile)
		}

		// Upload страница
		authorized.GET("/upload", controllers.UploadPage)
	}

	return r
}
