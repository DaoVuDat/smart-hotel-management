# Hotel Management System - Microservices Architecture

A comprehensive hotel management system built with microservices architecture, demonstrating real-world enterprise patterns and best practices.

## 🏗️ Architecture Overview

This system implements a distributed hotel management platform with the following core services:

- **User Service** (Go + Echo) - Authentication, user profiles, staff management
- **Booking Service** (Python + FastAPI) - Reservations, availability, pricing
- **Room Service** (Node.js + Fastify) - Room inventory, housekeeping, maintenance
- **Notification Service** (Go + Echo) - Email, SMS, push notifications
- **Payment Service** (Python + FastAPI) - Payment processing, billing, invoicing

## 🎯 Technology Stack

### Core Services

- **Go (Echo Framework)** - User & Notification Services
- **Python (FastAPI)** - Booking & Payment Services
- **Node.js (Fastify)** - Room Service

### Infrastructure

- **Databases**: PostgreSQL (primary), MongoDB (room data), Redis (cache/sessions)
- **Message Broker**: Apache Kafka
- **Authentication**: Keycloak
- **API Gateway**: Kong
- **Container Orchestration**: Kubernetes (K3s/K3d)
- **Development**: Tilt.dev
- **Monitoring**: Prometheus + Grafana

### Communication

- **External**: HTTP/REST via Kong API Gateway
- **Internal**: gRPC between services
- **Async**: Kafka for event-driven communication

## 🏛️ Service Architecture & Database Design

### Service-Database Mapping

```
User Service (Go)           → PostgreSQL (user_db)
├── Users, Staff, Roles
├── Authentication data
└── User preferences

Booking Service (Python)    → PostgreSQL (booking_db)
├── Reservations, Availability
├── Pricing rules
└── Guest information

Room Service (Node.js)      → MongoDB (room_db)
├── Room inventory (flexible schema)
├── Housekeeping status
├── Maintenance logs
└── Room features/amenities

Payment Service (Python)    → PostgreSQL (payment_db)
├── Transactions, Billing
├── Payment methods
└── Invoice history

Notification Service (Go)   → PostgreSQL (notification_db)
├── Message templates
├── Delivery logs
└── Preferences

Shared Cache/Sessions       → Redis
├── Session storage
├── Rate limiting
└── Temporary data
```

### Why This Database Distribution?

- **PostgreSQL**: ACID compliance for critical business data (users, bookings, payments)
- **MongoDB**: Flexible schema for room data with varying amenities and features
- **Redis**: High-performance caching and session management

## 🔄 Service Communication Flow

### Booking Flow Example

```
1. Client → Kong Gateway → Booking Service (check availability)
2. Booking Service → Room Service (gRPC: get available rooms)
3. Booking Service → User Service (gRPC: validate user)
4. Booking Service → Payment Service (gRPC: process payment)
5. Payment Service → Kafka (payment.completed event)
6. Notification Service ← Kafka (send confirmation)
7. Room Service ← Kafka (update room status)
```

### Event-Driven Architecture

```
Kafka Topics:
- booking.created
- booking.cancelled
- payment.completed
- payment.failed
- room.status.changed
- user.registered
- notification.sent
```

## 🏗️ Project Structure

```
hotel-management-system/
├── README.md
├── docker-compose.yml
├── Tiltfile
├── k8s/
│   ├── kong/
│   ├── keycloak/
│   ├── kafka/
│   ├── databases/
│   └── monitoring/
├── helm/
│   ├── user-service/
│   ├── booking-service/
│   ├── room-service/
│   ├── payment-service/
│   ├── notification-service/
│   └── infrastructure/
├── services/
│   ├── user-service/          # Go + Echo
│   │   ├── cmd/
│   │   ├── internal/
│   │   │   ├── api/
│   │   │   ├── domain/
│   │   │   ├── repository/
│   │   │   ├── db/
│   │   │   ├── service/
│   │   │   └── config/
│   │   ├── proto/
│   │   ├── migrations/
│   │   ├── Dockerfile
│   │   └── go.mod
│   ├── booking-service/       # Python + FastAPI
│   │   ├── app/
│   │   │   ├── api/
│   │   │   ├── core/
│   │   │   ├── models/
│   │   │   ├── services/
│   │   │   ├── repositories/
│   │   │   └── schemas/
│   │   ├── proto/
│   │   ├── alembic/
│   │   ├── Dockerfile
│   │   └── requirements.txt
│   ├── room-service/          # Node.js + Fastify
│   │   ├── src/
│   │   │   ├── controllers/
│   │   │   ├── models/
│   │   │   ├── services/
│   │   │   ├── repositories/
│   │   │   ├── routes/
│   │   │   └── proto/
│   │   ├── Dockerfile
│   │   └── package.json
│   ├── payment-service/       # Python + FastAPI
│   │   └── [similar structure to booking-service]
│   └── notification-service/  # Go + Echo
│       └── [similar structure to user-service]
├── shared/
│   ├── proto/                 # gRPC definitions
│   ├── events/                # Kafka event schemas
│   └── common/                # Shared utilities
├── web-ui/                    # Future frontend
└── docs/
    ├── api/
    ├── architecture/
    └── deployment/
```

