package main

import (
	"context"
	"database/sql"
	"log"
	"net"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"github.com/shshimamo/knowledge-auth/handler"
	"github.com/shshimamo/knowledge-auth/model"
	hellopb "github.com/shshimamo/knowledge-auth/pkg/grpc"
	"github.com/shshimamo/knowledge-auth/utils"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

func main() {
	var err error

	appEnv := model.NewAppEnv()

	db, err := utils.SetupDatabase(appEnv)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	go func() {
		port := getPort()
		h := setupHandler(db, appEnv)
		log.Fatal(http.ListenAndServe(":"+port, h))
	}()

	// gRPC Server
	s := grpc.NewServer(
		grpc.UnaryInterceptor(myUnaryServerInterceptor),
	)

	// Register Helloworld Server to gRPC Server
	hellopb.RegisterHelloworldServer(s, newMyServer())

	// Register HealthCheck Server to gRPC Server
	healthSrv := health.NewServer()
	healthSrv.SetServingStatus(hellopb.Helloworld_ServiceDesc.ServiceName, healthpb.HealthCheckResponse_SERVING)
	healthpb.RegisterHealthServer(s, healthSrv)

	// Register reflection for gRPCurl
	reflection.Register(s)

	// Create Lisnter
	gport := gRPCPort()
	listener, err := net.Listen("tcp", ":"+gport)
	if err != nil {
		panic(err)
	}

	// Start gRPC Server
	log.Printf("start gRPC server port: %v", gport)
	s.Serve(listener)
}

func setupHandler(db *sql.DB, appEnv model.AppEnv) http.Handler {
	auth := handler.NewAuthHandler(db, appEnv)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok"))
	})
	mux.HandleFunc("/signup", auth.Signup)
	mux.HandleFunc("/signin", auth.Signin)
	mux.HandleFunc("/signout", auth.Signout)

	var allowOrigins []string
	if appEnv == model.Production {
		allowOrigins = []string{"http://frontend.main.shshimamo.com"}
	} else {
		allowOrigins = []string{"http://localhost:3000"}
	}
	c := cors.New(cors.Options{
		AllowedOrigins:   allowOrigins,
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "HEAD", "POST", "PUT", "OPTIONS"},
		AllowedHeaders:   []string{"*"}, // Allow All HTTP Headers
	})
	h := c.Handler(mux)

	return h
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8081"
	} else {
		return port
	}
}

func gRPCPort() string {
	port := os.Getenv("GRPC_PORT")
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
	hellopb.UnimplementedHelloworldServer
}

func (s *myServer) GetServerResponse(ctx context.Context, message *hellopb.Message) (*hellopb.MessageResponse, error) {
	if message.Message == "error" {
		return nil, status.Error(
			codes.InvalidArgument,
			"Invalid message",
		)
	}

	result := "Thanks for talking to gRPC server!!! Welcome to hello world. Received message is: " + message.Message
	return &hellopb.MessageResponse{Message: result, Received: true}, nil
}
