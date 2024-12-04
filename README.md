
# GoResizer

Цей проект виконує такі завдання:
- Завантаження файлів у MinIO.
- Комунікація через RabbitMQ.
- Зберігання даних у MongoDB.
- Стиснення файлів через утиліту компресії.

## Системні вимоги

- Docker і Docker Compose
- Go 1.20 або новіше
- Доступ до MinIO, RabbitMQ, MongoDB

## Установка та запуск

### 1. Клонування репозиторію

```bash
git clone <URL_REPO>
cd goresizer
```

### 2. Налаштування

- Відредагуйте файл `config.yml` для підключення до MinIO, RabbitMQ та MongoDB.

### 3. Запуск через Docker

```bash
docker-compose up --build
```

### 4. Виконання `consumer.go`

Після запуску основного сервісу виконайте:

```bash
docker-compose exec app go run consumer/consumer.go
```

### 5. Тестування

Використовуйте API для завантаження (`/upload`) і завантаження файлів (`/download`).

## Структура проекту

- **main.go**: Основний файл для запуску API.
- **consumer/consumer.go**: Логіка споживача RabbitMQ.
- **internal/handlers/**: Хендлери для API.
- **internal/utils/**: Утиліти (JWT, компресія).
- **internal/user/**: Робота з користувачами.
- **pkg/**: Логування та підключення до MongoDB.