## 🚀 Language & Framework Rationale

### User Service (Go + Echo)

**Why Go?**

- Excellent concurrency for handling authentication
- Strong typing for security-critical operations
- Fast compilation and deployment

**Why Echo?**

- Lightweight, high-performance HTTP router
- Built-in middleware for JWT, CORS, logging
- Clean, simple API design

**Key Libraries:**

```go
// HTTP Framework
github.com/labstack/echo/v4
github.com/labstack/echo-jwt/v4

// Database
github.com/lib/pq (PostgreSQL driver)
github.com/jmoiron/sqlx

// gRPC
google.golang.org/grpc

// Configuration
github.com/spf13/viper

// Authentication
github.com/golang-jwt/jwt/v4

// Validation
github.com/go-playground/validator/v10
```

### Booking Service (Python + FastAPI)

**Why Python?**

- Excellent for complex business logic and calculations
- Rich ecosystem for data processing
- Great for pricing algorithms and availability calculations

**Why FastAPI?**

- Automatic API documentation
- Built-in data validation with Pydantic
- Async support for high performance

**Key Libraries:**

```python
# Web Framework
fastapi
uvicorn[standard]

# Database
asyncpg
databases[postgresql]
alembic

# gRPC
grpcio
grpcio-tools

# Validation & Serialization
pydantic

# Configuration
python-decouple

# Async HTTP Client
httpx

# Background Tasks
celery
redis
```

### Room Service (Node.js + Fastify)

**Why Node.js?**

- Excellent for I/O intensive operations (room status updates)
- Great for real-time features (room availability updates)
- Natural fit for MongoDB with rich ODM options

**Why Fastify?**

- Fastest Node.js web framework
- Built-in JSON schema validation
- Plugin architecture

**Key Libraries:**

```javascript
// Web Framework
fastify
@fastify/cors
@fastify/jwt

// Database
mongoose (MongoDB ODM)
ioredis (Redis client)

// gRPC
@grpc/grpc-js
@grpc/proto-loader

// Validation
ajv
fluent-json-schema

// Configuration
dotenv

// Logging
pino
```

### Payment Service (Python + FastAPI)

**Why Python for payments?**

- Robust libraries for financial calculations
- Excellent error handling for critical operations
- Strong typing with Pydantic for data validation

### Notification Service (Go + Echo)

**Why Go for notifications?**

- Excellent concurrency for handling multiple notification channels
- Reliable for critical communication
- Fast processing of notification queues

## 🔧 Inter-Service Communication

### HTTP Communication (External)

```
Client → Kong Gateway → Service
- Authentication via Keycloak JWT
- Rate limiting and throttling
- Request/response logging
- Load balancing
```

### gRPC Communication (Internal)

```proto
// Example: User Service Proto
syntax = "proto3";

package user;

service UserService {
  rpc GetUser(GetUserRequest) returns (GetUserResponse);
  rpc ValidateUser(ValidateUserRequest) returns (ValidateUserResponse);
}

message GetUserRequest {
  string user_id = 1;
}

message GetUserResponse {
  string user_id = 1;
  string email = 2;
  string name = 3;
  UserRole role = 4;
}
```

### Kafka Event Schema

```json
{
  "booking.created": {
    "booking_id": "uuid",
    "user_id": "uuid",
    "room_id": "uuid",
    "check_in": "datetime",
    "check_out": "datetime",
    "total_amount": "decimal",
    "timestamp": "datetime"
  }
}
```

## 🛡️ Security & Best Practices

### Authentication Flow

1. User authenticates with Keycloak
2. Keycloak issues JWT token
3. Kong validates JWT for all requests
4. Services receive validated user context

### Circuit Breker Pattern

```go
// Go example with circuit breaker
type CircuitBreaker struct {
    failures    int
    threshold   int
    timeout     time.Duration
    state       State
}

func (cb *CircuitBreaker) Call(operation func() error) error {
    if cb.state == Open {
        return ErrCircuitOpen
    }

    err := operation()
    if err != nil {
        cb.failures++
        if cb.failures >= cb.threshold {
            cb.state = Open
        }
        return err
    }

    cb.reset()
    return nil
}
```

### Retry Pattern with Exponential Backoff

```python
# Python example
@retry(
    retry=retry_if_exception_type(httpx.HTTPError),
    wait=wait_exponential(multiplier=1, min=4, max=10),
    stop=stop_after_attempt(3)
)
async def call_external_service():
    async with httpx.AsyncClient() as client:
        response = await client.get("http://service/api")
        return response.json()
```

### Outbox Pattern Implementation

```sql
-- Outbox table for reliable message publishing
CREATE TABLE outbox_events (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    aggregate_id UUID NOT NULL,
    event_type VARCHAR(255) NOT NULL,
    payload JSONB NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    processed_at TIMESTAMP NULL
);
```

## 🗄️ Database Schemas

### User Service (PostgreSQL)

```sql
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    phone VARCHAR(50),
    role user_role NOT NULL DEFAULT 'guest',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TYPE user_role AS ENUM ('guest', 'staff', 'manager', 'admin');
```

