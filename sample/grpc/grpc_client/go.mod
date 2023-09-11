module github.com/shshimamo/knowledge/examples/grpc_client

go 1.21.0

require (
	github.com/shshimamo/knowledge/protobufs/example/gen/pb_go v0.0.0-20230910081736-1eabec03d0bd
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230822172742-b8732ec3820d
	google.golang.org/grpc v1.58.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.15.0 // indirect
	golang.org/x/sys v0.12.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)

replace github.com/shshimamo/knowledge/protobufs/example/gen/pb_go => ../../../protobufs/example/gen/pb_go
