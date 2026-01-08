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
SMTP_PORT=6767
SMTP_USER=your_smtp_username
SMTP_PASS=your_smtp_password
TO_ADDRESS=personal@example.com
FROM_ADDRESS=registered_sender@example.com
APP_PORT=6969
```
Make sure `.env` is added to `.gitignore` to avoid committing secrets.

## Running Locally

Install dependencies and run the service:

```bash
go mod tidy
go run ./cmd/email-service
