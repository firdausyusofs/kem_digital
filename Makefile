# ------------------------------------------------------------------------------
#  Go
#
run:
	go run ./cmd/api/main.go

build:
	go build -o ./bin/api ./cmd/api/main.go

test:
	go test -v -cover ./...

seed:
	go run ./cmd/seed/main.go
