FROM docker.io/ollama/ollama:0.5.7

# Create app directory
WORKDIR /ai-analyzer

# Install dependencies
RUN apt-get update && \
    apt-get install -y \
    python3 \
    python3.10-venv \
    curl && \
    rm -rf /var/lib/apt/lists/*

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
ENTRYPOINT ["./run.sh"]