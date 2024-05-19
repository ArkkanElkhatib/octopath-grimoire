run: build
	@echo 'Running API server...'
	@./bin/api

build:
	@echo 'Building binaries...'
	@go build -o ./bin/api /cmd/api

clean:
	@echo 'Removing binaries...'
	@rm ./bin/api
