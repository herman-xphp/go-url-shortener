# High-Performance URL Shortener

A high-performance URL shortener service built with Go (Golang), Fiber, Redis, and PostgreSQL with a modern web interface.

## Features

- ✅ **Modern Web UI** - Beautiful, responsive interface for desktop and mobile
- ✅ **Shorten URLs** - Convert long URLs to compact short codes instantly
- ✅ **Custom Aliases** - Create branded short links with custom aliases
- ✅ **Click Tracking** - Track how many times your links are clicked
- ✅ **Fast Redirects** - Redis caching for lightning-fast redirects
- ✅ **Persistent Storage** - PostgreSQL for reliable data persistence
- ✅ **Clean Architecture** - Well-organized, maintainable codebase
- ✅ **Docker Support** - Easy deployment with Docker Compose

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

## Usage

### Web Interface

1. Open your browser and navigate to `http://localhost:3000`
2. Enter your long URL in the input field
3. (Optional) Add a custom alias for your short link
4. Click "Shorten URL" button
5. Copy your new short URL and share it!

The web interface features:
- 🎨 Modern, gradient design
- 📱 Fully responsive (works great on mobile)
- ⚡ Real-time validation
- 📋 One-click copy to clipboard
- 🎯 Custom alias support
- ✨ Smooth animations and transitions

### API Endpoints

You can also use the REST API directly:

#### 1. Health Check
```bash
GET /health
```

#### 2. Shorten URL
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

#### 3. Redirect to Original URL
```bash
GET /:shortCode
```
Redirects to the original URL and tracks click count.

## Access from Other Devices

Want to test from your phone or tablet? The URL shortener can be accessed from any device on the same network.

**Steps:**
1. Find your server IP: `ip addr show | grep "inet " | grep -v "127.0.0.1"`
2. Open firewall port: `./firewall-control.sh open`
3. Access from phone: `http://YOUR_SERVER_IP:3000`
4. Close port when done: `./firewall-control.sh close`

See [Firewall Management](#firewall-management) for details.

## Firewall Management

Use the provided script to easily manage firewall access:

```bash
# Open port temporarily (for testing)
./firewall-control.sh open

# Close port after testing
./firewall-control.sh close

# Check port status
./firewall-control.sh status

# Open permanently (with confirmation)
./firewall-control.sh permanent-open

# Close permanently
./firewall-control.sh permanent-close
```

**Security Best Practice:**
- Use `open` for testing sessions only
- Always `close` when done
- Avoid `permanent-open` unless necessary

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

See `.env.example` for all available configuration options:
- `SERVER_PORT` - HTTP server port (default: 3000)
- `DB_HOST`, `DB_PORT` - PostgreSQL connection
- `REDIS_ADDR` - Redis server address
- `BASE_URL` - Base URL for shortened links
- `SHORT_CODE_LENGTH` - Length of generated codes (default: 7)

## License

MIT
