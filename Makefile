postgres:
		docker run --name postgres12 -p 5432:5432  -e POSTGRES_USER=root  -e POSTGRES_PASSWORD=superM@n04 -d postgres:12-alpine

createdb:
		docker exec -it postgres12 createdb --username=root --owner=root ecommerce-2021

dropdb:
		docker exec -it postgres12 dropdb  ecommerce-2021

migrateup:
		migrate -path db/migration -database "postgresql://root:superM@n04@localhost:5432/ecommerce-2021?sslmode=disable" -verbose up

migratedown:
		migrate -path db/migration -database "postgresql://root:superM@n04@localhost:5432/ecommerce-2021?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown
