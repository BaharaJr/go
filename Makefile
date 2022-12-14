dropdb:
	docker exec -it go psql -U postgres -c "DROP DATABASE SIMPLEBANK"

createdb:
	docker exec -it go psql -U postgres -c "CREATE DATABASE SIMPLEBANK"

postgres:
	docker run --name go --network local-docker-network -p 5434:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d --restart unless-stopped postgres:14-alpine

migrateup:
	migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5434/simplebank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5434/simplebank?sslmode=disable" -verbose down

sqlc:
	sqlc generate
	
test:
	go test -v -cover ./...

.PHONY: postgres migrateup migratedown sqlc dropdb test