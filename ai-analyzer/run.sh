#!/usr/bin/bash

installOllama() {
  if command -v ollama &> /dev/null; then
    echo "ollama is already installed"
  else
    echo "ollama is not installed, installing most recent version..."

    if curl -fsSL https://ollama.com/install.sh | sh > install.log 2>&1; then
      echo "ollama installed successfully"
    else
      echo "ollama installation failed!"
      exit 1
    fi

    echo "ollama successfully installed!"
  fi
}

installAIModel() {
  if [[ $(ollama list | grep deepseek) != 0 ]]; then
    echo "Deepseek LLM not installed. Downloading Deepseek 1.5b model..."

    if [[ $(ollama pull deepseek-r1:1.5b) != 0 ]]; then
      echo "An error occurred while downloading the LLM model"
    fi

    echo "Successfully installed the LLM!"
  fi
}

# Install Ollama if needed
installOllama

# Install LLM model if needed
installAIModel

# Initialize and activate a Python virtual environment
python -m venv ai-analyzer
source ai-ananlyzer/bin/activate

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