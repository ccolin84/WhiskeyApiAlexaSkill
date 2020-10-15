build: main.zip
	@echo "Done with build!"

test:
	go test -v

main.zip: main
	zip main.zip main

main:
	go build -v -o main main.go alexa.go whiskey_api.go
