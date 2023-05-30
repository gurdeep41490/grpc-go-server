package main

import (
	"log"

	mgrpc "github.com/gurdeep41490/grpc-go-server/internal/adapter/grpc"
	app "github.com/gurdeep41490/grpc-go-server/internal/application"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(&logWriter{})

	hs := &app.HelloService{}

	grpcAdapter := mgrpc.NewGrpcAdapter(hs, 9090)

	grpcAdapter.Run()
}
