FROM docker.io/golang:1.23

# Create a directory that will be used to store the project files for the build step
WORKDIR /go/build

# Copy the source code
COPY ../.. .

# Copy the wait-for-db.sh script
COPY wait-for-db.sh /wait-for-db.sh
RUN chmod +x /wait-for-db.sh

# Expose the port that will be used for the API requests
EXPOSE 8000

# Build the Go app
RUN go build -o ai-fact-checker cmd/main.go

# Run the executable after waiting for the database to be ready
CMD ["/wait-for-db.sh", "news-db", "5432", "./ai-fact-checker"]