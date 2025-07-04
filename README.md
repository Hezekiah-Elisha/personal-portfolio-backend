# Personal Portfolio Backend

A RESTful API backend for a personal portfolio application built with Go, Gin framework, and MySQL database.

## Features

- **User Authentication**: JWT-based authentication with registration and login
- **User Management**: CRUD operations for users with role-based access
- **Experience Management**: Manage work experience entries
- **CORS Support**: Cross-origin resource sharing configuration
- **Database Migration**: Automatic database schema migration with GORM
- **Password Security**: Bcrypt password hashing
- **Environment Configuration**: Environment-based configuration management

## Tech Stack

- **Language**: Go
- **Web Framework**: Gin
- **Database**: MySQL
- **ORM**: GORM
- **Authentication**: JWT (JSON Web Tokens)
- **Password Hashing**: Bcrypt
- **Configuration**: Viper + godotenv

## Project Structure

```
personal-portfolio-backend/
├── config/
│   └── config.go              # Database configuration
├── controllers/
│   ├── auth.controller.go     # Authentication handlers
│   ├── user.controller.go     # User management handlers
│   └── experience.controller.go # Experience management handlers
├── middlewares/
│   └── auth.middleware.go     # JWT authentication middleware
├── models/
│   ├── user.model.go          # User data model
│   └── experience.model.go    # Experience data model
├── routes/
│   ├── auth.route.go          # Authentication routes
│   ├── user.route.go          # User routes
│   └── experience.route.go    # Experience routes
├── main.go                    # Application entry point
└── .env                       # Environment variables
```

## Installation

1. **Clone the repository**

   ```bash
   git clone <repository-url>
   cd personal-portfolio-backend
   ```

2. **Install dependencies**

   ```bash
   go mod tidy
   ```

3. **Set up environment variables**
   Create a `.env` file in the root directory:

   ```env
   DB_USER=your_db_user
   DB_PASSWORD=your_db_password
   DB_HOST=localhost
   DB_PORT=3306
   DB_NAME=portfolio_db
   JWT_SECRET=your_jwt_secret_key
   GIN_MODE=debug
   ```

4. **Set up MySQL database**

   ```sql
   CREATE DATABASE portfolio_db;
   ```

5. **Run the application**
   ```bash
   go run main.go
   ```

The server will start on `http://localhost:8081`

## API Endpoints

### Authentication

- `POST /auth/register` - Register a new user
- `POST /auth/login` - Login user

### Users

- `GET /users/` - Get all users
- `POST /users/` - Create a new user (requires authentication)

### Experiences

- `GET /experiences/` - Get all experiences
- `POST /experiences/` - Create a new experience (requires authentication)

## Request/Response Examples

### Register User

```bash
POST /auth/register
Content-Type: application/json

{
  "username": "johndoe",
  "email": "john@example.com",
  "password": "password123"
}
```

### Login

```bash
POST /auth/login
Content-Type: application/json

{
  "email": "john@example.com",
  "password": "password123"
}
```

### Create Experience

```bash
POST /experiences/
Authorization: Bearer <jwt_token>
Content-Type: application/json

{
  "title": "Software Developer",
  "description": "Developed web applications using Go and React",
  "company": "Tech Corp",
  "location": "San Francisco, CA",
  "start_date": "2023-01-01",
  "end_date": "2023-12-31",
  "user_id": 1
}
```

## Environment Variables

| Variable      | Description                   | Example           |
| ------------- | ----------------------------- | ----------------- |
| `DB_USER`     | MySQL username                | `root`            |
| `DB_PASSWORD` | MySQL password                | `password`        |
| `DB_HOST`     | MySQL host                    | `localhost`       |
| `DB_PORT`     | MySQL port                    | `3306`            |
| `DB_NAME`     | Database name                 | `portfolio_db`    |
| `JWT_SECRET`  | JWT signing secret            | `your-secret-key` |
| `GIN_MODE`    | Gin mode (debug/release/test) | `debug`           |

## Development

### Running in Development Mode

```bash
export GIN_MODE=debug
go run main.go
```

### Building for Production

```bash
export GIN_MODE=release
go build -o portfolio-backend
./portfolio-backend
```

### Database Migration

The application automatically runs database migrations on startup using GORM's `AutoMigrate` feature.

## CORS Configuration

The application is configured to accept requests from:

- `http://localhost` (for local development)
- `https://github.com` (example external origin)

Modify the CORS configuration in `main.go` to suit your frontend requirements.

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Important Resources

- [Bcrypt](https://mojoauth.com/hashing/bcrypt-in-go/) - Password hashing in Go using Bcrypt
- [jwt-token](https://medium.com/@cheickzida/golang-implementing-jwt-token-authentication-bba9bfd84d60) - JWT token authentication in Go

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
