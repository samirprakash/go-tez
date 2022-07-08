BINARY_NAME=tezApp

build:
	@go mod vendor
	@go mod tidy
	@echo "Building Tez..."
	@go build -o tmp/${BINARY_NAME} .
	@echo "Tez built!"

run: build
	@echo "Starting Tez..."
	@./tmp/${BINARY_NAME} &
	@echo "Tez started!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm tmp/${BINARY_NAME}
	@echo "Cleaned!"

test:
	@echo "Testing..."
	@go test ./...
	@echo "Done!"

start: run

stop:
	@echo "Stopping Tez..."
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "Stopped Tez!"

restart: stop start