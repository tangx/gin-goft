
up:
	cd cmd/demo/ && go run .
tidy:
	go mod tidy

compose.up:
	docker-compose up -d
compose.down:
	docker-compose down
