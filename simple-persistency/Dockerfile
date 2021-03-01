# https://dev.to/plutov/docker-and-go-modules-3kkn
FROM golang:1.15.6 as builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o bookmark_service
EXPOSE 8080

ENTRYPOINT ["/app/bookmark_service"]
