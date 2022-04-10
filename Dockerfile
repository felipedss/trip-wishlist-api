FROM golang:1.16.5 as development

# Add a work directory
WORKDIR /app

# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy app files
COPY . .

# Expose port
EXPOSE 8080

# Start app
CMD GO_ENVIRONMENT=test go run *.go
