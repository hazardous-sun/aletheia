#!/usr/bin/env bash

# Start Ollama service in the background
echo "Starting Ollama service..."
ollama serve > /dev/null 2>&1 &

# Wait briefly to ensure Ollama is ready
sleep 2

# Start FastAPI application
echo "Starting FastAPI application..."
exec uvicorn ai_api:app --host 0.0.0.0 --port 7654