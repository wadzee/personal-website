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

# Copy Go executable
COPY --from=builder /build/main ./main

# Copy static file
COPY --from=builder /build/static ./static

ENV PORT 3000
EXPOSE $PORT

# Run the app
CMD [ "/app/main" ]


