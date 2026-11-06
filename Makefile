.PHONY: test vet build fmt
fmt:
	gofmt -w ./cmd ./internal ./pkg
test:
	go test ./...
vet:
	go vet ./...
build:
	go build -o bin/nmp-snmp ./cmd/nmp-snmp
