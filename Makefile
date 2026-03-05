.PHONY: swag proto

swag:
	swag init -g cmd/app/main.go

proto:
	protoc --go_out=. --go_opt=module=github.com/mobqom/questions --go-grpc_out=. --go-grpc_opt=module=github.com/mobqom/questions proto/v1/question.proto proto/v1/option.proto
