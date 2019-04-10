.PHONY: build clean run

build: clean
	@go build -o ./bin/club ./cmd/club

clean:
	@rm -rf ./bin/*

race: clean
	@go build -race -o ./bin/club ./cmd/club

run:
	@go run ./cmd/club