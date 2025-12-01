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
