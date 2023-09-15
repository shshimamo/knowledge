module github.com/shshimamo/knowledge/sample/grpc_hello/grpc_server

go 1.21

require (
	github.com/shshimamo/knowledge/protobufs/example/hello/gen/pb_go v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.58.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.15.0 // indirect
	golang.org/x/sys v0.12.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230822172742-b8732ec3820d // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)

replace github.com/shshimamo/knowledge/protobufs/example/hello/gen/pb_go => ../../../protobufs/example/hello/gen/pb_go
