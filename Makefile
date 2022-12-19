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

.PHONY: server deps docs migrateup migratedown chart
