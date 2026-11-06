# 🍽️ LetsEat -- Restaurant Management Backend API

**LetsEat** is a RESTful backend API designed to manage restaurant
information, food categories, and menu items. The project is built using
Go (Golang), Echo Framework, and PostgreSQL, following a layered
architecture to maintain clean, modular, and maintainable code.

The primary goal of this project is to develop a scalable backend system
that handles restaurant management and food menu operations efficiently.

------------------------------------------------------------------------

## 🚀 Technology Stack

  Technology            Purpose
  --------------------- ------------------------------------------
  Go (Golang) 1.24      Backend programming language
  Echo v4               HTTP web framework
  PostgreSQL 17         Relational database
  pgx / pgxpool v5      PostgreSQL driver and connection pooling
  Docker                Containerization
  Docker Compose        Multi-container management
  Postman               API development and testing
  Git & GitHub/GitLab   Version control

------------------------------------------------------------------------

## ✨ Features & Functionalities

### 1. Restaurant Management

-   Create a new restaurant.
-   Retrieve restaurant information.
-   Retrieve individual restaurant details.
-   Update restaurant information.
-   Delete restaurants.
-   Manage restaurant active/inactive status.
-   Support restaurants with the same name at different locations.

### 2. Food Category Management

-   Create food categories.
-   Retrieve all categories.
-   Retrieve individual category details.
-   Update category information.
-   Delete categories.

### 3. Menu Item Management

-   Create food menu items.
-   Retrieve all menu items.
-   Retrieve individual menu item details.
-   Update menu item information.
-   Update food item prices.
-   Manage food item availability.
-   Delete menu items.

### 4. Search Functionality

-   Search for restaurants and food items.
-   Filter search results based on supported search parameters.

### 5. Health Check

-   API health endpoint to verify server availability.

------------------------------------------------------------------------

## 🏗️ Project Architecture

The project follows a layered architecture inspired by Clean
Architecture principles.

**Request flow:**

`Client → Echo Router → Handler → Service → Repository → PostgreSQL`

Each layer has a separate responsibility:

  Layer          Responsibility
  -------------- ----------------------------------------
  Routes         Defines API endpoints and HTTP methods
  Handlers       Handles HTTP requests and responses
  Services       Contains business logic and validation
  Repositories   Handles database operations
  Models         Defines application data structures
  Database       Manages PostgreSQL connections

This separation makes the application easier to maintain, test, and
extend.

------------------------------------------------------------------------

## 📁 Project Structure

``` text
LetsEat/
│
├── cmd/
│   └── api/
│       └── main.go
│
├── internal/
│   ├── config/
│   ├── database/
│   ├── handlers/
│   ├── models/
│   ├── repositories/
│   ├── routes/
│   └── services/
│
├── migrations/
│   ├── 001_init.sql
│   └── 002_menu_management.sql
│
├── postman/
│
├── tests/
│
├── .env
├── .env.example
├── .gitignore
├── docker-compose.yml
├── go.mod
├── go.sum
└── README.md
```

------------------------------------------------------------------------

## ⚙️ Getting Started

### Prerequisites

Make sure you have installed:

-   Go 1.24 or later
-   Docker and Docker Compose
-   Git
-   Postman

### 1. Clone the Repository

``` bash
git clone https://github.com/Barrun-Kabir-Rifat/LetsEat.git
cd LetsEat
```

### 2. Configure Environment Variables

Create your environment file:

``` bash
cp .env.example .env
```

Configure the database and server settings:

``` env
DB_HOST=localhost
DB_PORT=5433
DB_USER=letseat
DB_PASSWORD=your_password
DB_NAME=letseat_db
DB_SSLMODE=disable

SERVER_PORT=8080
```

**Note:** Never commit your actual `.env` file or database credentials
to GitHub.

### 3. Start PostgreSQL Using Docker Compose

``` bash
docker compose up -d
```

Check running containers:

``` bash
docker compose ps
```

### 4. Run Database Migrations

Apply the SQL migration files in order:

1.  `migrations/001_init.sql`
2.  `migrations/002_menu_management.sql`

### 5. Install Go Dependencies

``` bash
go mod tidy
```

### 6. Run the Application

``` bash
go run ./cmd/api
```

The API server will be available at:

``` text
http://localhost:8080
```

------------------------------------------------------------------------

## 🐳 Docker

LetsEat uses Docker to run PostgreSQL in an isolated container
environment.

The PostgreSQL configuration includes:

-   PostgreSQL 17 image.
-   Persistent database storage using Docker volumes.
-   Environment-based database configuration.
-   Health checks using `pg_isready`.
-   Port mapping from host port `5433` to container port `5432`.

Docker Compose is used to manage the container lifecycle.

------------------------------------------------------------------------

## 🔌 API Health Check

**Endpoint:**

``` http
GET /api/v1/health
```

**Base URL:**

``` text
http://localhost:8080
```

Use Postman to send requests and inspect API responses.

------------------------------------------------------------------------

## 🧪 API Testing

Postman is used for testing the REST API.

The project includes testing for:

-   Restaurant CRUD operations.
-   Category CRUD operations.
-   Menu item CRUD operations.
-   Price and availability updates.
-   Search functionality.
-   API health checks.

------------------------------------------------------------------------

## 🗄️ Database

LetsEat uses PostgreSQL as its primary relational database.

The database currently includes the following major entities:

  Entity        Description
  ------------- ----------------------------------------------------
  Restaurants   Stores restaurant information
  Categories    Stores food categories
  Menu Items    Stores food item details, prices, and availability

Database migrations are maintained in the `migrations/` directory.

------------------------------------------------------------------------

## 🛠️ Development Commands

Run the application:

``` bash
go run ./cmd/api
```

Format Go source files:

``` bash
go fmt ./...
```

Update dependencies:

``` bash
go mod tidy
```

Run tests:

``` bash
go test ./...
```

Check Docker containers:

``` bash
docker compose ps
```

Stop Docker containers:

``` bash
docker compose down
```

Stop containers and remove database volumes (this deletes persisted
database data):

``` bash
docker compose down -v
```

------------------------------------------------------------------------

## 🔮 Future Improvements

Potential improvements planned for future development:

-   User authentication and authorization using JWT.
-   Role-based access control.
-   Restaurant owner and customer management.
-   Order management.
-   Food cart functionality.
-   Payment integration.
-   Advanced search and filtering.
-   API documentation using Swagger/OpenAPI.
-   Automated unit and integration testing.
-   Production-ready Docker deployment.

------------------------------------------------------------------------

## 👨‍💻 Author

**Barrun Kabir Rifat**

-   GitHub: [Barrun-Kabir-Rifat](https://github.com/Barrun-Kabir-Rifat)
-   Project: [LetsEat](https://github.com/Barrun-Kabir-Rifat/LetsEat)

------------------------------------------------------------------------

*LetsEat -- Building a structured and scalable restaurant management
backend with Go.*
