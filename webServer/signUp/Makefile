postgres:
	docker run --name intellias -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it intellias createdb --username=root --owner=root users

dropdb:
	docker exec -it intellias dropdb users

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/users?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/users?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc
