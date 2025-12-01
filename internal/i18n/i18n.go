package i18n

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// Locale represents a language locale
type Locale string

const (
	LocaleEN Locale = "en"
	LocaleRU Locale = "ru"
)

// I18n manages translations
type I18n struct {
	translations map[Locale]map[string]string
	defaultLocale Locale
	mu           sync.RWMutex
}

var instance *I18n
var once sync.Once

// GetInstance returns singleton instance of I18n
func GetInstance() *I18n {
	once.Do(func() {
		instance = &I18n{
			translations:  make(map[Locale]map[string]string),
			defaultLocale: LocaleEN,
		}
	})
	return instance
}

// LoadTranslations loads translation files from directory
func (i *I18n) LoadTranslations(dir string) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	locales := []Locale{LocaleEN, LocaleRU}
	
	for _, locale := range locales {
		filename := filepath.Join(dir, fmt.Sprintf("%s.json", locale))
		data, err := os.ReadFile(filename)
		if err != nil {
			return fmt.Errorf("failed to read translation file %s: %w", filename, err)
		}

		var translations map[string]string
		if err := json.Unmarshal(data, &translations); err != nil {
			return fmt.Errorf("failed to parse translation file %s: %w", filename, err)
		}

		i.translations[locale] = translations
	}

	return nil
}

// T translates a key for given locale
func (i *I18n) T(locale Locale, key string, args ...interface{}) string {
	i.mu.RLock()
	defer i.mu.RUnlock()

	// Try to get translation for requested locale
	if translations, ok := i.translations[locale]; ok {
		if translation, ok := translations[key]; ok {
			if len(args) > 0 {
				return fmt.Sprintf(translation, args...)
			}
			return translation
		}
	}

	// Fallback to default locale
	if locale != i.defaultLocale {
		if translations, ok := i.translations[i.defaultLocale]; ok {
			if translation, ok := translations[key]; ok {
				if len(args) > 0 {
					return fmt.Sprintf(translation, args...)
				}
				return translation
			}
		}
	}

	// Return key if no translation found
	return key
}

// SetDefaultLocale sets the default locale
func (i *I18n) SetDefaultLocale(locale Locale) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.defaultLocale = locale
}

// GetAvailableLocales returns list of available locales
func (i *I18n) GetAvailableLocales() []Locale {
	i.mu.RLock()
	defer i.mu.RUnlock()

	locales := make([]Locale, 0, len(i.translations))
	for locale := range i.translations {
		locales = append(locales, locale)
	}
	return locales
}
