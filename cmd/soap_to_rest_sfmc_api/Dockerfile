# syntax=docker/dockerfile:1

FROM golang:1.18-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

WORKDIR /app/cmd/soap_to_rest_sfmc_api
RUN go build -o /sfmc_api

# Use a minimal image for the final container
FROM alpine:latest

WORKDIR /root/

COPY --from=build /sfmc_api .

EXPOSE 8080

CMD ["./sfmc_api"]
