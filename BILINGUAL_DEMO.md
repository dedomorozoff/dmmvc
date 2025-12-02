# Bilingual Demo Data Implementation / Реализация двуязычных демо данных

## Summary / Краткое описание

DMMVC now includes full bilingual support (English and Russian) for all demo data, UI elements, and API responses.

DMMVC теперь включает полную двуязычную поддержку (английский и русский) для всех демо данных, элементов UI и ответов API.

## What's New / Что нового

### 0. Language Switcher in Header / Переключатель языков в шапке

**Visual Language Switcher:**
- ✅ 🇬🇧 EN / 🇷🇺 RU buttons in header
- ✅ Active language highlighted
- ✅ Available on all pages
- ✅ Preserves current URL when switching
- ✅ Smooth transitions and hover effects

Click the language button in the header to instantly switch between English and Russian!

Нажмите на кнопку языка в шапке для мгновенного переключения между английским и русским!

### 1. Bilingual Demo Users / Двуязычные демо пользователи

**English Users:**
- john_doe (john@example.com)
- jane_smith (jane@example.com)
- bob_johnson (bob@example.com)

**Russian Users:**
- ivan_ivanov (ivan@example.ru)
- maria_petrova (maria@example.ru)
- alexey_sidorov (alexey@example.ru)

All demo users have password: `password123`

### 2. Translated Pages / Переведенные страницы

**Home Page / Главная страница:**
- ✅ Welcome message / Приветственное сообщение
- ✅ Subtitle / Подзаголовок
- ✅ All 12 feature cards / Все 12 карточек возможностей
- ✅ Documentation links / Ссылки на документацию
- ✅ Buttons / Кнопки

**Dashboard / Панель управления:**
- ✅ Page title / Заголовок страницы
- ✅ Welcome message / Приветственное сообщение
- ✅ All dashboard cards / Все карточки панели
- ✅ Admin section / Раздел администратора

**Profile / Профиль:**
- ✅ Page title / Заголовок страницы
- ✅ User information / Информация о пользователе
- ✅ Action buttons / Кнопки действий

**Upload / Загрузка:**
- ✅ Page title / Заголовок страницы
- ✅ Upload sections / Разделы загрузки
- ✅ Form buttons / Кнопки форм

**Login / Вход:**
- ✅ Page title / Заголовок страницы
- ✅ Form labels / Метки форм
- ✅ Error messages / Сообщения об ошибках
- ✅ Login button / Кнопка входа

**WebSocket:**
- ✅ Page title / Заголовок страницы
- ✅ Connection status / Статус подключения
- ✅ Input placeholder / Подсказка ввода
- ✅ Send button / Кнопка отправки

### 3. Translated API Responses / Переведенные ответы API

All API endpoints now return localized messages:

Все API эндпоинты теперь возвращают локализованные сообщения:

- User management / Управление пользователями
- Email service / Email сервис
- Queue service / Сервис очередей
- Error messages / Сообщения об ошибках

### 4. New Documentation / Новая документация

- `docs/DEMO_DATA.md` - Demo data reference (bilingual)
- `docs/DEMO_DATA.ru.md` - Russian version
- `docs/LANGUAGE_SWITCHING.md` - Language switching guide
- `docs/DEMO_EXAMPLES.md` - Usage examples with both languages

## Files Modified / Измененные файлы

### Backend / Бэкенд

1. **internal/database/seeder.go**
   - Added `SeedDemoUsers()` function
   - Creates 6 demo users (3 English + 3 Russian)

2. **cmd/server/main.go**
   - Added call to `database.SeedDemoUsers()`

3. **internal/controllers/api_example.go**
   - All messages now use i18n: `i18nT(c, "api.user.created")`

4. **internal/controllers/email_example.go**
   - Email service messages localized

5. **internal/controllers/queue_example.go**
   - Queue service messages localized

6. **internal/controllers/home_controller.go**
   - Added locale and T function to HomePage, DashboardPage, ProfilePage

7. **internal/controllers/upload_example.go**
   - Added i18n support to UploadPage

8. **internal/controllers/auth_controller.go**
   - Added i18n to LoginPage and LoginPost
   - Error messages localized

9. **internal/controllers/websocket.go**
   - Added i18n to WebSocketDemo

### Frontend / Фронтенд

7. **templates/pages/home.html**
   - All text replaced with `{{call .T "key"}}` template calls
   - Dynamic documentation links based on locale

8. **templates/pages/dashboard.html**
   - Dashboard fully translated
   - Welcome message with username

9. **templates/pages/profile.html**
   - Profile page translated
   - User information labels

10. **templates/pages/upload.html**
    - Upload page headers translated
    - Form buttons localized

11. **templates/pages/login.html**
    - Login page fully translated
    - Form labels and buttons

12. **templates/pages/websocket.html**
    - WebSocket page translated
    - Status messages and buttons

13. **templates/partials/header.html**
    - Added visual language switcher
    - Shows both EN and RU options
    - Active language highlighted

