#!/usr/bin/env bash

installOllama() {
  curl -fsSL https://ollama.com/install.sh | sh
}

installAIModel() {
  ollama pull deepseek-r1:1.5b
}

# Install Ollama if needed
installOllama

# Start Ollama service in the background
ollama serve > /dev/null 2>&1 &

# Install LLM model if needed
installAIModel

# Initialize and activate a Python virtual environment
python3 -m venv ai-analyzer
source ai-analyzer/bin/activate

# Download required libraries
pip install -r requirements.txt

# Activate the connector
python3 ai_api.py