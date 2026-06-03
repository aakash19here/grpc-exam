proto:
	protoc \
		--proto_path=proto \
		--go_out=proto \
		--go-grpc_out=proto \
		proto/*.proto
	@echo "Proto files generated in the 'proto' directory."


.PHONY: proto server client_unary client_server client_client client_bidi