# High-Performance URL Shortener

A high-performance URL shortener service built with Go (Golang), Fiber, Redis, and PostgreSQL.

## Features

- ✅ Shorten long URLs to compact short codes
- ✅ Custom alias support
- ✅ Click tracking and analytics
- ✅ Redis caching for fast redirects
- ✅ PostgreSQL for persistent storage
- ✅ Clean Architecture design
- ✅ Docker support

## Project Structure

This project follows the Standard Go Project Layout (Clean Architecture):

- `cmd/api`: Application entry point
- `internal/core`: Business logic (Domain models, Ports, Services)
- `internal/adapters`: Implementations of ports (PostgreSQL, Redis)
- `internal/api`: HTTP handlers and routes
- `pkg`: Shared libraries (Config, Logger)

## Tech Stack

- **Language**: Go 1.21+
- **Web Framework**: Fiber v2
- **Database**: PostgreSQL 15
- **Cache**: Redis 7
- **Containerization**: Docker & Docker Compose

## Getting Started

### Prerequisites

- Go 1.21+ installed
- PostgreSQL 15+ (or use Docker)
- Redis 7+ (or use Docker)
- Docker & Docker Compose (optional)

### Installation

#### Option 1: Using Docker Compose (Recommended)

1. Clone the repository
2. Copy environment file:
   ```bash
   cp .env.example .env
   ```
3. Start all services:
   ```bash
   docker-compose up -d
   ```
4. API will be available at `http://localhost:3000`

#### Option 2: Manual Setup

1. Clone the repository
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Copy and configure environment:
   ```bash
   cp .env.example .env
   # Edit .env with your database credentials
   ```
4. Start PostgreSQL and Redis (or use Docker):
   ```bash
   docker-compose up -d postgres redis
   ```
5. Run the application:
   ```bash
   go run cmd/api/main.go
   ```

## API Endpoints

### 1. Health Check
```bash
GET /health
```

### 2. Shorten URL
```bash
POST /api/shorten
Content-Type: application/json

{
  "original_url": "https://example.com/very/long/url",
  "custom_alias": "mylink",  // optional
  "expires_at": "2024-12-31T23:59:59Z"  // optional
}
```

Response:
```json
{
  "short_url": "http://localhost:3000/abc123",
  "original_url": "https://example.com/very/long/url",
  "short_code": "abc123",
  "created_at": "2024-12-03T10:00:00Z"
}
```

### 3. Redirect to Original URL
```bash
GET /:shortCode
```
Redirects to the original URL and tracks click count.

## Testing

```bash
# Test URL shortening
curl -X POST http://localhost:3000/api/shorten \
  -H "Content-Type: application/json" \
  -d '{"original_url": "https://github.com"}'

# Test redirect
curl -L http://localhost:3000/abc123
```

## Environment Variables

See `.env.example` for all available configuration options.

## License

MIT
