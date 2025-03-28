#!/usr/bin/bash

installOllama() {
  if [[ $(ollama -v) != 0 ]]; then
    echo "ollama is not installed, installing most recent version..."

    if [[ $(curl -fsSL https://ollama.com/install.sh | sh) != 0 ]]; then
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

# Activate Python virtual environment
source ai-ananlyzer/bin/activate

# Download required libraries
pip install -r requirements.txt

# Activate the connector
python3 connector.py