# Language Switching Guide / Руководство по переключению языков

## Quick Start / Быстрый старт

DMMVC supports bilingual interface (English and Russian) out of the box.

DMMVC поддерживает двуязычный интерфейс (английский и русский) из коробки.

## How to Switch Language / Как переключить язык

### Method 1: Language Switcher in Header / Метод 1: Переключатель в шапке

Click on the language buttons in the header:

Нажмите на кнопки языка в шапке:

- 🇬🇧 EN - Switch to English / Переключить на английский
- 🇷🇺 RU - Switch to Russian / Переключить на русский

The active language is highlighted. The switcher is available on all pages.

Активный язык подсвечен. Переключатель доступен на всех страницах.

### Method 2: URL Parameter / Метод 2: Параметр URL

Add `?lang=en` or `?lang=ru` to any URL:

Добавьте `?lang=en` или `?lang=ru` к любому URL:

```
http://localhost:8080/?lang=en    # English
http://localhost:8080/?lang=ru    # Russian
```

### Method 3: Browser Settings / Метод 3: Настройки браузера

The framework automatically detects your browser's language preference from the `Accept-Language` header.

Фреймворк автоматически определяет предпочитаемый язык браузера из заголовка `Accept-Language`.

### Method 4: API Requests / Метод 4: API запросы

For API calls, use the `Accept-Language` header:

Для API вызовов используйте заголовок `Accept-Language`:

```bash
# English
curl -H "Accept-Language: en" http://localhost:8080/api/users

# Russian
curl -H "Accept-Language: ru" http://localhost:8080/api/users
```

Or use the `lang` query parameter:

Или используйте параметр запроса `lang`:

```bash
curl http://localhost:8080/api/users?lang=ru
```

## What Gets Translated / Что переводится

### Web Interface / Веб-интерфейс

- ✅ Home page (welcome message, features, documentation links)
- ✅ Navigation menu
- ✅ Login/Register forms
- ✅ Dashboard
- ✅ Profile page
- ✅ Upload page
- ✅ WebSocket demo
- ✅ All buttons and labels

### API Responses / API ответы

- ✅ Success messages
- ✅ Error messages
- ✅ Validation messages
- ✅ Status messages

### Demo Data / Демо данные

- ✅ Demo user names (English: John Doe, Jane Smith / Russian: Иван Иванов, Мария Петрова)
- ✅ Email templates
- ✅ Sample messages

## Default Language / Язык по умолчанию

The default language is **English** (`en`).

Язык по умолчанию - **английский** (`en`).

To change the default, modify `.env`:

Чтобы изменить язык по умолчанию, измените `.env`:

```env
DEFAULT_LOCALE=ru
```

## Supported Languages / Поддерживаемые языки

Currently supported:

В настоящее время поддерживаются:

- 🇬🇧 English (`en`)
- 🇷🇺 Russian (`ru`)

## Adding More Languages / Добавление новых языков

To add a new language (e.g., Spanish):

Чтобы добавить новый язык (например, испанский):

1. Create `locales/es.json`:

```json
{
  "app.name": "DMMVC",
  "home.welcome": "Bienvenido a DMMVC",
  ...
}
```

2. Copy all keys from `locales/en.json` and translate them.

Скопируйте все ключи из `locales/en.json` и переведите их.

3. Restart the server / Перезапустите сервер:

```bash
go run cmd/server/main.go
```

4. Use the new language / Используйте новый язык:

```
http://localhost:8080/?lang=es
```

## Testing / Тестирование

### Test in Browser / Тест в браузере

1. Open `http://localhost:8080/?lang=en` - should see English
2. Open `http://localhost:8080/?lang=ru` - should see Russian

### Test API / Тест API

```bash
# Test English API response
curl -H "Accept-Language: en" http://localhost:8080/api/users | jq

# Test Russian API response
curl -H "Accept-Language: ru" http://localhost:8080/api/users | jq
```

## Troubleshooting / Устранение неполадок

### Language not switching / Язык не переключается

1. Clear browser cache / Очистите кэш браузера
2. Check that locale files exist in `locales/` directory
3. Restart the server / Перезапустите сервер

### Missing translations / Отсутствующие переводы

If you see a translation key instead of text (e.g., `home.welcome`):

Если вы видите ключ перевода вместо текста (например, `home.welcome`):

1. Check that the key exists in `locales/en.json` and `locales/ru.json`
2. Verify JSON syntax is correct (no trailing commas)
3. Restart the server / Перезапустите сервер

## Related Documentation / Связанная документация

- [Internationalization (i18n)](I18N.md) - Full i18n documentation
- [Demo Data](DEMO_DATA.md) - Bilingual demo users and data
- [API Documentation](SWAGGER.md) - API endpoints

---

For more information, see [I18N.md](I18N.md) and [I18N.ru.md](I18N.ru.md).

Для получения дополнительной информации см. [I18N.md](I18N.md) и [I18N.ru.md](I18N.ru.md).
