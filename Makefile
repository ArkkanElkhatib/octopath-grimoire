run: build
	@echo 'Running server...'
	@./bin/server

build:
	@echo 'Building binaries...'
	@go build -o ./bin/server ./cmd/server

clean:
	@echo 'Removing binaries...'
	@rm ./bin/server
