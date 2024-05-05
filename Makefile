createdb:
	docker exec -it nps-postgres createdb --username=root --owner=root nps-postgres

dropdb:
		docker exec -it nps-postgres dropdb nps-postgres

migrateUp:
		migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/nps-postgres?sslmode=disable" -verbose up

migrateDown:
		migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/nps-postgres?sslmode=disable" -verbose down

server:
	go run main.go

test:
	go test -v -cover .   