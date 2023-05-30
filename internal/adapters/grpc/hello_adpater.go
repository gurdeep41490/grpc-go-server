package grpc

import (
	"context"

	"github.com/gurdeep41490/my-grpc-protogen/protogen/go/hello"
)

func (a *GrpcAdapter) SayHello(ctx context.Context, request *HelloRequest) (*HelloResponse, error) {
	greet := a.helloService.GenerateHello(request.name)

	return &hello.HelloResponse{
		Greet: greet,
	}, nil
}
