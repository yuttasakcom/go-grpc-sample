package services

import (
	context "context"
	"fmt"
)

type calculatorServer struct {
}

func NewCalculatorServer() CalulatorServer {
	return calculatorServer{}
}

func (c calculatorServer) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	result := fmt.Sprintf("Hello, %s!", req.Name)
	res := &HelloResponse{
		Message: result,
	}
	return res, nil
}

func (c calculatorServer) mustEmbedUnimplementedCalulatorServer() {}
