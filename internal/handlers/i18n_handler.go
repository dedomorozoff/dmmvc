package handlers

import (
	"dmmvc/internal/i18n"
	"net/http"

	"github.com/gin-gonic/gin"
)

type I18nHandler struct{}

func NewI18nHandler() *I18nHandler {
	return &I18nHandler{}
}

// SetLocale handles locale change requests
func (h *I18nHandler) SetLocale(c *gin.Context) {
	var req struct {
		Locale string `json:"locale" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid locale",
		})
		return
	}

	locale := i18n.Locale(req.Locale)
	
	// Validate locale
	validLocales := map[i18n.Locale]bool{
		i18n.LocaleEN: true,
		i18n.LocaleRU: true,
	}

	if !validLocales[locale] {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Unsupported locale",
		})
		return
	}

	// Set cookie
	c.SetCookie(i18n.LocaleCookieName, string(locale), 365*24*60*60, "/", "", false, false)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"locale":  locale,
	})
}

// GetLocales returns available locales
func (h *I18nHandler) GetLocales(c *gin.Context) {
	locales := i18n.GetInstance().GetAvailableLocales()
	currentLocale := i18n.GetLocale(c)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"locales": locales,
			"current": currentLocale,
		},
	})
}
