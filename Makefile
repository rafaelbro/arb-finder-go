build:
	go build -o dist/bin/ ./...

linux:
	GOOS=linux GOARCH=amd64 go build -o dist/bin-linux-amd64/ ./...

clean:
	rm -rf dist/
