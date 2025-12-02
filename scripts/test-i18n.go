package main

import (
	"fmt"
	"os"
)

func main() {
	// Проверяем переменную окружения
	defaultLocale := os.Getenv("DEFAULT_LOCALE")
	if defaultLocale == "" {
		defaultLocale = "en (not set)"
	}
	
	fmt.Println("=== i18n Configuration Test ===")
	fmt.Printf("DEFAULT_LOCALE: %s\n", defaultLocale)
	fmt.Println("\nTo test:")
	fmt.Println("1. Set DEFAULT_LOCALE=ru in .env")
	fmt.Println("2. Restart the server")
	fmt.Println("3. Visit http://localhost:8080/i18n")
	fmt.Println("4. Check that default language is Russian")
	fmt.Println("\nOr test via API:")
	fmt.Println("curl http://localhost:8080/api/locales")
}
