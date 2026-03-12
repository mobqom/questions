.PHONY: build swag proto run clean
build:
	go build -o bin/app cmd/app/main.go

clean:
	rm -rf bin/

swag:
	swag init -g cmd/app/main.go

proto:
	protoc --go_out=. --go_opt=module=github.com/mobqom/questions --go-grpc_out=. --go-grpc_opt=module=github.com/mobqom/questions proto/v1/question.proto proto/v1/option.proto

run: build
	./bin/app
