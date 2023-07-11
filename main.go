package main

import (
	"context"
	"log"
	"net"

	"github.com/MrDavudov/demo-grpc/adder"
	"google.golang.org/grpc"
)

type myAdderServer struct {
	adder.UnimplementedAdderServer
}

func (s myAdderServer) Add(ctx context.Context, in *adder.AddRequest) (*adder.AddResponce, error) {
	return &adder.AddResponce{
		Result: int32(in.X) + int32(in.Y),
	}, nil
}

func main() {
	l, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatal("cannot create listener: %s", err)
	}

	srv := grpc.NewServer()
	s := &myAdderServer{}

	adder.RegisterAdderServer(srv, s)

	err = srv.Serve(l)
	if err != nil {
		log.Fatalf("imposible to serve: %s", err)
	}
	log.Print("Start server")
}