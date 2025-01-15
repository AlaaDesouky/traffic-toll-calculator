genobu:
	@go build -o bin/obu ./obu
	@./bin/obu

receiver: kafka
	@go build -o bin/receiver ./receiver
	@./bin/receiver

calculator: kafka
	@go build -o bin/calculator ./calculator
	@./bin/calculator

aggregator:
	@go build -o bin/aggregator ./aggregator
	@./bin/aggregator

kafka:
	docker compose up -d

.PHONY: calculator aggregator