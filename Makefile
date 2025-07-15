APP_NAME=accounting

run:
	go run ./cmd/$(APP_NAME)

build:
	go build -o ./bin/$(APP_NAME) ./cmd/$(APP_NAME)