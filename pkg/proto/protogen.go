package proto

//go:generate protoc --go_out=../pb --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:../pb --go-grpc_opt=paths=source_relative message.proto
//go:generate protoc --go_out=../pb --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:../pb --go-grpc_opt=paths=source_relative product.proto
//go:generate protoc --go_out=../pb --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:../pb --go-grpc_opt=paths=source_relative common.proto
//go:generate protoc --go_out=../pb --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:../pb --go-grpc_opt=paths=source_relative pagination.proto
//go:generate protoc --go_out=../pb --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:../pb --go-grpc_opt=paths=source_relative services.proto
//go:generate protoc --go_out=../pb --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:../pb --go-grpc_opt=paths=source_relative types.proto
//go:generate protoc --go_out=../pb --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:../pb --go-grpc_opt=paths=source_relative user.proto

//go:generate protoc -I . --grpc-gateway_out=../pb --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true services.proto
