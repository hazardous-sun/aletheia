FROM docker.io/postgres:12

# Set environment variables
ENV POSTGRES_USER={$DB_USER}
ENV POSTGRES_PASSWORD=${DB_PASSWORD}
ENV POSTGRES_DB=${DB_NAME}

# Copy the initialization script to the Docker image
COPY initialize_db.sql /docker-entrypoint-initdb.d/