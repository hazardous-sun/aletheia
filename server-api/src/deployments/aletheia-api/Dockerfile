FROM docker.io/golang:1.23-alpine AS builder

WORKDIR /go/build

COPY ../../. .

RUN go mod download

# Install Delve
RUN go install github.com/go-delve/delve/cmd/dlv@latest

RUN go build -o aletheia-api src/cmd/main.go

#-----------------------------------------------------------------------------------------------------------------------

FROM alpine:latest

ENV DB_HOST=${DB_HOST}
ENV DB_PORT=${DB_PORT}

# Install netcat (nc) for the wait-for-db.sh script
RUN apk add --no-cache netcat-openbsd

# Copy the wait-for-db.sh script and set permissions
COPY src/deployments/aletheia-api/wait-for-db.sh /wait-for-db.sh
RUN chmod +x /wait-for-db.sh

# Copy the binary from the builder stage
COPY --from=builder /go/build/aletheia-api /aletheia-api

# Copy Delve from the builder stage
COPY --from=builder /go/bin/dlv /dlv

# Add a non-root user for Delve
RUN adduser -D -g '' appuser
USER appuser

EXPOSE 8000 40000

CMD ["/wait-for-db.sh", "$DB_HOST", "$DB_PORT", "/dlv", "exec", "/aletheia-api", "--headless", "--listen=:40000", "--api-version=2", "--accept-multiclient"]