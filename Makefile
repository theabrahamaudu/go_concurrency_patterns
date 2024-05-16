build:
	@go build -o bin/conc_patterns cmd/main/main.go

run: build
	./bin/conc_patterns

test:
	@go test -v ./...

clean:
	@go clean
	rm ./bin/conc_patterns