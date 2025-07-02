# Development Guide

## Code Structure
```
.
├── cmd/          # CLI commands and entry points
├── config/       # Configuration types and parsing
├── server/       # Core server implementation
│   ├── db/       # Database migrations and connection handling
│   ├── repositories/ # Data access layer
│   ├── services/ # Business logic and gRPC service implementations
│   ├── sender/   # Email sending implementation
│   └── utils/    # Shared utilities (hashing, tokens, etc.)
├── docs/         # Project documentation
└── script/       # Database initialization scripts
```

## Testing
- Run unit tests: `make test`
- Integration tests require running PostgreSQL and Redis instances
- Use mock implementations for isolated testing

## Code Style
- Follow Go's standard formatting (gofmt)
- Use zerolog for structured logging
- Keep line length under 120 characters
- Write comprehensive unit tests for new features
