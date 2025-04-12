# Server API

A REST API built using the Gin framework in Go for managing languages and news outlets, with web crawling capabilities
to collect news data.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
    - [Environment Variables](#environment-variables)
    - [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
    - [Languages](#languages)
    - [News Outlets](#news-outlets)
    - [Crawlers](#crawlers)
- [Project Structure](#project-structure)
- [Database](#database)
- [Testing](#testing)
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

- **Containerized Environment**:
    - Docker/Podman setup for easy deployment
    - PostgreSQL database integration

## Prerequisites

Before you begin, ensure you have the following installed:

- [Podman](https://podman.io/) or Docker
- [Podman Compose](https://github.com/containers/podman-compose/) or Docker Compose
- [Go](https://golang.org/doc/install) (for local development)

## Getting Started

### Environment Variables

The application requires these environment variables:

| Variable    | Description              | Default Value |
|-------------|--------------------------|---------------|
| DB_HOST     | PostgreSQL database host | `news-db`     |
| DB_PORT     | PostgreSQL database port | `5432`        |
| DB_USER     | PostgreSQL username      | `postgres`    |
| DB_PASSWORD | PostgreSQL password      | `1234`        |
| DB_NAME     | PostgreSQL database name | `postgres`    |
| SERVER_PORT | Port for the API server  | `8000`        |

### Running the Application

1. Make the run script executable:
   bash
   chmod +x run.sh

2. Start the application:
   bash
   ./run.sh

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

## API Endpoints

### Languages

- **Create Language**:

  POST /language

  Request Body:
  json
  {
  "name": "english"
  }

- **List Languages**:

  GET /languages

- **Get Language by ID**:

  GET /languageId/:languageId

- **Get Language by Name**:

  GET /languageName/:languageName

### News Outlets

- **Create News Outlet**:

  POST /newsOutlet

  Request Body:
  json
  {
  "Name": "Example News",
  "QueryUrl": "https://example.com/search?q=KEYWORDS_HERE",
  "HtmlSelector": ".article a",
  "language": "english",
  "credibility": 80
  }

- **List News Outlets**:

  GET /newsOutlets

- **Get Outlet by ID**:

  GET /newsOutletId/:newsOutletId

- **Get Outlet by Name**:

  GET /newsOutletName/:newsOutletName

### Crawlers

- **Start Crawling**:

  POST /crawl

  Request Body:
  json
  {
  "pagesToVisit": 5,
  "query": "latest news"
  }

## Project Structure

The project follows a layered architecture:

src/
├── cmd/ # Main application entry point
├── controllers/ # HTTP request handlers
├── db/ # Database connection and configuration
├── deployments/ # Container deployment files
├── errors/ # Custom error definitions
├── models/ # Data structures
├── repositories/ # Database interaction layer
└── usecases/ # Business logic

## Database

The PostgreSQL database is initialized with these tables:

sql
CREATE TABLE languages (
Id SERIAL PRIMARY KEY,
Name VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE news_outlet (
Id SERIAL PRIMARY KEY,
Name VARCHAR(255) UNIQUE NOT NULL,
QueryUrl TEXT NOT NULL,
HtmlSelector TEXT NOT NULL,
LanguageId INT NOT NULL,
Credibility INT NOT NULL,
FOREIGN KEY (LanguageId) REFERENCES languages (Id)
ON UPDATE CASCADE ON DELETE CASCADE
);

## Testing

The project includes comprehensive tests for:

- Error handling
- Data models
- API responses

Run tests with:
bash
go test ./...

## TODO

- Implement security measures for Gin to block proxy communication
- Develop authentication for API endpoints
- Improve crawler repository to correctly collect website data
- Add more comprehensive error handling
- Implement rate limiting
- Add API documentation (Swagger/OpenAPI)