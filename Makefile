genobu:
	@go build -o bin/obu ./obu
	@./bin/obu

start-receiver:
	@go build -o bin/receiver ./receiver
	@./bin/receiver