
# GoResizer

This project performs the following tasks:
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

### 2. Customize
using the command
``` bash
make up
```
-start the services that are in `docker-compose.yml` get access keys for minIO
### 3. Configuration.

- Edit the `config.yml` file to connect to MinIO, RabbitMQ and MongoDB.``` yaml
mongodb:
  host: <your host> 
  port: <your port> 
  database: <your database name> 
  auth_db: <your auth_token> 
  username: <your username>
  password: <your password>
  collection: <your_collection>

minio:
  endpoint: <your endpoint> 
  storage: <your storage name>
  secret_k: <your secret key> 
  access_k: <your access key> 

rabbitmq:
  username: "<your username>"
  password: "<your password>"
  host: "<your hots>"
  port: "<your port>"
  queuename: "<your queue name (queue will generate automatically)>"
### 4. Tun project
`make up` - run docker-compose with services minIO, MongoDB, RabbitMQ
`make down` - stop docker-compose
`make logs` - check docker-compose logs
` make run` - run all project

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
