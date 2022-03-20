build:
	echo hello

protobuf:
	protoc --go_out=. --go_opt=paths=source_relative pb/service.proto

grpc:
	protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative --go_out=. --go_opt=paths=source_relative pb/service.proto