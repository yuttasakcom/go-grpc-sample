package main

import (
	"fmt"
	"net"

	"github.com/yuttasakcom/go-grpc-sample/services"
	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	services.RegisterCalulatorServer(s, services.NewCalculatorServer())

	fmt.Println("Server is running on port :8080")
	s.Serve(listener)
}
