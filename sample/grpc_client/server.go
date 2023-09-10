package main

import (
	"context"
	"fmt"
	"github.com/shshimamo/knowledge/examples/grpc_client/model"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"os"

	pb_example "github.com/shshimamo/knowledge/protobufs/example/gen/pb_go"

	_ "google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	client pb_example.HelloServiceClient
	appEnv model.AppEnv
)

func main() {
	log.Println("main start.")
	appEnv = model.NewAppEnv()
	log.Println("appEnv: " + string(appEnv))
	conn, err := setGRPCClient()
	defer func() { _ = conn.Close() }()
	if err != nil {
		log.Fatal("Connection failed. error: " + err.Error())
		return
	}

	port := getPort()
	log.Println("port: " + port)
	h := setupHandler()
	log.Println("Listen And Serve...")
	log.Fatal(http.ListenAndServe(":"+port, h))
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8083"
	} else {
		return port
	}
}

func setupHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok"))
	})
	mux.HandleFunc("/grpc", sendGRPCRequest)

	return mux
}

func setGRPCClient() (*grpc.ClientConn, error) {
	log.Println("Start setGRPCClient")
	var address string
	if appEnv == model.Production {
		address = "sample-grpc-server-service:8082"
	} else {
		address = "localhost:8082"
	}
	log.Println("address: " + address)
	conn, err := grpc.Dial(
		address,
		grpc.WithUnaryInterceptor(unaryClientInterceptor),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Println("grpc.Dial err: " + err.Error())
		return nil, err
	}

	client = pb_example.NewHelloServiceClient(conn)
	log.Println("Finish setGRPCClient")
	return conn, nil
}

func sendGRPCRequest(w http.ResponseWriter, r *http.Request) {
	req := &pb_example.Message{
		Message: "hoge",
	}

	ctx := context.Background()
	md := metadata.New(map[string]string{"type": "unary", "from": "client"})
	ctx = metadata.NewOutgoingContext(ctx, md)

	var header, trailer metadata.MD
	res, err := client.GetServerResponse(ctx, req, grpc.Header(&header), grpc.Trailer(&trailer))

	var message string
	if err != nil {
		if stat, ok := status.FromError(err); ok {
			message = fmt.Sprintf("code: %s\nmessage: %s\ndetails: %s\n", stat.Code(), stat.Message(), stat.Details())
		} else {
			message = err.Error()
		}
	} else {
		message = fmt.Sprintf("header: %s\ntrailer: %s\nmessage: %s\n", header, trailer, res.GetMessage())
	}

	_, _ = w.Write([]byte(message))
	w.WriteHeader(http.StatusOK)
}

func unaryClientInterceptor(ctx context.Context, method string, req, res interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	fmt.Println("[pre] unary client interceptor", method, req)
	err := invoker(ctx, method, req, res, cc, opts...)
	fmt.Println("[post]  unary client interceptor", res)
	return err
}
