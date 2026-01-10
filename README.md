# Email Service

A Go microservice that receives contact form submissions and sends emails via SMTP.

## Features

- Receives HTTP POST requests with JSON payload containing name, email, subject, and message.
- Sends email to a configured recipient address.
- Sets visitor's email as the Reply-To address for direct replies.
- Uses Gin framework for HTTP server.
- Supports SMTP email sending via `wneessen/go-mail`.
- Dockerized for easy deployment.

## Project Structure

```
cmd/email-service/main.go
docker/
├── go/ # Docker file and script for golang app
└── nginx/ # Docker file and script for nginx
    └── certs/ # Folder to store certs
internal/
├── config/ # Configuration loader
├── handler/ # HTTP handlers
├── logger/ # Logging middleware
├── model/ # Email data models
├── sender/ # SMTP sender implementation
└── service/ # Business logic
Dockerfile
docker-compose.yml
go.mod
.env
README.md
```
## Configuration

Create a `.env` file in the project root with the following variables:
```yaml
SMTP_HOST=smtp.example.com
SMTP_PORT=587
SMTP_USER=your_smtp_username
SMTP_PASS=your_smtp_password
TO_ADDRESS=personal@example.com
FROM_ADDRESS=registered_sender@example.com
# Port your golang application is running inside it's docker container
PORT=6969
# Port your NGINX image is listening on inside the container
NGINX_PORT=80
# Port that allows for external interfacing
HOST_PORT=8000
```
Make sure `.env` is added to `.gitignore` to avoid committing secrets.

## Running Locally

### For only HTTP with Golang
Install dependencies and run the service:

```bash
go mod tidy
go run ./cmd/email-service
```
### For HTTP with Docker
```bash
docker-compose build --no-cache
docker-compose up -d

# To tear down
docker-compose down
```

# For HTTPS
```bash
openssl req -x509 -nodes -days 365 \   
-newkey rsa:2048 \
-keyout docker/nginx/certs/nginx.key \   
-out docker/nginx/certs/nginx.crt \ 
-subj "/CN=localhost"

docker-compose build --no-cache
docker-compose up -d

# To tear down
docker-compose down
```

