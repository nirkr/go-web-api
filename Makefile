postgres:
	docker run --name myPostgres -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:12-alpine
createdb:
	docker exec -it myPostgres createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it myPostgres dropdb simple_bank
migrateup:
	migrate -path db/migration -database "posgresql://root:secret@loaclhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "posgresql://root:secret@loaclhost:5432/simple_bank?sslmode=disable" -verbose down
run:
	go run cmd/main.go

.PHONY: createdb dropdb postgres migrateup migratedown
