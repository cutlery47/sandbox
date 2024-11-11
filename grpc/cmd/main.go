package main

import (
	"context"
	"grpc-service/internal"
	pb "grpc-service/protos/grpc-service"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx := context.Background()

	grpcServer := grpc.NewServer()
	userService := internal.NewUserService()
	pb.RegisterUserServiceServer(grpcServer, userService)

	// grpc mux
	rmux := runtime.NewServeMux()
	// http mux
	mux := http.NewServeMux()
	mux.Handle("/", rmux)

	// register grpc gateway
	if err := pb.RegisterUserServiceHandlerServer(ctx, rmux, userService); err != nil {
		log.Fatal("RegisterUserServiceHandlerServer:", err)
	}

	reflection.Register(grpcServer)

	grpcAddr, err := net.ResolveTCPAddr("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("ResolveTCPAddr:", err)
	}

	httpAddr, err := net.ResolveTCPAddr("tcp", "localhost:8081")
	if err != nil {
		log.Fatal("ResolveTCPAddr:", err)
	}

	grpcLisnr, err := net.ListenTCP("tcp", grpcAddr)
	if err != nil {
		log.Fatal("ListenTCP:", err)
	}

	httpLisnr, err := net.ListenTCP("tcp", httpAddr)
	if err != nil {
		log.Fatal("ListenTCP:", err)
	}

	exitChan := make(chan os.Signal, 1)
	signal.Notify(exitChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := grpcServer.Serve(grpcLisnr); err != nil {
			log.Fatal("gRPC-server error: ", err)
		}
	}()

	go func() {
		if err := http.Serve(httpLisnr, mux); err != nil {
			log.Fatal("http-server error:", err)
		}
	}()

	<-exitChan
	log.Println("shutdown")
	grpcServer.GracefulStop()
}
