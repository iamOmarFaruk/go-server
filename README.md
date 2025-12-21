# Go Web App

A simple web application built with Go, demonstrating routing, controllers, templates, and Docker deployment.

## Features

- **Go 1.21** with Chi router
- **MVC Architecture** with proper separation of concerns
- **HTML Templates** with layout inheritance
- **Responsive Design** with modern CSS
- **Docker Support** for containerized deployment
- **Health Check** endpoint
- **Static File Serving** for CSS, JS, and images

## Project Structure

```
go-server/
├── cmd/
│   └── server/
│       └── main.go              # Application entry point
├── internal/
│   ├── handlers/                # HTTP handlers (controllers)
│   ├── services/               # Business logic
│   └── models/                 # Data structures
├── templates/                  # HTML templates
├── static/                     # Static files (CSS, JS, images)
├── Dockerfile                  # Container configuration
├── docker-compose.yml          # Development workflow
└── go.mod                      # Go module definition
```

## Quick Start

### Using Docker (Recommended)

1. **Build and run with Docker Compose:**
   ```bash
   docker-compose up --build
   ```

2. **Access the application:**
   - Open your browser and visit: http://localhost:8080

3. **Stop the application:**
   ```bash
   docker-compose down
   ```

### Manual Setup (Requires Go 1.21+)

1. **Install dependencies:**
   ```bash
   go mod download
   ```

2. **Run the application:**
   ```bash
   go run cmd/server/main.go
   ```

3. **Access the application:**
   - Open your browser and visit: http://localhost:8080

## Available Routes

- `/` - Home page with hero section and features
- `/about` - About page with company info and team
- `/testimonials` - Testimonials page with client stories
- `/health` - Health check endpoint
- `/static/*` - Static files (CSS, JS, images)

## Development

### Adding New Pages

1. **Create Model:** Add data structures in `internal/models/page.go`
2. **Create Service:** Add business logic in `internal/services/page_service.go`
3. **Create Handler:** Add HTTP handler in `internal/handlers/page_handler.go`
4. **Create Template:** Add HTML template in `templates/`
5. **Add Route:** Register the route in `cmd/server/main.go`

### Hot Reload (Development)

For hot reload during development, uncomment the volume mounts in `docker-compose.yml` and rebuild:

```yaml
volumes:
  - .:/app
  - /app/vendor
```

## Technologies Used

- **Go 1.21** - Programming language
- **Chi** - HTTP router
- **html/template** - Template engine
- **Docker** - Containerization
- **Responsive CSS** - Modern styling

## Deployment

This application is designed for easy deployment on platforms like:
- **Railway** - "Vercel for Go" with auto-deployment
- **Render** - Simple PaaS with free tier
- **Fly.io** - Edge-focused platform
- **Any Docker-compatible platform**

## Health Check

The application includes a health check endpoint at `/health` that returns HTTP 200 status, making it suitable for load balancers and container orchestrators.

## Security

- Runs as non-root user in Docker container
- Includes security headers via middleware
- Input validation and proper error handling
- Static files served with proper headers

## Performance

- Built with Go's performance and concurrency
- Compression middleware for faster load times
- Optimized Docker image with multi-stage build
- Minimal runtime dependencies