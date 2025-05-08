# Go Builder Plate

## Project Description

Go Builder Plate is a starter template for Go projects. It provides a structured foundation to quickly set up and start developing Go applications.

## Features

### Core Features

- RESTful API using Echo Framework
- JWT Authentication
- Request validation
- Standardized response format
- Structured error handling

### Database Integration

- PostgreSQL support
  - GORM implementation for ORM
  - Raw SQL queries implementation
  - Connection pooling
  - Migration support

### Payment Integration

- Multiple payment gateway support
  - Midtrans
  - Xendit
  - Stripe
- Payment status webhooks
- Transaction history
- Payment method management

### Docker Support

- Multi-stage build optimization
- Docker Compose for development
- Production-ready container
- Environment variables management
- Container health checks

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/hnifmaghfur/go-builder-plate.git
   ```
2. Navigate to the project directory:
   ```sh
   cd go-builder-plate
   ```
3. Install dependencies:
   ```sh
   go mod tidy
   ```

## Usage

1. Build the project:
   ```sh
   go build
   ```
2. Run the project:
   ```sh
   ./go-builder-plate
   ```

## Folder Structure

```
go-builder-plate/
├── cmd/                      # Main application entry points
│   |── server/               # Server application
│   |   └── main.go          # Entry point
│   └── docker/               # Docker related files
│       └── DockerFile        # Dockerfile for development environment
├── internal/                 # Private application code
│   ├── auth/                 # Authentication related code
│   │   └── jwt.go            # JWT implementation
│   ├── handler/              # HTTP request handlers
│   │   └── http/             # HTTP handlers
│   │       └── v1/           # API version 1
│   │           └── user/     # User-related handlers
│   ├── model/                # Data models
│   │   └── user.go           # User model definition
│   ├── repository/           # Data access layer
│   │   └── user.go           # User repository
│   ├── service/              # Business logic layer
│   │   └── user.go           # User-related services
│   └── middleware/           # HTTP middleware
│       └── auth.go           # Authentication middleware
├── pkg/                      # Library code that can be used by external applications
│   └── utils/                # Utility functions
│       └── response.go       # HTTP response helpers
├── api/                      # API contract files (OpenAPI/Swagger)
│   └── openapi/              # OpenAPI specifications
├── configs/                  # Configuration files
│   ├── app.yaml              # Application configuration
│   └── database.yaml         # Database configuration
├── migrations/               # Database migrations
├── scripts/                  # Build and utility scripts
├── .gitignore                # Git ignore file
├── go.mod                    # Go module file
├── docker-compose.yaml       # Docker compose file
├── go.sum                    # Go module checksum
├── Makefile                  # Common commands
└── README.md                 # Project documentation
```

## Contribution Guidelines

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -m 'Add some feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Open a pull request.

## Contributors

- [hnifmaghfur](https://github.com/hnifmaghfur)

## License

This project is licensed under the MIT License.
