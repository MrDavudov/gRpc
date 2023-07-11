generate_grpc_code:
	protoc \
	--go_out=adder \
	--go_opt=paths=source_relative \
	--go-grpc_out=adder \
	--go-grpc_opt=paths=source_relative \
	adder.proto