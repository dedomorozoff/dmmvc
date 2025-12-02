// Language switcher functionality
class LanguageSwitcher {
    constructor() {
        this.currentLocale = this.detectLocale();
        this.init();
    }

    detectLocale() {
        // Check cookie
        const cookies = document.cookie.split(';');
        for (let cookie of cookies) {
            const [name, value] = cookie.trim().split('=');
            if (name === 'locale') {
                return value;
            }
        }
        return 'en';
    }

    init() {
        // Set current locale in select
        const select = document.getElementById('locale-select');
        if (select) {
            select.value = this.currentLocale;
            select.addEventListener('change', (e) => {
                this.changeLocale(e.target.value);
            });
        }
    }

    async changeLocale(locale) {
        try {
            const response = await fetch('/api/locale', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ locale })
            });

            const result = await response.json();

            if (result.success) {
                // Reload page to apply new locale
                window.location.reload();
            } else {
                console.error('Failed to change locale:', result.error);
            }
        } catch (error) {
            console.error('Error changing locale:', error);
        }
    }
}

// Initialize language switcher when DOM is ready
if (document.readyState === 'loading') {
    document.addEventListener('DOMContentLoaded', () => {
        new LanguageSwitcher();
    });
} else {
    new LanguageSwitcher();
}
