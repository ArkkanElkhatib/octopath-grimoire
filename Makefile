run-web: build-api
	@echo 'Running Web server...'
	@./bin/web

build-web:
	@echo 'Building Web binaries...'
	@go build -o ./bin/web /cmd/web

run-api: build-api
	@echo 'Running API server...'
	@./bin/api

build-api:
	@echo 'Building API binaries...'
	@go build -o ./bin/api /cmd/api

clean:
	@echo 'Removing binaries...'
	@rm ./bin/api
	@rm ./bin/web
