package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/gurdeep41490/grpc-go-server/internal/port"
	"github.com/gurdeep41490/my-grpc-protogen/protogen/go/hello"
	"google.golang.org/grpc"
)

type GrpcAdapter struct {
	helloService port.HelloServicePort
	grpcPort     int
	server       *grpc.Server
	hello.HelloServiceServer
}

func NewGrpcAdapter(helloService port.HelloServicePort, grpcPort int) *GrpcAdapter {
	return &GrpcAdapter{
		helloService: helloService,
		grpcPort:     grpcPort,
	}
}

func (a *GrpcAdapter) Run() {
	var err error

	listen, err := net.Listen("tcp", fmt.Sprintf(" :%d", a.grpcPort))

	if err != nil {
		log.Fatalf("Failed to listen to port %d : %v\n", a.grpcPort, err)
	}

	log.Println("Server listening at port ", a.grpcPort)

	grpcServer := grpc.NewServer()

	a.server = grpcServer

	hello.RegisterHelloServiceServer(grpcServer, a)

	if err = grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to listen to port %d : %v\n", a.grpcPort, err)
	}
}

func (a *GrpcAdapter) Stop() {
	a.server.Stop()
}
