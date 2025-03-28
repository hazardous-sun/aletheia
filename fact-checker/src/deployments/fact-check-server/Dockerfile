FROM docker.io/golang:1.23-alpine as builder

WORKDIR /go/build

COPY ../../. .

RUN go mod download

RUN go build -o fact-check-server src/cmd/main.go

#-----------------------------------------------------------------------------------------------------------------------

FROM alpine:latest

ENV DB_HOST=${DB_HOST}
ENV DB_PORT=${DB_PORT}

# Install netcat (nc) for the wait-for-db.sh script
RUN apk add --no-cache netcat-openbsd

# Copy the wait-for-db.sh script and set permissions
COPY src/deployments/fact-check-server/wait-for-db.sh /wait-for-db.sh
RUN chmod +x /wait-for-db.sh

# Copy the binary from the builder stage
COPY --from=builder /go/build/fact-check-server /fact-check-server

EXPOSE 8000

CMD ["/wait-for-db.sh", "$DB_HOST", "$DB_PORT", "./fact-check-server"]