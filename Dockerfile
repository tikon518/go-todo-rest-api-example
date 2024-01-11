# Use a minimal base image
FROM golang:1.16-alpine

# Set the working directory
WORKDIR /app

# Copy the source code to the container
COPY . .

# Install the required dependencies
RUN apk add --no-cache git mysql-client && \
    go get github.com/go-sql-driver/mysql && \
    go get github.com/gorilla/mux && \
    go get github.com/joho/godotenv && \
    go mod download

# Build the Go binary
RUN go build -o main .

# Expose port 8080 to the host
EXPOSE 8080

# Set the environment variables for the database connection
ENV DB_DIALECT=mysql \
    DB_HOST=103.195.49.199 \
    DB_PORT=5589 \
    DB_USER=dbadmin \
    DB_PASSWORD=aXa#iMH*US%*Jzfkr \
    DB_NAME=shopify \
    DB_CHARSET=utf8mb4

# Start the application
CMD ["./main"]
