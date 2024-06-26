gate:
	@go build -o bin/gate gateway/main.go
	@./bin/gate

obu:
	@go build -o bin/obu obu/main.go
	@./bin/obu

reciever:
	@go build -o bin/reciever ./data_reciever
	@./bin/reciever


calculator:
	@go build -o bin/calculator ./distance_calculator
	@./bin/calculator

aggregator:
	@go build -o bin/aggregator ./aggregator
	@./bin/aggregator


proto:
	protoc --go_out=. --go_opt=paths=source_relative types/ptypes.proto  --go-grpc_out=. --go-grpc_opt=paths=source_relative types/ptypes.proto


 
.PHONY: obu , aggregator