# Notes Sharing Application

A full-stack web application for sharing and managing notes, built with Go (backend) and Next.js (frontend).

## Project Structure

The project consists of two main components:

### Backend (`notes-sharing-app-be`)

- Built with Go
- RESTful API architecture
- Features:
  - User authentication
  - Note management (CRUD operations)
  - Image handling
  - JWT-based authentication
  - Database migrations

### Frontend (`notes-sharing-app-fe`)

- Built with Next.js and TypeScript
- Modern UI components using Shadcn UI
- Features:
  - Responsive design
  - User authentication
  - Note creation and management
  - Image upload support
  - Client-side routing

## Prerequisites

- Go 1.23 or later
- Node.js 18.x or later
- Docker and Docker Compose
- PostgreSQL (if running without Docker)

## Getting Started

1. Clone the repository:

```bash
git clone https://github.com/rizkyIqbal/notes-sharing-app.git
cd notes-sharing-app
```

2. Start the application using Docker Compose:

```bash
docker-compose up -d
```

This will start both the frontend and backend services.

### Manual Setup

#### Backend

1. Navigate to the backend directory:

```bash
cd notes-sharing-app-be
```

2. Install dependencies:

```bash
go mod download
```

3. Run the migrations:

```bash
# Ensure your database is running and configured
go run cmd/server/main.go
```

#### Frontend

1. Navigate to the frontend directory:

```bash
cd notes-sharing-app-fe
```

2. Install dependencies:

```bash
npm install
```

3. Start the development server:

```bash
npm run dev
```

## Features

- User Authentication

  - Register
  - Login
  - JWT-based session management

- Notes Management

  - Create, read, update, and delete notes
  - Image attachments support
  - Public and private notes
  - Rich text editing

- User Interface
  - Responsive design
  - Modern and clean UI
  - Mobile-friendly navigation

## Directory Structure

```
├── notes-sharing-app-be/     # Backend application
│   ├── cmd/                  # Application entry points
│   ├── internal/            # Internal packages
│   │   ├── config/         # Configuration
│   │   ├── guards/         # Middleware
│   │   ├── handler/        # HTTP handlers
│   │   ├── models/         # Data models
│   │   ├── repository/     # Data access layer
│   │   └── service/        # Business logic
│   ├── migrations/         # Database migrations
│   └── pkg/                # Shared packages
│
├── notes-sharing-app-fe/     # Frontend application
│   ├── src/
│   │   ├── app/           # Next.js pages
│   │   ├── components/    # React components
│   │   ├── helper/        # Helper functions
│   │   ├── hooks/         # Custom React hooks
│   │   ├── lib/           # Utilities and services
│   │   └── types/         # TypeScript types
│   └── public/            # Static assets
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
