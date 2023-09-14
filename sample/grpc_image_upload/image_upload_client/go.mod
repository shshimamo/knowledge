module github.com/shshimamo/knowledge/sample/grpc_image_upload/image_upload_client

go 1.21.1

replace github.com/shshimamo/knowledge/protobufs/example/image_uploader/gen/pb_go => ../../../protobufs/example/image_upload/gen/pb_go/

require (
	github.com/shshimamo/knowledge/protobufs/example/image_uploader/gen/pb_go v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.58.1
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.12.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
	golang.org/x/text v0.11.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230711160842-782d3b101e98 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)
