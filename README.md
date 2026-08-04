# Event-Driven Email Service with RabbitMQ

A simple Go project demonstrating asynchronous email processing using RabbitMQ.

Instead of sending emails directly from the API, email requests are published to a RabbitMQ queue. A separate consumer listens to the queue and processes the emails asynchronously.

---

## Tech Stack

- Go
- Gin
- RabbitMQ
- JSON
- Goroutines
- Channels

---

## Architecture

```text
                HTTP Request
                     │
                     ▼
              Gin REST API
                     │
                     ▼
          Publish Message (JSON)
                     │
                     ▼
             RabbitMQ Queue
                     │
                     ▼
          Consumer (Goroutine)
                     │
                     ▼
          Process Email Request
                     │
                     ▼
              Send Email
```

---

## Flow

1. Client sends an HTTP request.
2. Gin receives the request.
3. The API publishes a message to RabbitMQ.
4. RabbitMQ stores the message in the queue.
5. The consumer listens for new messages.
6. The message is unmarshaled into a Go struct.
7. The email is processed and sent asynchronously.

---

## Project Structure

```text
.
├── cmd/
├── handlers/
├── rabbitmq/
├── models/
├── config/
├── main.go
└── README.md
```

---

## Run RabbitMQ

Using Docker

```bash
docker run -d \
  --hostname rabbitmq \
  --name rabbitmq \
  -p 5672:5672 \
  -p 15672:15672 \
  rabbitmq:3-management
```

RabbitMQ Dashboard

```
http://localhost:15672
```

Default credentials

```
Username: guest
Password: guest
```

---

## Install Dependencies

```bash
go mod tidy
```

---

## Run the API

```bash
go run main.go
```

or

```bash
go run ./cmd/api
```

---

## Run the Consumer

```bash
go run ./cmd/consumer
```

---

## Example Request

```http
POST /projects
Content-Type: application/json

{
    "reference_id": "REF-001",
    "client_name": "Mahmoud",
    "email": "mahmoud@example.com",
    "company": "ValtQ",
    "project_name": "Website Development"
}
```

---

## Queue Message

```json
{
    "reference_id": "REF-001",
    "client_name": "Mahmoud",
    "email": "mahmoud@example.com",
    "company": "ValtQ",
    "project_name": "Website Development"
}
```

---

## Learning Objectives

- RabbitMQ Fundamentals
- Producer / Consumer Pattern
- Event-Driven Architecture
- Asynchronous Processing
- Goroutines
- Channels
- JSON Serialization
- Queue-Based Communication

---

## Author

Mahmoud Ramadan Abbas
