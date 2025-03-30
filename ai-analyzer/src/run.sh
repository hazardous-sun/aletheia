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

export ORIGINAL_POST_CONTENT="ducks do not 'quack', they only 'bark'"
echo $ORIGINAL_POST_CONTENT

export ONLINE_NEWS_CONTENT="ducks only 'quack' and nothing else"
echo $ONLINE_NEWS_CONTENT

export USER_CONTEXT="can you confirm to me if ducks actually 'bark'?"
echo $USER_CONTEXT

# Activate the connector
python3 connector.py