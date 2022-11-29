// proto generator --1  (source_relative)

protoc --go_out=./proto-gen --go_opt=paths=source_relative \
    --go-grpc_out=./proto-gen --go-grpc_opt=paths=source_relative \
    protos/dice.proto



// proto generator --2

protoc --go_out=./proto-gen \
    --go-grpc_out=./proto-gen \
    protos/dice.proto