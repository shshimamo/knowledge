package main

import (
	"context"
	"log"
	"net"
	"os"

	"google.golang.org/grpc/codes"

	pb_example "github.com/shshimamo/knowledge/protobufs/example/hello/gen/pb_go"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

func main() {
	// gRPC Server
	s := grpc.NewServer(
		grpc.UnaryInterceptor(myUnaryServerInterceptor),
	)

	// Register Helloworld Server to gRPC Server
	pb_example.RegisterHelloServiceServer(s, newMyServer())

	// Register HealthCheck Server to gRPC Server
	healthSrv := health.NewServer()
	healthSrv.SetServingStatus(pb_example.HelloService_ServiceDesc.ServiceName, healthpb.HealthCheckResponse_SERVING)
	healthpb.RegisterHealthServer(s, healthSrv)

	// Register reflection for gRPCurl
	reflection.Register(s)

	// Create Lisnter
	port := getPort()
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}

	// Start gRPC Server
	log.Printf("start gRPC server port: %v", port)
	s.Serve(listener)
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8082"
	} else {
		return port
	}
}

func myUnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("[pre] my unary server interceptor 1: ", info.FullMethod)
	res, err := handler(ctx, req)
	log.Println("[post] my unary server interceptor 1: ", res)
	return res, err
}

func newMyServer() *myServer {
	return &myServer{}
}

type myServer struct {
	pb_example.UnimplementedHelloServiceServer
}

func (s *myServer) GetServerResponse(ctx context.Context, message *pb_example.Message) (*pb_example.MessageResponse, error) {
	if message.Message == "error" {
		return nil, status.Error(
			codes.InvalidArgument,
			"Invalid message",
		)
	}

	result := "Thanks for talking to gRPC server!!! Welcome to hello world. Received message is: " + message.Message
	return &pb_example.MessageResponse{Message: result, Received: true}, nil
}

func (s *myServer) Optional(ctx context.Context, message *pb_example.Message) (*pb_example.OptionalResponse, error) {
	//m := &pb_example.Message{Message: "hoge"}
	return &pb_example.OptionalResponse{MessageStr: "hoge"}, nil
}
