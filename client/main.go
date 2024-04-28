package main

import (
	"log"

	"github.com/yuttasakcom/go-grpc-sample/client/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	creds := insecure.NewCredentials()
	cc, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	calculatorClient := services.NewCalulatorClient(cc)
	calculatorService := services.NewCalculatorService(calculatorClient)

	err = calculatorService.Hello("Yuttasak")
	if err != nil {
		log.Fatal(err)
	}
}
