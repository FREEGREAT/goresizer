
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
git clone https://github.com/FREEGREAT/goresizer.git
cd goresizer
```

### 2. Налаштування

- Відредагуйте файл `config.yml` для підключення до MinIO, RabbitMQ та MongoDB.

### 3. Запуск проекту
```bash
make up && make run-main
```

### 5. Тестування
Звертаючись до ендпоінта
http://localhost:8080/login та http://localhost:8080/signup
потрібно передати `email` `password`
``` json
{
  "email": "testuser@example.com",
  "password": "password123"
}
```

Звертаючись до ендпоінта
http://localhost:8080/api/upload?resizepercent=0.5
потрібно передати `resizepercent` який слугує значенням компресії фото в %, відповідно до цього значення в посиланні може бути тільки в границі (0,1)
Також потрібно передати в хедер `Accsess` токен
![зображення](https://github.com/user-attachments/assets/8f0e11ff-c574-4118-a712-d000242dd2f5)

Звертаючись до ендпоінта
http://localhost:8080/api/download?filename=image.jpg
Потрібно передати id(назву)файлу, який зберігається в minIO для того, щоб завантажити його. Зберігається файл в директорії `/tmp/download/pp`
Також потрібно передати в хедер `Accsess` токен
![зображення](https://github.com/user-attachments/assets/562265d2-a9a5-4877-b17b-85e489cfd5a5)


