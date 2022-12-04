MYSQL_HOSTNAME := 127.0.0.1
MYSQL_PORT := 3306
MYSQL_USERNAME := root
MYSQL_PASSWORD :=
MYSQL_DATABASE := kurester

server:
	go run main.go

deps:
	go mod download

docs:
	swag init -g httpserver/httpserver.go

migrateup:
	migrate -path pkg/database/migration -database "mysql://${MYSQL_USERNAME}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOSTNAME}:${MYSQL_PORT})/${MYSQL_DATABASE}?parseTime=true" -verbose up

migratedown:
	migrate -path pkg/database/migration -database "mysql://${MYSQL_USERNAME}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOSTNAME}:${MYSQL_PORT})/${MYSQL_DATABASE}?parseTime=true" -verbose down

.PHONY: server deps docs migrateup migratedown
