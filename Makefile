
build: main.zip
	@echo "Done!"

test:
	go test -v

main.zip: main
	zip main.zip main

main:
	go build main.go alexa.go whiskey_api.go
