genobu:
	@go build -o bin/obu ./obu
	@./bin/obu

receiver: kafka
	@go build -o bin/receiver ./receiver
	@./bin/receiver

kafka:
	docker compose up -d