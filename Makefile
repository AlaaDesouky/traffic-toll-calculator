genobu:
	@go build -o bin/obu ./obu
	@./bin/obu

receiver: kafka
	@go build -o bin/receiver ./receiver
	@./bin/receiver

calculator: kafka
	@go build -o bin/calculator ./calculator
	@./bin/calculator

kafka:
	docker compose up -d

.PHONY: calculator