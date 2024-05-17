# syntax=docker/dockerfile:1

FROM golang:1.22.2-alpine3.19 AS build

WORKDIR /app

# Copier go.mod et go.sum depuis la racine du projet
COPY ../../go.mod ../../go.sum ./
RUN echo "Après la copie de go.mod et go.sum:" && ls -l /app

RUN ls -l /app   # Diagnostic: Lister le contenu après la copie de go.mod et go.sum
RUN go mod download

# Copier le reste du code
COPY ../../ ./
RUN ls -l /app   # Diagnostic: Lister le contenu après la copie du reste du code


WORKDIR /app/cmd/soap_to_rest_sfmc_api
RUN ls -l /app/cmd/soap_to_rest_sfmc_api  # Diagnostic: Lister le contenu avant la construction du binaire

RUN go build -o /soap_to_rest_sfmc_api
RUN ls -l /app/cmd/soap_to_rest_sfmc_api  # Diagnostic: Lister le contenu après la construction du binaire


FROM alpine:latest

WORKDIR /root/

# Copier le binaire construit dans l'image finale
COPY --from=build . .

# Diagnostic: Liste le contenu du répertoire
RUN ls -l  ./

# Ajoute les permissions d'exécution au binaire
RUN chmod +x ./soap_to_rest_sfmc_api

EXPOSE 8080

CMD ["./soap_to_rest_sfmc_api"]