14. **templates/partials/base_foot.html**
    - Added switchLanguage() JavaScript function
    - Preserves URL parameters when switching

### Localization / Локализация

13. **locales/en.json**
    - Added 34+ keys for home page
    - Added 10+ keys for dashboard
    - Added 6+ keys for profile
    - Added 10+ keys for upload
    - Added 8+ keys for login
    - Added 6+ keys for websocket
    - Added 20+ API message keys

14. **locales/ru.json**
    - Added 34+ keys for home page
    - Added 10+ keys for dashboard
    - Added 6+ keys for profile
    - Added 10+ keys for upload
    - Added 8+ keys for login
    - Added 6+ keys for websocket
    - Added 20+ API message keys

### Documentation / Документация

15. **README.md** - Added link to Demo Data docs
16. **README.ru.md** - Added link to Demo Data docs
17. **docs/DEMO_DATA.md** - New bilingual documentation
18. **docs/DEMO_DATA.ru.md** - Russian version
19. **docs/LANGUAGE_SWITCHING.md** - Language switching guide
20. **docs/DEMO_EXAMPLES.md** - Usage examples

## How to Use / Как использовать

### Switch Language in Browser / Переключить язык в браузере

**Method 1: Click the language switcher in header**

Look for 🇬🇧 EN / 🇷🇺 RU buttons in the top navigation bar. Click to switch instantly!

Найдите кнопки 🇬🇧 EN / 🇷🇺 RU в верхней панели навигации. Нажмите для мгновенного переключения!

**Method 2: Use URL parameter**

```
http://localhost:8080/?lang=en    # English
http://localhost:8080/?lang=ru    # Russian
```

### Switch Language in API / Переключить язык в API

```bash
# English
curl -H "Accept-Language: en" http://localhost:8080/api/users

# Russian
curl -H "Accept-Language: ru" http://localhost:8080/api/users
```

### Login with Demo Users / Вход с демо пользователями

```bash
# English user
curl -X POST http://localhost:8080/login \
  -d "username=john_doe&password=password123"

# Russian user
curl -X POST http://localhost:8080/login \
  -d "username=ivan_ivanov&password=password123"
```

## Translation Keys Added / Добавленные ключи переводов

### Home Page / Главная страница
- `home.welcome` - Welcome message
- `home.subtitle` - Subtitle
- `home.documentation` - Documentation title
- `home.feature.*` - 12 feature cards (title + description)
- `home.doc.*` - Documentation link labels

### API Messages / API сообщения
- `api.user.*` - User management (6 keys)
- `api.email.*` - Email service (8 keys)
- `api.queue.*` - Queue service (6 keys)

### Demo Data / Демо данные
- `demo.user.*` - Demo user names
- `demo.email.*` - Demo email templates

**Total: 90+ new translation keys**

## Testing / Тестирование

### Quick Test / Быстрый тест

1. Start server / Запустите сервер:
```bash
go run cmd/server/main.go
```

2. Open in browser / Откройте в браузере:
```
http://localhost:8080/?lang=en
http://localhost:8080/?lang=ru
```

3. Verify demo users created / Проверьте создание демо пользователей:
```bash
# Check logs for:
# Demo user created: john_doe (password: password123)
# Demo user created: ivan_ivanov (password: password123)
```

### API Test / Тест API

```bash
# Test English response
curl -H "Accept-Language: en" http://localhost:8080/api/email/status

# Test Russian response
curl -H "Accept-Language: ru" http://localhost:8080/api/email/status
```

## Benefits / Преимущества

✅ **Better UX** - Users can use the app in their native language
✅ **Demo Ready** - Perfect for demonstrations to international audiences
✅ **Easy to Extend** - Simple to add more languages
✅ **Consistent** - All UI and API messages are translated
✅ **Professional** - Shows attention to internationalization

✅ **Лучший UX** - Пользователи могут использовать приложение на родном языке
✅ **Готово для демо** - Идеально для демонстраций международной аудитории
✅ **Легко расширить** - Просто добавить больше языков
✅ **Последовательно** - Все UI и API сообщения переведены
✅ **Профессионально** - Показывает внимание к интернационализации

## Next Steps / Следующие шаги

To add more languages:

Чтобы добавить больше языков:

1. Create `locales/[lang].json` (e.g., `es.json` for Spanish)
2. Copy all keys from `en.json` and translate
3. Restart server
4. Use `?lang=[lang]` in URL

## Documentation Links / Ссылки на документацию

- [Demo Data Documentation](docs/DEMO_DATA.md)
- [Language Switching Guide](docs/LANGUAGE_SWITCHING.md)
- [Demo Examples](docs/DEMO_EXAMPLES.md)
- [i18n Guide](docs/I18N.md)

---

**Created**: December 2, 2024
**Languages**: English (en), Russian (ru)
**Demo Users**: 6 (3 English + 3 Russian)
**Translation Keys**: 100+
**Translated Pages**: Home, Dashboard, Profile, Upload, Login, WebSocket
