.PHONY: proto
proto:
	protoc --go_out=plugins=grpc:internal/common/genproto/books -I api/protobuf api/protobuf/books.proto
