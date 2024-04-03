FROM golang:1.22.1 AS builder

# Set working directory
WORKDIR /build

# Copy files
COPY . .

# Download Go Modules
RUN go mod download

# Build Go App
RUN go build -o main ./cmd

# Set Distro Image
FROM gcr.io/distroless/base-debian12 as final

WORKDIR /app

# Copy the built Go exec
COPY --from=builder /build/main ./main

# Run the app
CMD [ "/app/main" ]


