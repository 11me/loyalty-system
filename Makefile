migrateup:
	migrate -path db/migrations -verbose -database 'postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable' up

migratedown:
	migrate -path db/migrations -verbose -database 'postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable' down

.phony: migrateup migratedown
