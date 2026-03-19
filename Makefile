.PHONY: build swag run clean
build:
	go build -o bin/app cmd/app/main.go

clean:
	rm -rf bin/

swag:
	swag init -g cmd/app/main.go

run: build
	./bin/app
