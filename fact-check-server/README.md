# Fact Check Server

This section of the project is a REST API built using the Gin framework in Go.

## Table of Contents

- [TODO](#todo)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
    - [Environment Variables](#environment-variables)
    - [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
  - [Languages](#languages)
  - [News Outlets](#news-outlets)
- [Project Structure](#project-structure)
- [Docker Compose](#docker-compose)
- [Database Initialization](#database-initialization)

## TODO

- Apply security measures to Gin to block communication with proxies
- Develop authentication measures for the API calls

## Features

- Manage accepted languages
    - [x] Add languages
    - [x] List all languages
    - [x] Retrieve a language by ID
    - [x] Retrieve a language by name
- Manage accepted news outlets
    - [x] Add news outlets
    - [x] List all news outlets
    - [ ] Retrieve a news outlet by ID
    - [x] Retrieve a news outlet by name
- Dockerized environment for easy setup
- PostgreSQL database for data storage

## Prerequisites

Before you begin, ensure you have the following installed:

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Go](https://golang.org/doc/install) (optional, for local development)

## Getting Started

### Environment Variables

The application requires the following environment variables to be set:

- `DB_HOST`: The hostname of the PostgreSQL database (default: `news-db`).
- `DB_PORT`: The port of the PostgreSQL database (default: `5432`).
- `DB_USER`: The username for the PostgreSQL database (default: `postgres`).
- `DB_PASSWORD`: The password for the PostgreSQL database (default: `1234`).
- `DB_NAME`: The name of the PostgreSQL database (default: `postgres`).

These variables are set in the [`run-server.sh`](run-server.sh) script.

### Running the Application

```bash
# Make the run-server.sh script executable:
chmod +x run-server.sh

# Initialize the API server
./run-server.sh
```

#### The [run-server.sh](run-server.sh) script will:

1. Set the necessary environment variables.
2. Clean up any previously created containers.
3. Start the Docker containers using podman-compose.

Finally, the Go application will be available at http://localhost:8000

#### run-server.sh parameters

- `-d`: deletes the `pgdata/` volume used to store the languages and news outlets
- `-r`: deletes the images from the project before initializing the pod
- `--DB_HOST`: overwrites the DB_HOST environment variable
    - Example: `--DB_HOST=localhost`
- `--DB_PORT`: overwrites the DB_PORT environment variable
- `--DB_USER`: overwrites the DB_USER environment variable
- `--DB_PASSWORD`: overwrites the DB_PASSWORD environment variable
- `--DB_NAME`: overwrites the DB_NAME environment variable

## API Endpoints

### Languages

- Create a new language
    - `POST /language`
    - Request Body Example:
      ```json
      {
        "name": "english"
      }
      ```
    - Response Example:
      ```json
      {
        "id": 4,
        "name": "german"
      }
      ```
- List all languages
    - `GET /languages`
    - Response Example:
      ```json
      [
        {
          "id": 1,
          "name": "spanish"
        },
        {
        "id": 2,
          "name": "portuguese"
        },
        {
          "id": 3,
          "name": "english"
        },
        {
          "id": 4,
          "name": "german"
        }
      ]
      ```
- Retrieve a language by ID:
    - `GET /languageId/:languageId`
    - Response Example:
      ```json
      {
        "id": 3,
        "name": "english"
      }
      ```
- Retrieve a language by name:
    - `GET /languageName/:languageName`
    - Response Example:
      ```json
      {
        "id": 4,
        "name": "german"
      }
      ```

### News Outlets

- Create a new news outlet
    - `POST /newsOutlet`
    - Request Body Example:
      ```json
      {
        "name": "example",
        "url": "example.com",
        "language": "english",
        "credibility": 10
      }
      ```
    - Response Example:
      ```json
      {
        "id": 1,
        "name": "example",
        "url": "example.com",
        "language": "english",
        "credibility": 10
      }
      ```
- List all news outlets
    - `GET /newsOutlets`
    - Response Example:
      ```json
      [
        {
          "id": 1,
          "name": "example",
          "url": "example.com",
          "language": "english",
          "credibility": 10
        }
      ]
      ```
- Retrieve a news outlet by name:
    - `GET /newsOutletName/:newsOutletName`
    - Response Example:
      ```json
      {
        "id": 1,
        "name": "example",
        "url": "example.com",
        "language": "english",
        "credibility": 10
      }
      ```

## Project Structure

The project is structured into four layers:

1. Controller: Handles HTTP requests and responses.
2. Use Case: Contains the business logic.
3. Repository: Manages data access and interaction with the database.
4. Model: Defines the data models and database schema.

## Docker Compose

The `docker-compose.yml` file defines two services:

- **fact-check-server**: The Go application server.
- **news-db**: The PostgreSQL database.

The `fact-check-server` service depends on the `news-db` service, ensuring that the database is up and running before the
application starts.

## Database Initialization

```sql
CREATE TABLE languages
(
    ID   SERIAL PRIMARY KEY,
    NAME VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE news_outlet
(
    ID          SERIAL PRIMARY KEY,
    NAME        VARCHAR(255) UNIQUE NOT NULL,
    URL         TEXT                NOT NULL,
    LANGUAGE    INT                 NOT NULL,
    CREDIBILITY INT                 NOT NULL
);

ALTER TABLE news_outlet
    ADD CONSTRAINT fk_language
        FOREIGN KEY (language) REFERENCES languages (id);
```
