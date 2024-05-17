# syntax=docker/dockerfile:1

FROM golang:1.18-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

WORKDIR /app/cmd/mongo_db_api
RUN go build -o /mongo_db_api

# Use a minimal image for the final container
FROM alpine:latest

WORKDIR /root/

COPY --from=build /mongo_db_api .

EXPOSE 8081

CMD ["./mongo_db_api"]
