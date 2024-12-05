
# GoResizer

Цей проект виконує такі завдання:
- Storage files in MinIO.
- Communicate using by RabbitMQ.
- Storage users ing MongoDB.
- Compress files using func `Compress()` and gomodule `"github.com/nfnt/resize"`.

## System requiretments

- Docker Compose
- Go 1.23
- MinIO, RabbitMQ, MongoDB

## Install and running

### 1. Clone repo

```bash
git clone https://github.com/FREEGREAT/goresizer.git
git clone "<URL_REPO>"
cd goresizer
```

### 2. Налаштування
використовуючи команду
```bash
make up
```
-запустіть сервіси, які знаходяться в `docker-compose.yml` отримайте ключі доступу для minIO
### 2. Configuration

- Відредагуйте файл `config.yml` для підключення до MinIO, RabbitMQ та MongoDB.
``` yaml
mongodb:
  host: <your host> reuired
  port: <your port> required
  database: <your database name> required
  auth_db: <your auth_token> not required
  username: <your username> not required
  password: <your password> not required
  collection: <your_collection> required

minio:
  endpoint: <your endpoint> required
  storage: <your storage name> required
  secret_k: <your secret key> required
  access_k: <your access key> required

rabbitmq:
  username: "<your username>"
  password: "<your password>"
  host: "<your hots>"
  port: "<your port>"
  queuename: "<your queue name (queue generate automatically)>"


```

### 3. Run using by docker

```bash
docker-compose up --build
```

### 5. Testing

Calling to endpoints
http://localhost:8080/login and http://localhost:8080/signup
you must to transfer  `email` and `password`
``` json
{
  "email": "testuser@example.com",
  "password": "password123"
}
```

Calling to endpoint
http://localhost:8080/api/upload?resizepercent=0.5
you must to trasfer `resizepercent` which serves as a photo compression value in %, accordingly, this value in the link can only be in the range (0,1)
You also need to pass a token to the `Accsess` header
![Picture](https://github.com/user-attachments/assets/8f0e11ff-c574-4118-a712-d000242dd2f5)

Addressing the endpoint
http://localhost:8080/api/download?filename=image.jpg
You need to pass the id (name) of the file stored in minIO in order to download it. The file is stored in the directory `/tmp/download/pp`.
You also need to pass a token to the `Accsess` header
![Picture](https://github.com/user-attachments/assets/562265d2-a9a5-4877-b17b-85e489cfd5a5)



## Project structure

- **main.go**: The main file to run the API.
- **internal/handlers/**:Handlers for the API.
- **internal/utils/**: Utilities (JWT, compression).
- **internal/user/**: Work with users.
- **internal/pkg/**: Logging and connection to MongoDB, Minio.
