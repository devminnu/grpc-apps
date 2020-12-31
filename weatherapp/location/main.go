package main

import (
	"fmt"
	"net"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/devminnu/weatherapp/location/config"
	"github.com/devminnu/weatherapp/location/db"
	"github.com/devminnu/weatherapp/location/logger"
	"github.com/devminnu/weatherapp/location/service"

	// @grpc - Enable
	pb "github.com/devminnu/weatherapp/pb/location"
)

func main() {
	config.Load()

	logger.Init()
	db.Init()

	// deps := service.Dependencies{
	// 	DB: db.GetStorer(db.Get()),
	// }



	// @grpc
	go GRPCServe()
	server.Run(addr)
}

func GRPCServe() {
	host := config.ReadEnvString("GRPC_HOST")
	port := config.ReadEnvInt("GRPC_PORT")
	// tls := config.ReadEnvBool("TLS")
	// certFile := config.ReadEnvString("CERT_FILE")
	// keyFile := config.ReadEnvString("KEY_FILE")

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		logger.Get().Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	if tls {
		if certFile == "" {
			logger.Get().Fatalf("No certificate file specified")
		}
		if keyFile == "" {
			logger.Get().Fatalf("No key file specified")
		}
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			logger.Get().Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer()

	s := service.GrpcServer{
		DB: db.GetStorer(db.Get()),
	}
	logger.Get().Infof("Grpc config: %v", s)

	// @grpc - Enable after proto file is in place
	pb.RegisterLocationMgrServer(grpcServer, &s)

	logger.Get().Infof("Listening for gRPC on %s:%d", host, port)

	grpcServer.Serve(lis)
}
