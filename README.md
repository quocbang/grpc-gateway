# gRPC Gateway Server

A gRPC-based server with REST API gateway, handling authentication, database interactions, and asynchronous tasks.

## Features
- REST API Gateway
- JWT Authentication
- PostgreSQL database integration
- Redis caching
- Async task processing
- SMTP email integration

## Prerequisites
- Go 1.20+
- Docker
- PostgreSQL
- Redis

## Installation
```bash
git clone https://github.com/quocbang/grpc-gateway.git
cd grpc-gateway
go mod tidy
```

## Configuration
1. Create `server/config.yaml`:
```yaml
dev_mode: false
database:
  postgres:
    address: localhost
    port: 54324
    name: quocbang
    schema: quocbang
    username: postgres
    password: quocbang
  redis:
    address: localhost:63795
    password: api_server_basic_password
server:
  sender:
    smtp:
      smtp_server: "smtp.gmail.com"
      smtp_port: 587
      sender_email: "your-email@gmail.com"
      password: "your-app-password"
  auth:
    secret_key_path: ./server/secret/secret_key.key
    access_token_life_time: 12m
    refresh_token_life_time: 168h
```

2. Generate secret key:
```bash
mkdir -p server/secret
openssl rand -hex 64 > server/secret/secret_key.key
```

## Running the Server
```bash
make build    # Build binary
make local-db # Start database containers
make run      # Start server
```

## Testing
```bash
make test
```

## API Documentation
Access Swagger UI at `http://localhost:8080/swagger/` after starting the server.

## License
MIT
