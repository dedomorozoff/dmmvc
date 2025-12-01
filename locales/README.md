# Translation Files

This directory contains translation files for the DMMVC application.

## Available Languages

- `en.json` - English (default)
- `ru.json` - Russian

## File Format

Translation files use JSON format with key-value pairs:

```json
{
  "key": "Translation text",
  "greeting": "Hello, %s!",
  "app.name": "Application Name"
}
```

## Adding New Languages

1. Create a new JSON file with the locale code (e.g., `es.json` for Spanish)
2. Copy the structure from `en.json`
3. Translate all values
4. Update `internal/i18n/i18n.go` to include the new locale
5. Update middleware and handlers

## Translation Keys

Use hierarchical keys for better organization:

- `app.*` - Application-level translations
- `nav.*` - Navigation items
- `auth.*` - Authentication related
- `common.*` - Common UI elements
- `error.*` - Error messages

## Placeholders

Use standard Go format specifiers:
- `%s` - String
- `%d` - Integer
- `%f` - Float

Example:
```json
{
  "welcome": "Welcome, %s!",
  "count": "You have %d items"
}
```

## See Also

- [i18n Documentation](../docs/I18N.md)
- [Examples](../docs/EXAMPLES.md)
