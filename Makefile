postgres:
	sudo docker run --name postgres14 -p 8080:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres

createdb:
	sudo docker exec -it postgres14 createdb --username=root --owner=root simple_bank

dropdb:
	sudo docker exec -it postgres14 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:password@localhost:8080/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:password@localhost:8080/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test: 
	go test -v -cover ./...

server: 
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/Emmrys-Jay/simple-bank/db/sqlc Store

.PHONY: createdb dropdb postgres migratedown migrateup sqlc test server mock