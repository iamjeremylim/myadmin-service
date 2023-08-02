# Build stage
FROM golang:1.18-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o services ./cmd/services/main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz

# Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/services .
COPY --from=builder /app/migrate ./migrate
COPY --from=builder /app/seed ./seed
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./migration

# Install postgresql-client
RUN apk add --no-cache postgresql-client

EXPOSE 8080
CMD [ "/app/services" ]
ENTRYPOINT [ "/app/start.sh" ]
