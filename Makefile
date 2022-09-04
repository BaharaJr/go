postgres:
	docker run --name go --network local-docker-network -p 5434:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:14-alpine

migrateup:
	migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5434/simple?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5434/simple?sslmode=disable" -verbose down

.PHONY: postgres migrateup migratedown