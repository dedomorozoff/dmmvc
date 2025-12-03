package controllers

import (
	"dmmvc/internal/i18n"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HomePage отображает главную страницу
func HomePage(c *gin.Context) {
	features, _ := c.Get("features")
	
	data := gin.H{
		"title":    "Home",
		"features": features,
	}
	
	// Add i18n support if enabled
	addI18nData(c, data)
	
	c.HTML(http.StatusOK, "pages/home.html", data)
}

// DashboardPage отображает панель управления
func DashboardPage(c *gin.Context) {
	username, _ := c.Get("username")
	role, _ := c.Get("role")
	features, _ := c.Get("features")

	data := gin.H{
		"title":    "Dashboard",
		"username": username,
		"role":     role,
		"features": features,
	}
	
	// Add i18n support if enabled
	addI18nData(c, data)
	
	c.HTML(http.StatusOK, "pages/dashboard.html", data)
}

// ProfilePage отображает профиль пользователя
func ProfilePage(c *gin.Context) {
	username, _ := c.Get("username")
	role, _ := c.Get("role")
	features, _ := c.Get("features")

	data := gin.H{
		"title":    "Profile",
		"username": username,
		"role":     role,
		"features": features,
	}
	
	// Add i18n support if enabled
	addI18nData(c, data)
	
	c.HTML(http.StatusOK, "pages/profile.html", data)
}

// I18nDemoPage displays i18n demo page
func I18nDemoPage(c *gin.Context) {
	data := gin.H{
		"title": "i18n Demo",
	}
	
	// Add translations using i18n.T helper
	data["app_name"] = i18nT(c, "app.name")
	data["app_description"] = i18nT(c, "app.description")
	data["nav_home"] = i18nT(c, "nav.home")
	data["nav_dashboard"] = i18nT(c, "nav.dashboard")
	data["nav_profile"] = i18nT(c, "nav.profile")
	data["nav_upload"] = i18nT(c, "nav.upload")
	data["nav_websocket"] = i18nT(c, "nav.websocket")
	data["nav_logout"] = i18nT(c, "nav.logout")
	data["common_submit"] = i18nT(c, "common.submit")
	data["common_cancel"] = i18nT(c, "common.cancel")
	data["common_save"] = i18nT(c, "common.save")
	data["common_delete"] = i18nT(c, "common.delete")
	data["locale"] = i18nLocale(c)
	
	c.HTML(http.StatusOK, "pages/i18n_demo.html", data)
}

// Helper functions for i18n
func i18nT(c *gin.Context, key string, args ...interface{}) string {
	return i18n.T(c, key, args...)
}

func i18nLocale(c *gin.Context) string {
	return string(i18n.GetLocale(c))
}

// addI18nData adds i18n support to template data if i18n is enabled
func addI18nData(c *gin.Context, data gin.H) {
	featuresVal, exists := c.Get("features")
	if !exists {
		return
	}
	
	// Check if i18n is enabled via context (set by middleware)
	if i18nEnabled, exists := c.Get("i18n_enabled"); exists && i18nEnabled.(bool) {
		locale := i18n.GetLocale(c)
		data["locale"] = i18nLocale(c)
		data["T"] = i18n.TFunc(locale)
	} else {
		// Fallback: provide dummy function that returns the key
		data["locale"] = "en"
		data["T"] = func(key string, args ...interface{}) string {
			return key
		}
	}
	
	_ = featuresVal // avoid unused variable warning
}
