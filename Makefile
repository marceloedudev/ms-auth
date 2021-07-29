include env_make

migrate_up:
	migrate -database $(POSTGRES_URI) -path migrations up

migrate_down:
	migrate -database $(POSTGRES_URI) -path migrations down

mocks:
	mockgen -source=internal/user/usecase/interface.go -destination=internal/user/usecase/mock/user.go -package=mock
	mockgen -source=internal/client/usecase/interface.go -destination=internal/client/usecase/mock/client.go -package=mock

test:
	go test -v -coverprofile cover.out -tags testing ./... -failfast
	go tool cover -html=cover.out -o coverage.html

bench:
	go test -v -bench=. -benchmem ./...

test_race:
	 go test -race ./...

start:
	go run cmd/main.go

proto:
	protoc --go_out=internal/user/infra/grpc_user/pb --go_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=internal/user/infra/grpc_user/pb --go-grpc_opt=paths=source_relative --proto_path=internal/user/infra/grpc_user/protofiles internal/user/infra/grpc_user/protofiles/*.proto

evans:
	evans -r repl

docs_generate:
	swag init -dir ./cmd --parseDependency --parseVendor
