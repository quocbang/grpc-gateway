# API Reference

## Accessing Documentation
1. Start the server using `make run`
2. Visit [Swagger UI](http://localhost:8080/swagger-ui/) for interactive API documentation

## Endpoints
- Authentication: `/v1/auth/*`
- User Management: `/v1/users/*`
- Product Management: `/v1/products/*`

## Generating Documentation
```bash
# Requires protoc and grpc-gateway
make generate
```
Generated OpenAPI specs will be available in `docs/swagger/`
