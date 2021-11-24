.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/catalog/healthy.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/catalog/shared.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/catalog/volume.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/catalog/path.proto
