package i18n

import (
	"github.com/gin-gonic/gin"
)

const (
	LocaleContextKey = "locale"
	LocaleCookieName = "locale"
)

// Middleware adds locale detection to gin context
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		locale := detectLocale(c)
		c.Set(LocaleContextKey, locale)
		c.Next()
	}
}

// detectLocale detects user's preferred locale
func detectLocale(c *gin.Context) Locale {
	// 1. Check query parameter
	if lang := c.Query("lang"); lang != "" {
		if locale := parseLocale(lang); locale != "" {
			// Set cookie for future requests
			c.SetCookie(LocaleCookieName, string(locale), 365*24*60*60, "/", "", false, false)
			return locale
		}
	}

	// 2. Check cookie
	if lang, err := c.Cookie(LocaleCookieName); err == nil {
		if locale := parseLocale(lang); locale != "" {
			return locale
		}
	}

	// 3. Check Accept-Language header
	if lang := c.GetHeader("Accept-Language"); lang != "" {
		if locale := parseLocale(lang); locale != "" {
			return locale
		}
	}

	// 4. Default to English
	return LocaleEN
}

// parseLocale parses locale string
func parseLocale(lang string) Locale {
	// Simple parsing - take first 2 characters
	if len(lang) >= 2 {
		switch lang[:2] {
		case "en":
			return LocaleEN
		case "ru":
			return LocaleRU
		}
	}
	return ""
}

// GetLocale returns locale from gin context
func GetLocale(c *gin.Context) Locale {
	if locale, exists := c.Get(LocaleContextKey); exists {
		if l, ok := locale.(Locale); ok {
			return l
		}
	}
	return LocaleEN
}

// T is a helper function to translate in gin handlers
func T(c *gin.Context, key string, args ...interface{}) string {
	locale := GetLocale(c)
	return GetInstance().T(locale, key, args...)
}

// GetTranslations returns a map of common translations for templates
func GetTranslations(c *gin.Context) map[string]string {
	locale := GetLocale(c)
	i18n := GetInstance()
	
	// Return commonly used translations
	return map[string]string{
		"app_name":        i18n.T(locale, "app.name"),
		"app_description": i18n.T(locale, "app.description"),
		"nav_home":        i18n.T(locale, "nav.home"),
		"nav_dashboard":   i18n.T(locale, "nav.dashboard"),
		"nav_profile":     i18n.T(locale, "nav.profile"),
		"nav_upload":      i18n.T(locale, "nav.upload"),
		"nav_websocket":   i18n.T(locale, "nav.websocket"),
		"nav_logout":      i18n.T(locale, "nav.logout"),
		"common_submit":   i18n.T(locale, "common.submit"),
		"common_cancel":   i18n.T(locale, "common.cancel"),
		"common_save":     i18n.T(locale, "common.save"),
		"common_delete":   i18n.T(locale, "common.delete"),
		"locale":          string(locale),
	}
}
