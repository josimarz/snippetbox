test:
	go test -v ./...

run:
	go run ./cmd/web

build:
	go build -o /tmp/web ./cmd/web/
	cp -r ./tls /tmp/