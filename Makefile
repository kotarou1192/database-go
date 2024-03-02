build:
	go build -o dist/app
run:
	test -f dist/app && dist/app || go run main.go
test:
	go test -v ./...
