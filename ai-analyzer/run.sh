#!/usr/bin/bash

installOllama() {
  _=$(ollama -v)

  if [[ "$?" != 0 ]]; then
    echo "ollama is not installed, installing most recent version..."
    _=$(curl -fsSL https://ollama.com/install.sh | sh)

    if [[ "$?" != 0 ]]; then
      echo "ollama installation failed!"
      exit 1
    fi

    echo "ollama successfully installed!"
  fi
}

instalAIModel() {
  _=$(ollama list | grep deepseek)
  if [[ $? != 0 ]]; then
    echo "Deepseek LLM not installed. Downloading Deepseek 1.5b model..."
    _=$(ollama pull deepseek-r1:1.5b)

    if [[ $? != 0 ]]; then
      echo "An error occurred while downloading the LLM model"
    fi

    echo "Successfully installed the LLM!"
  fi
}

# Install Ollama if needed
installOllama

# Install LLM model if needed
instalAIModel

# Activate Python virtual environment
source ./ai-analyzer/bin/activate

# Download required libraries
pip install -r requirements.txt

python3 connector.py