FROM golang:1.22.1 as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o /server ./cmd

FROM gcr.io/distroless/base-debian11 as final

COPY --from=builder /server /server

ENV PORT 3000
EXPOSE $PORT

ENTRYPOINT ["/server"]