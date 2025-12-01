// DMMVC Framework JavaScript

document.addEventListener('DOMContentLoaded', function() {
    console.log('DMMVC Framework loaded');
    
    // Автоматическое скрытие alert сообщений через 5 секунд
    const alerts = document.querySelectorAll('.alert');
    alerts.forEach(alert => {
        setTimeout(() => {
            alert.style.opacity = '0';
            setTimeout(() => {
                alert.remove();
            }, 300);
        }, 5000);
    });
    
    // Подтверждение удаления
    const deleteForms = document.querySelectorAll('form[action*="/delete"]');
    deleteForms.forEach(form => {
        form.addEventListener('submit', function(e) {
            if (!confirm('Вы уверены, что хотите удалить?')) {
                e.preventDefault();
            }
        });
    });
});
