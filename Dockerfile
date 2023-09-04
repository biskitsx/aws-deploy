# syntax=docker/dockerfile:1

FROM golang:1.18-alpine


# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download


COPY ./ ./

# Build
RUN go build -o ./myapp


EXPOSE 8000

# Run
CMD ["./myapp"]