# Aletheia

[![🚀 Build with Dependencies](https://github.com/hazardous-sun/aletheia/actions/workflows/go.yml/badge.svg)](https://github.com/hazardous-sun/aletheia/actions/workflows/go.yml)

> "If a lie is only printed often enough, it becomes a quasi-truth, and if such a truth is repeated often enough, it
> becomes an article of belief, a dogma, and men will die for it."
> - Isabella Jane Blagden

This project implements a comprehensive fact-checking platform developed in Go and Python. It helps users verify online
content by comparing it against information from reputable news sources through a combination of web crawling and AI
analysis. The system follows a distributed architecture with three specialized components working in concert.

## System Components

### [Client Application](client/README.md)

A cross-platform GUI built using Fyne framework that serves as the user interface. Key features:

- Configurable input fields for URLs, context prompts, and media types
- Robust error handling with color-coded logging
- Dynamic interface generation based on configuration
- Comprehensive API communication layer

### [Server API](server-api/README.md)

The core service built with Gin framework that orchestrates the fact-checking process:

- Manages relational data (languages/news outlets) in PostgreSQL
- Concurrent web crawling engine with configurable depth
- RESTful endpoints with comprehensive error handling
- Integrated with the AI Analyzer for content processing
- Detailed request/response logging

### [AI Analyzer](ai-analyzer/README.md)

The analytical component built with FastAPI and Ollama:

- Utilizes deepseek-r1:1.5b model for semantic analysis
- Advanced HTML processing and link extraction
- Content comparison algorithms
- Configurable API endpoints for integration
- Health checks and monitoring

## Getting Started

### Prerequisites

- Podman/Docker with compose support
- Go 1.23+ (for client and server development)
- Python 3.10+ (for AI analyzer)
- Ollama (automatically installed in container)

### Installation

1. Configure development environment:
   ```bash
   ./dev-setup.sh
   ```
2. Launch the full system:
   ```bash
   ./run.sh
   ```
   Customize ports and settings as needed:
   ```bash
   ./run.sh --SERVER_PORT=8080 --AI_PORT=9000 --DEBUG=true
   ```

## Operational Workflow

1. **Submission**: User provides content through the Client interface
2. **Processing**: Server API validates input and initiates crawlers
3. **Collection**: Concurrent crawlers gather relevant news content
4. **Analysis**: AI Analyzer compares submissions with collected data
5. **Reporting**: Results are compiled and returned to the user

## Configuration Management

The system is highly configurable through environment variables and runtime arguments:

| Component   | Key Variables          | Default Values             |
|-------------|------------------------|----------------------------|
| Server API  | `SERVER_PORT`, `DEBUG` | `8000`, `false`            |
| Database    | `DB_*` series          | PostgreSQL defaults        |
| AI Analyzer | `AI_PORT`, `MODEL`     | `7654`, `deepseek-r1:1.5b` |

## Development Practices

### Git Workflow

- Enforced commit message format via `commit-msg` hook
- Automatic testing before push via `pre-push` hook
- Changelog generation using git-cliff

### Testing

Run the complete test suite:

```bash
go test ./...
```

### Documentation

- Architectural diagrams in `docs/diagrams/`
- Component-specific documentation:
    - [Client Details](client/README.md)
    - [Server API Specifications](server-api/README.md)
    - [AI Analyzer Implementation](ai-analyzer/README.md)

## Contributing Guidelines

1. Always run `dev-setup.sh` to configure local hooks
2. Follow the commit message convention for changelog generation
3. Document new endpoints or features in the respective READMEs
4. Update the changelog:
   ```bash
   git-cliff -o CHANGELOG.md
   ```

## Roadmap

### Immediate Priorities

- Enhanced crawler reliability and error recovery
- News outlet credibility validation

### Future Development

- JWT-based authentication system
- Rate limiting and API quotas
- Comprehensive monitoring endpoints
- Swagger/OpenAPI documentation
- Advanced content similarity algorithms

For detailed implementation specifics, please refer to each component's dedicated documentation. The system is designed
for extensibility, with clear separation between the presentation layer (Client), business logic (Server API), and
analytical processing (AI Analyzer).
