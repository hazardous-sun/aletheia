FROM docker.io/ollama/ollama:0.5.7

# Create app directory
WORKDIR /ai-analyzer
RUN apt-get update && apt-get install python3 -y && apt-get install python3.10-venv -y && apt-get install -y curl && rm -rf /var/lib/apt/lists/*

# Copy the project to the container
COPY src/* .
RUN chmod +x ./run.sh

ENTRYPOINT ["./run.sh"]