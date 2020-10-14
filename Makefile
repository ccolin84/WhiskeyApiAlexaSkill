
build: main.zip
	@echo "Done!"

main.zip: main
	zip main.zip main

main:
	go build main.go alexa.go whiskey_api.go
