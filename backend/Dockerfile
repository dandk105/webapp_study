FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o myapp

FROM alpine:3.18

# Set the POSTGRES_USER, POSTGRES_PASSWORD, and POSTGRES_DB environment variables
ENV DB_USER=myuser
ENV DB_PASSWORD=mypassword
ENV DB_NAME=mydatabase


COPY --from=builder /app/myapp /usr/local/bin/myapp

ENTRYPOINT ["/usr/local/bin/myapp"]
EXPOSE 5000