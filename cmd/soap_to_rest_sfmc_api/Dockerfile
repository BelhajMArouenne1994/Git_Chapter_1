# syntax=docker/dockerfile:1

FROM golang:1.22.2-alpine3.19 AS build

WORKDIR /app

# Copier go.mod et go.sum depuis la racine du projet
COPY ../../go.mod ../../go.sum ./
RUN go mod download

# Copier le reste du code
COPY ../../ ./

WORKDIR /app/cmd/soap_to_rest_sfmc_api
RUN go build -o /sfmc_api

FROM alpine:latest

WORKDIR /root/

# Copier le binaire construit dans l'image finale
COPY --from=build /sfmc_api .

# Diagnostic: Liste le contenu du répertoire
RUN ls -l  ./

# Ajoute les permissions d'exécution au binaire
RUN chmod +x ./sfmc_api

EXPOSE 8080

CMD ["./sfmc_api"]
