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
├── src/                    # Source code directory
│   ├── auth/              # Authentication related code
│   │   └── jwt.go        # JWT implementation
│   ├── handlers/          # HTTP request handlers
│   │   └── user.go       # User-related handlers
│   ├── models/           # Data models
│   │   └── user.go       # User model definition
│   ├── services/         # Business logic layer
│   │   └── user.go       # User-related services
│   ├── utils/            # Utility functions
│   │   └── response.go   # HTTP response helpers
│   └── validation/       # Input validation
│       └── user.go       # User-related validations
├── docker/                # Docker related files
│   ├── dev/              # Development environment
│   └── prod/             # Production environment
├── .dockerignore         # Docker ignore file
├── Dockerfile            # Main Docker build file
├── docker-compose.yml    # Docker compose configuration
├── .gitignore            # Git ignore file
├── go.mod                # Go module file
├── go.sum                # Go module checksum
├── main.go               # Application entry point
└── README.md             # Project documentation
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
