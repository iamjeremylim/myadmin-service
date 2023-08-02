postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root my_admin

dropdb:
	docker exec -it postgres12 dropdb my_admin

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/my_admin?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/my_admin?sslmode=disable" -verbose down

server:
	go run ./cmd/services/main.go

.PHONY: postgres createdb dropdb migrateup migratedown server