### Booking Service (PostgreSQL)

```sql
CREATE TABLE bookings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    room_id UUID NOT NULL,
    check_in DATE NOT NULL,
    check_out DATE NOT NULL,
    guests_count INTEGER NOT NULL,
    total_amount DECIMAL(10,2) NOT NULL,
    status booking_status DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TYPE booking_status AS ENUM ('pending', 'confirmed', 'cancelled', 'completed');
```

### Room Service (MongoDB)

```javascript
// Room Document Schema
{
  _id: ObjectId,
  room_number: String,
  room_type: {
    name: String,
    base_price: Number,
    capacity: Number
  },
  amenities: [String],
  status: String, // available, occupied, maintenance, cleaning
  floor: Number,
  features: {
    has_balcony: Boolean,
    has_kitchen: Boolean,
    wifi_speed: String,
    // flexible schema allows easy addition of new features
  },
  housekeeping: {
    last_cleaned: Date,
    assigned_staff: String,
    notes: String
  },
  created_at: Date,
  updated_at: Date
}
```

## 🚀 Development Setup

### Prerequisites

- Docker & Docker Compose
- Kubernetes (K3s/K3d)
- Tilt.dev
- Go 1.21+
- Python 3.11+
- Node.js 18+

### Local Development with Tilt

```bash
# Clone repository
git clone <repo-url>
cd hotel-management-system

# Start K3d cluster
k3d cluster create hotel-dev --port "8080:80@loadbalancer"

# Install Tilt
curl -fsSL https://raw.githubusercontent.com/tilt-dev/tilt/master/scripts/install.sh | bash

# Start development environment
tilt up

# Access services
# Kong Gateway: http://localhost:8080
# Keycloak: http://localhost:8081
# Grafana: http://localhost:3000
```

### Production Deployment

```bash
# Deploy infrastructure
helm install kong kong/kong -f helm/kong/values.yaml
helm install keycloak bitnami/keycloak -f helm/keycloak/values.yaml
helm install kafka bitnami/kafka -f helm/kafka/values.yaml

# Deploy services
helm install user-service ./helm/user-service
helm install booking-service ./helm/booking-service
helm install room-service ./helm/room-service
helm install payment-service ./helm/payment-service
helm install notification-service ./helm/notification-service
```

## 📊 Monitoring & Observability

### Metrics Collection

- **Prometheus** for metrics collection
- **Grafana** for visualization
- **Jaeger** for distributed tracing

### Key Metrics to Monitor

- Request latency and throughput
- Database connection pool usage
- Kafka consumer lag
- Circuit breaker states
- Authentication success/failure rates

## 🔄 CI/CD Pipeline

```yaml
# .github/workflows/ci.yml
name: CI/CD Pipeline
on:
  push:
    branches: [main, develop]
  pull_request:
    branches: [main]

jobs:
  test:
    runs-on: ubuntu-latest
    services:
      postgres:
        image: postgres:15
        env:
          POSTGRES_PASSWORD: test
        options: >-
          --health-cmd pg_isready
          --health-interval 10s
          --health-timeout 5s
          --health-retries 5
    steps:
      - uses: actions/checkout@v3
      - name: Run tests
        run: make test-all

  build-and-deploy:
    needs: test
    runs-on: ubuntu-latest
    if: github.ref == 'refs/heads/main'
    steps:
      - name: Build and push images
        run: make build-images
      - name: Deploy to staging
        run: make deploy-staging
```

## 🧪 Testing Strategy

### Unit Tests

- Go: `go test` with testify
- Python: `pytest` with pytest-asyncio
- Node.js: `jest` with supertest

### Integration Tests

- Database integration tests
- gRPC service tests
- Kafka event publishing/consuming tests

### End-to-End Tests

- API workflow tests
- Cross-service communication tests

## 📚 API Documentation

### REST API Documentation

- Swagger/OpenAPI 3.0 for all HTTP APIs
- Auto-generated from FastAPI and Echo
- Available at `/docs` endpoint for each service

### gRPC Documentation

- Protocol buffer definitions in `/shared/proto`
- Generated client libraries for each language
- gRPC reflection enabled for development

## 🔮 Future Enhancements

### Real-time Features

- WebSocket support for live room availability
- Real-time booking notifications
- Live housekeeping status updates

### Advanced Features

- Machine learning for pricing optimization
- IoT integration for smart room controls
- Mobile app integration
- Multi-tenant support for hotel chains

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Write tests for new functionality
4. Ensure all tests pass
5. Create a Pull Request

## 📄 License

MIT License - see LICENSE file for details

---

## 🏃‍♂️ Quick Start Commands

```bash
# Development
make dev-start          # Start all services in development mode
make dev-stop           # Stop all services
make test-all           # Run all tests
make lint-all           # Run linters

# Production
make build-images       # Build all Docker images
make deploy-prod        # Deploy to production
make backup-db          # Backup all databases
```

This project demonstrates enterprise-level microservices architecture with real-world patterns including circuit breakers, retry mechanisms, event sourcing, CQRS, and comprehensive observability.
