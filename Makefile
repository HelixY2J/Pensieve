run: build
	@./bin/pensieve

build:
	@go build -o bin/pensieve

test:
	@go test -v ./...