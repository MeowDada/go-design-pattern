package main

import (
	"fmt"
)

type Service interface {
	GetServiceName() string
}

type FoodService struct {}

func (c *FoodService) GetServiceName() string {
	return "Food Service"
}

type Client struct {
	service Service
}

func (c *Client) CallService() {
	fmt.Println(c.service.GetServiceName())
}

func main() {
	
	// Golang does not have constructor, so we use constructor-like factory method to replace it
	client := Client {
		service: &FoodService{},
	}

	client.CallService()
}