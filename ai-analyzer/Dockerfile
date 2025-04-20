FROM python:3.10-slim

# Create app directory
WORKDIR /ai-analyzer

# Install system dependencies
RUN apt-get update && \
    apt-get install -y \
    curl \
    pciutils \
    && \
    rm -rf /var/lib/apt/lists/*

# Install Ollama
RUN curl -fsSL https://ollama.com/install.sh | sh

# Start Ollama, wait for it to initialize, then pull the model
RUN ollama serve > /dev/null 2>&1 & \
    sleep 5 && \
    ollama pull deepseek-r1:1.5b

# Copy the project to the container
COPY src/ ./

# Install Python dependencies
RUN python3 -m venv venv && \
    . venv/bin/activate && \
    pip install --no-cache-dir -r requirements.txt

# Make run.sh executable
RUN chmod +x ./run.sh

# Expose the FastAPI port
EXPOSE 7654

# Start the service
ENTRYPOINT ["/bin/bash", "-c", "source venv/bin/activate && ./run.sh"]