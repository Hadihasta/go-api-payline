# this folder for simplify runing docker on local on any device 
# run with make postgres make createdb make dropdb
postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb: 
	docker exec -it postgres12 createdb --username=postgres --owner=postgres payline_v1

dropdb: 
	docker exec -it postgres12 dropdb --username=postgres payline_v1

migrateup: 
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5432/payline_v1?sslmode=disable" -verbose up

migratedown: 
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5432/payline_v1?sslmode=disable" -verbose down

.PHONY:postgres createdb dropdb migrateup migratedown