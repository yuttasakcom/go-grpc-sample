package services

import (
	"context"
	"fmt"
)

type CalculatorService interface {
	Hello(name string) error
}

type calculatorService struct {
	calulatorClient CalulatorClient
}

func NewCalculatorService(calulatorClient CalulatorClient) CalculatorService {
	return &calculatorService{calulatorClient}
}

func (c *calculatorService) Hello(name string) error {
	req := &HelloRequest{
		Name: name,
	}
	res, err := c.calulatorClient.Hello(context.Background(), req)
	if err != nil {
		return err
	}
	fmt.Println("Message: ", res.Message)
	return nil
}
