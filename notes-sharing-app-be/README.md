# Notes Sharing App Backend

The backend service for the Notes Sharing Application, built with Go. This service provides a RESTful API for note management, user authentication, and image handling.

## Tech Stack

- *Go* - Main programming language
- *PostgreSQL* - Database
- *JWT* - Authentication
- *Docker* - Containerization

## Project Structure


.
├── cmd/
│   └── server/          # Application entry point
│       └── main.go
├── internal/
│   ├── config/         # Application configuration
│   ├── guards/         # Middleware implementations
│   ├── handler/        # HTTP request handlers
│   ├── models/         # Data models
│   ├── repository/     # Data access layer
│   └── service/        # Business logic layer
├── migrations/         # Database migrations
└── pkg/               # Shared packages
    └── jwt/           # JWT utilities


## Getting Started

### Prerequisites

- Go 1.21 or later
- PostgreSQL
- Docker (optional)

### Local Development Setup

1. Clone the repository:

bash
git clone https://github.com/rizkyIqbal/notes-sharing-app.git
cd notes-sharing-app/notes-sharing-app-be


2. Install dependencies:

bash
go mod download


3. Set up the database:

- Create a PostgreSQL database
- Run the migrations in the migrations folder

4. Run the application:

bash
go run cmd/server/main.go


### Docker Setup

1. Build the Docker image:

bash
docker build -t notes-app-backend .


2. Run the container:

bash
docker run -p 8080:8080 notes-app-backend


## API Endpoints

### Authentication

- POST /register - Register a new user
- POST /login - User login
- POST /auth/refresh - Refresh access token

### Notes

- GET /notes - Get all notes
- GET /notes/{id} - Get a specific note
- POST /notes - Create a new note
- PUT /notes/{id} - Update a note
- DELETE /notes/{id} - Delete a note

### Images

- POST /notes/{id}/images - Upload note images
- GET /notes/{id}/images/ - Get an image

## Development

### Project Components

1. *Handlers*: Located in internal/handler/

   - Handle HTTP requests
   - Validate input
   - Call appropriate services

2. *Services*: Located in internal/service/

   - Implement business logic
   - Coordinate between different repositories
   - Handle data transformations

3. *Repositories*: Located in internal/repository/

   - Handle database operations
   - Implement data access patterns

4. *Models*: Located in internal/models/

   - Define data structures
   - Implement model-specific methods

5. *Middleware*: Located in internal/guards/
   - Authentication middleware
   - Logging middleware
   - JSON response middleware

### Testing

Run tests with:

bash
go test ./...


## Configuration

The application can be configured using environment variables:

env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=notes_db
JWT_SECRET=your-secret-key


## Error Handling

The application uses standard HTTP status codes:

- 200: Success
- 201: Created
- 400: Bad Request
- 401: Unauthorized
- 403: Forbidden
- 404: Not Found
- 500: Internal Server Error

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

This project is licensed under the MIT License.
