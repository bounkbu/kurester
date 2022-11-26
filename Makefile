server:
	go run main.go

deps:
	go mod download

docs:
	swag init -g httpserver/httpserver.go

migrateup:
	migrate -path pkg/database/migration -database "root:secret@tcp(mysql://127.0.0.1:3306)/kurester?parseTime=true" -verbose up

migratedown:
	migrate -path pkg/database/migration -database "root:secret@tcp(mysql://127.0.0.1:3306)/kurester?parseTime=true" -verbose down

.PHONY: server deps docs migrateup migratedown
