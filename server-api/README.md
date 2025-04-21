# Aletheia Server API

A REST API built using the Gin framework in Go for managing languages and news outlets, with web crawling capabilities to collect news data and integrate with an AI analyzer.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
  - [Environment Variables](#environment-variables)
  - [Running the Application](#running-the-application)
  - [Debugging the Application](#debugging-the-application)
    - [For JetBrains Users](#for-jetbrains-users)
    - [For VS Code Users](#for-vs-code-users)
- [API Endpoints](#api-endpoints)
  - [Languages](#languages)
  - [News Outlets](#news-outlets)
  - [Crawlers](#crawlers)
- [Project Structure](#project-structure)
- [Database](#database)
- [Testing](#testing)
- [Architecture](#architecture)
- [TODO](#todo)

## Features

- **Language Management**:
  - Add new languages
  - List all languages
  - Retrieve language by ID or name

- **News Outlet Management**:
  - Add news outlets with associated language
  - List all news outlets
  - Retrieve outlet by ID or name
  - Store credibility scores

- **Web Crawling**:
  - Crawl news outlets using configured query URLs and HTML selectors
  - Store crawled page bodies for analysis
  - Integration with AI analyzer service for link extraction
  - Concurrent crawling with configurable page limits

- **Error Handling**:
  - Comprehensive error logging with different levels (info, warning, error)
  - Color-coded console output for different log levels
  - Specific error types for different components

- **Containerized Environment**:
  - Docker/Podman setup for easy deployment
  - PostgreSQL database integration
  - Debugging support with Delve
  - Database initialization scripts

## Prerequisites

Before you begin, ensure you have the following installed:

- [Podman](https://podman.io/) or Docker
- [Podman Compose](https://github.com/containers/podman-compose/) or Docker Compose
- [Go](https://golang.org/doc/install) (1.23+ for local development)

## Getting Started

### Environment Variables

The application requires these environment variables:

| Variable        | Description                          | Default Value  |
|-----------------|--------------------------------------|----------------|
| DB_HOST         | PostgreSQL database host             | `news-db`      |
| DB_PORT         | PostgreSQL database port             | `5432`         |
| DB_USER         | PostgreSQL username                  | `postgres`     |
| DB_PASSWORD     | PostgreSQL password                  | `1234`         |
| DB_NAME         | PostgreSQL database name             | `postgres`     |
| SERVER_PORT     | Port for the API server              | `8000`         |
| AI_ANALYZER_URL | URL for the AI analyzer service      | `http://localhost:7654` |

### Running the Application

1. Make the run script executable:
   ```bash
   chmod +x run.sh
   ```

2. Start the application:
   ```bash
   ./run.sh
   ```

The script will:
1. Set environment variables
2. Clean up previous containers
3. Build and start the containers
4. The API will be available at http://localhost:8000

#### Script Options:
- `-C` or `--CLEAR`: Deletes the `pgdata/` volume
- `-R` or `--RESET`: Deletes project images before initialization
- `--DB_*`: Override specific database connection parameters
- `--SERVER_PORT`: Override the server port
- `--AI_ANALYZER_URL`: Override the AI analyzer service URL

### Debugging the Application

#### For JetBrains Users
1. The debug configuration is already included in `.idea/runConfigurations/`
2. Just run the "Remote Debug" configuration after starting the containers

#### For VS Code Users
1. Install recommended extensions
2. Two debug configurations are available:
- **Attach to Podman Container**: Attaches to a running container
- **Launch in Podman Container**: Builds and starts containers before debugging

## API Endpoints

### Languages

- **Create Language**:
  ```
  POST /language
  ```
  Request Body:
  ```json
  {
    "name": "english"
  }
  ```

- **List Languages**:
  ```
  GET /languages
  ```

- **Get Language by ID**:
  ```
  GET /languageId/:languageId
  ```

- **Get Language by Name**:
  ```
  GET /languageName/:languageName
  ```

### News Outlets

- **Create News Outlet**:
  ```
  POST /newsOutlet
  ```
  Request Body:
  ```json
  {
    "Name": "Example News",
    "QueryUrl": "https://example.com/search?q=QUERY_HERE",
    "HtmlSelector": ".article a",
    "language": "english",
    "credibility": 80
  }
  ```

- **List News Outlets**:
  ```
  GET /newsOutlets
  ```

- **Get Outlet by ID**:
  ```
  GET /newsOutletId/:newsOutletId
  ```

- **Get Outlet by Name**:
  ```
  GET /newsOutletName/:newsOutletName
  ```

### Crawlers

- **Start Crawling**:
  ```
  POST /crawl
  ```
  Request Body:
  ```json
  {
    "pagesToVisit": 5,
    "query": "latest news"
  }
  ```

## Project Structure

The project follows a clean architecture pattern with clear separation of concerns:

```
src/
├── cmd/               # Entry point (main.go)
├── controllers/       # HTTP request handlers
├── db/                # Database connection and configuration
├── deployments/       # Container deployment files
├── errors/            # Custom error definitions and logging
├── models/            # Data structures and business objects
├── repositories/      # Database interaction layer
└── usecases/          # Business logic
```

## Database

The PostgreSQL database is initialized with these tables:

```sql
CREATE TABLE languages (
    Id   SERIAL PRIMARY KEY,
    Name VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE news_outlet (
    Id           SERIAL PRIMARY KEY,
    Name         VARCHAR(255) UNIQUE NOT NULL,
    QueryUrl     TEXT                NOT NULL,
    HtmlSelector TEXT                NOT NULL,
    LanguageId   INT                 NOT NULL,
    Credibility  INT                 NOT NULL,
    FOREIGN KEY (LanguageId) REFERENCES languages (Id) 
    ON UPDATE CASCADE ON DELETE CASCADE
);
```

## Testing

The project includes comprehensive tests for:

- Error handling and constants
- Data models and JSON serialization
- API response structures
- Database repository layer

Run tests with:
```bash
go test ./...
```

## Architecture

The application follows a layered architecture:

1. **Controllers**: Handle HTTP requests/responses
2. **Use Cases**: Contain business logic
3. **Repositories**: Handle data persistence
4. **Models**: Define data structures
5. **Errors**: Centralized error handling

Key design patterns:
- Dependency injection
- Separation of concerns
- Concurrent web crawling
- Comprehensive logging

## TODO

- High priority:
  - Improve crawler repository to correctly collect website data
  - Add validation for news outlet credibility scores (0-100)

- Medium priority:
  - Add API documentation (Swagger/OpenAPI)
  - Refactor the code to add more comprehensive error handling
  - Implement retry logic for failed crawls

- Low priority:
  - Implement security measures for Gin to block proxy communication
  - Develop authentication for API endpoints
  - Implement rate limiting
  - Add health check endpoints
  - Implement proper shutdown handling
