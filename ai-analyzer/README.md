# AI Content Analyzer

A FastAPI-based service that analyzes user-submitted content against reputable news sources using Ollama's AI models.

## Features

- **Content Analysis**: Compares user posts with news content to identify alignment or contradictions
- **Link Extraction**: Extracts relevant news article links from HTML content with titles
- **AI-Powered**: Uses Ollama's `deepseek-r1:1.5b` model for semantic analysis
- **REST API**: Provides endpoints for easy integration with other applications
- **CORS Support**: Configured to accept requests from any origin
- **Large Request Handling**: Supports requests up to 10MB in size

## Prerequisites

- Docker or Podman
- Python 3.10+
- Ollama (will be installed automatically in container)

## Installation

### Using Docker/Podman

1. Build the container:
   ```bash
   podman build . -t ai-analyzer
   ```

2. Run the container (default port 7654):
   ```bash
   podman run -d --name ai-analyzer -p 7654:7654 ai-analyzer
   ```

   To use a different port:
   ```bash
   podman run -d --name ai-analyzer -p [YOUR_PORT]:7654 ai-analyzer
   ```

### Manual Installation

1. Install system dependencies:
   ```bash
   sudo apt-get update && sudo apt-get install -y curl pciutils
   ```

2. Install Ollama:
   ```bash
   curl -fsSL https://ollama.com/install.sh | sh
   ```

3. Clone the repository and install Python dependencies:
   ```bash
   cd src
   python3 -m venv venv
   source venv/bin/activate
   pip install -r requirements.txt
   ```

4. Start the service:
   ```bash
   ./run.sh
   ```

## API Endpoints

### `POST /getLinks`

Extracts news article links with titles from HTML content.

**Request:**
```json
{
  "html_content": "<html>...</html>"
}
```

**Response:**
```json
[
  {"title": "Article Title 1", "url": "https://example.com/article1"},
  {"title": "Article Title 2", "url": "https://example.com/article2"}
]
```

### `POST /analyze`

Analyzes post content against news sources.

**Request:**
```json
{
  "post_content": "User's original post content",
  "news_content": "Content from reputable news sources",
  "user_context": "Optional additional context"
}
```

**Response:**
```json
{
  "success": true,
  "analysis": "AI-generated analysis comparing the post to news sources"
}
```

## Configuration

- **Port**: Change the service port by modifying the `PORT` environment variable or using the `--PORT=` argument in `run.sh`
- **AI Model**: Default model is `deepseek-r1:1.5b`. Change this in `content_analyzer.py` and `link_extractor.py`
- **Request Size**: Configured for 10MB max request size in `ai_api.py`

## Project Structure

```
.
├── Dockerfile            # Container configuration
├── README.md            # Project documentation
├── run.sh               # Main execution script (container version)
└── src/
    ├── ai_api.py        # FastAPI application and endpoints
    ├── content_analyzer.py # Core content analysis logic
    ├── link_extractor.py  # HTML link extraction logic
    ├── requirements.txt # Python dependencies
    └── run.sh          # Setup and execution script for manual installation
```

## Error Handling

The API provides detailed error responses for:
- Missing or invalid input data
- JSON parsing errors
- Internal server errors
- Model availability issues

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.