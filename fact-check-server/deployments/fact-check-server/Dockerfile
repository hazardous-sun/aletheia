FROM docker.io/golang:1.23-alpine as builder

WORKDIR /go/build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ai-fact-checker cmd/main.go

#-----------------------------------------------------------------------------------------------------------------------

FROM alpine:latest

# Install netcat (nc) for the wait-for-db.sh script
RUN apk add --no-cache netcat-openbsd

# Copy the wait-for-db.sh script and set permissions
COPY deployments/fact-check-server/wait-for-db.sh /wait-for-db.sh
RUN chmod +x /wait-for-db.sh

# Copy the binary from the builder stage
COPY --from=builder /go/build/ai-fact-checker /ai-fact-checker

EXPOSE 8000

CMD ["/wait-for-db.sh", "news-db", "5432", "./ai-fact-checker"]