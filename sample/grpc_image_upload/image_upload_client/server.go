package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"os"

	pb "github.com/shshimamo/knowledge/protobufs/example/image_uploader/gen/pb_go"
	"google.golang.org/grpc"
)

var (
	client pb.ImageUploadServiceClient
)

const (
	address  = "localhost:50051"
	fileName = "neko.jpg"
)

func main() {
	conn, err := setGRPCClient()
	defer func() { _ = conn.Close() }()
	if err != nil {
		log.Fatal("Connection failed. error: " + err.Error())
		return
	}

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Failed to open file: %v", err)
		return
	}
	defer file.Close()

	stream, err := client.Upload(context.Background())
	if err != nil {
		fmt.Printf("Failed to start stream: %v", err)
		return
	}

	// メタデータを最初に送信
	err = stream.Send(&pb.ImageUploadRequest{
		File: &pb.ImageUploadRequest_FileMeta_{
			FileMeta: &pb.ImageUploadRequest_FileMeta{
				Filename: fileName,
			},
		},
	})
	if err != nil {
		fmt.Printf("Failed to send metadata: %v", err)
		return
	}

	// ファイルの内容をチャンクとして送信
	buf := make([]byte, 1024*100)
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Failed to read file: %v", err)
			return
		}

		err = stream.Send(&pb.ImageUploadRequest{
			File: &pb.ImageUploadRequest_Data{
				Data: buf[:n],
			},
		})
		if err != nil {
			fmt.Printf("Failed to send chunk: %v", err)
			return
		}
		fmt.Printf("send chunk: %v\n", n)
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Printf("Failed to receive response: %v", err)
		return
	}

	fmt.Printf(
		"Uploaded successfully. UUID: %s, size: %s, content_type: %s, filename: %s",
		resp.GetUuid(),
		resp.GetSize(),
		resp.GetContentType(),
		resp.GetFilename(),
	)
}

func setGRPCClient() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, err
	}

	client = pb.NewImageUploadServiceClient(conn)
	return conn, nil
}
