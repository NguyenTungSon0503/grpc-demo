.PHONY: proto

proto-js:
    # https://www.npmjs.com/package/grpc_tools_node_protoc_ts
	@grpc_tools_node_protoc \
    --js_out=import_style=commonjs,binary:./dist \
    --grpc_out=grpc_js:./dist \
    --plugin=protoc-gen-grpc=`which grpc_tools_node_protoc_plugin` \
    -I ./proto \
    ./proto/*.proto

proto-ts:
    # https://www.npmjs.com/package/protoc-gen-ts
	@protoc -I=./proto --ts_out=dist helloworld.proto

update-proto:
	@cat ../go-server/proto/helloworld.proto > ./proto/helloworld.proto