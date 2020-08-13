package main

import (
	"fmt"
	engineGrpc "matching-engine/engineGrpc"
	"matching-engine/server"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":9096"
)

func main() {
	gs := grpc.NewServer()
	cs := server.NewEngine()
	engineGrpc.RegisterEngineServer(gs, cs)

	reflection.Register(gs)

	l, err := net.Listen("tcp", port)
	if err != nil {
		e := fmt.Errorf("Unable to listen server, err: %v", err)
		fmt.Println(e)
		os.Exit(1)
	}
	fmt.Printf("grpc server listening to %s\n", port)
	gs.Serve(l)
}
