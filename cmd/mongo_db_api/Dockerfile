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

WORKDIR /app/cmd/mongo_db_api
RUN ls -l /app/cmd/mongo_db_api  # Diagnostic: Lister le contenu avant la construction du binaire
RUN go build -o ./mongo_db_api
RUN ls -l /app/cmd/mongo_db_api  # Diagnostic: Lister le contenu après la construction du binaire

FROM alpine:latest

WORKDIR /root/

# Copier le binaire depuis le stage de build
COPY --from=build /app/cmd/mongo_db_api/mongo_db_api .

# Diagnostic: Liste le contenu du répertoire de travail
RUN ls -l /root

# Ajouter les permissions d'exécution au binaire
RUN chmod +x ./mongo_db_api

RUN ls -l /root

EXPOSE 8081

CMD ["./mongo_db_api"]
