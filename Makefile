run-server:
	@go run cmd/api-server/main.go

run-nats:
	@nats-server -js

run-redis:
	@redis-server
