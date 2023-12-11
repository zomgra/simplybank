postgres:
	docker run --name postgres12 -d -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret postgres:12-alpine
createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simply_bank

dropdb:
	docker exec -it postgres12 dropdb simply_bank

migrateup:
	migrate -path internal/db/migrations -database "postgresql://root:secret@localhost:5432/simply_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path internal/db/migrations -database "postgresql://root:secret@localhost:5432/simply_bank?sslmode=disable" -verbose down

test:
	go test -v -cover ./... 

sqlc_generate:
	docker run --rm -v $(CURDIR):/src -w /src sqlc/sqlc generate
.PHONY: sqlc_generate sqlc migrateup migratedown createdb postgres dropdb test