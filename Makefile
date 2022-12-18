SHELL:=/bin/bash

proto :
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative src/shared/grpc/**/*.proto
run :
	go run main.go
lint :
	golangci-lint run --print-issued-lines=false --exclude-use-default=false --enable=goimports  --enable=unconvert --enable=unparam --enable=gosec --timeout=2m
test:
	ENV=test go test -race -coverprofile coverage.cov -cover ./... && go tool cover -func coverage.cov
test_coverage:
	ENV=test go test ./... -coverprofile coverage.out && go tool cover -html=coverage.out -o coverage.html
mock:
	rm -r -f src/mocks/
	mockery --dir src --output src/mocks --all --keeptree
