FROM golang:1.21.2-alpine

WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /product-catalog-service

EXPOSE 8080

# Run
CMD ["/product-catalog-service"]