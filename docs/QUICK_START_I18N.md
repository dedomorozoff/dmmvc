# Quick Start: Bilingual Interface / Быстрый старт: Двуязычный интерфейс

## 🌍 Language Switcher / Переключатель языков

DMMVC now has a beautiful language switcher in the header!

DMMVC теперь имеет красивый переключатель языков в шапке!

### How to Use / Как использовать

1. **Look at the top navigation bar** / **Посмотрите на верхнюю панель навигации**
   
   You'll see: 🇬🇧 EN | 🇷🇺 RU

2. **Click on the language you want** / **Нажмите на нужный язык**
   
   - 🇬🇧 EN - English
   - 🇷🇺 RU - Русский

3. **The page will reload in the selected language** / **Страница перезагрузится на выбранном языке**

### Features / Возможности

✅ **Visual Switcher** - Easy to find and use
✅ **Active State** - Current language is highlighted
✅ **All Pages** - Available everywhere in the app
✅ **URL Preservation** - Keeps your current page and parameters
✅ **Smooth Animation** - Beautiful hover and transition effects

## 📄 Translated Pages / Переведенные страницы

All pages are fully bilingual:

Все страницы полностью двуязычные:

| Page | English | Russian |
|------|---------|---------|
| Home | ✅ | ✅ |
| Login | ✅ | ✅ |
| Dashboard | ✅ | ✅ |
| Profile | ✅ | ✅ |
| Upload | ✅ | ✅ |
| WebSocket | ✅ | ✅ |

## 🎨 Design / Дизайн

The language switcher features:

Переключатель языков имеет:

- **Flag Emojis** - 🇬🇧 🇷🇺 for visual recognition
- **Active State** - Highlighted current language
- **Hover Effects** - Smooth animations on hover
- **Responsive** - Works on mobile and desktop
- **Accessible** - Clear labels and tooltips

## 🚀 Demo Users / Демо пользователи

Try logging in with demo users:

Попробуйте войти с демо пользователями:

**English Users:**
```
Username: john_doe
Password: password123
```

**Russian Users:**
```
Username: ivan_ivanov
Password: password123
```

## 💡 Tips / Советы

1. **Language persists** - Your language choice is saved in cookies
2. **Works everywhere** - Switch language on any page
3. **API too** - API responses are also localized
4. **Easy to extend** - Add more languages by creating new locale files

## 🔧 For Developers / Для разработчиков

### Adding a new language / Добавление нового языка

1. Create `locales/es.json` (for Spanish)
2. Copy keys from `locales/en.json`
3. Translate all values
4. Add flag to header: 🇪🇸 ES
5. Restart server

### Using translations in code / Использование переводов в коде

```go
// In controller
locale := i18n.GetLocale(c)
c.HTML(http.StatusOK, "page.html", gin.H{
    "locale": string(locale),
    "T": i18n.TFunc(locale),
})
```

```html
<!-- In template -->
<h1>{{call .T "page.title"}}</h1>
<p>{{call .T "page.description"}}</p>
```

## 📚 More Documentation / Дополнительная документация

- [Full i18n Guide](I18N.md) - Complete internationalization guide
- [Language Switching](LANGUAGE_SWITCHING.md) - All methods to switch language
- [Demo Data](DEMO_DATA.md) - Bilingual demo users and data
- [Demo Examples](DEMO_EXAMPLES.md) - API examples in both languages

## 🎉 Try It Now! / Попробуйте сейчас!

1. Start the server / Запустите сервер:
   ```bash
   go run cmd/server/main.go
   ```

2. Open in browser / Откройте в браузере:
   ```
   http://localhost:8080
   ```

3. Click the language switcher in the header! / Нажмите переключатель языков в шапке!

---

**Enjoy the bilingual experience!** 🌍

**Наслаждайтесь двуязычным интерфейсом!** 🌍
