obu:
	@go build -o bin/obu obu/main.go
	@./bin/obu

reciever:
	@go build -o bin/data_reciever data_reciever/main.go
	@./bin/data_reciever

.PHONY: obu