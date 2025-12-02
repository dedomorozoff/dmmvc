package i18n

import (
	"html/template"
)

// TemplateFuncs returns template functions for i18n
func TemplateFuncs(locale Locale) template.FuncMap {
	return template.FuncMap{
		"t": func(key string, args ...interface{}) string {
			return GetInstance().T(locale, key, args...)
		},
		"locale": func() string {
			return string(locale)
		},
	}
}

// GetTemplateFuncs returns template functions that work with gin context
// Note: These functions expect the locale to be passed in the template data as ".locale"
func GetTemplateFuncs() template.FuncMap {
	return template.FuncMap{
		"T": func(key string, args ...interface{}) string {
			// This will be called from template with current locale from context
			// We need to get locale from the template data, not from gin context
			// For now, use default locale - will be overridden by context-aware version
			return GetInstance().T(LocaleEN, key, args...)
		},
	}
}

// TFunc returns a translation function for a specific locale
func TFunc(locale Locale) func(string, ...interface{}) string {
	return func(key string, args ...interface{}) string {
		return GetInstance().T(locale, key, args...)
	}
}
