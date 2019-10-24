package main

import (
	"fmt"
)

type ServiceSetter interface {
	SetService(Service)
}

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

// SetService implements Service interface and do the stuff of interface injection
func (c *Client) SetService(service Service) {
	c.service = service
}

func main() {
	
	client := Client{}
	client.SetService(&FoodService{})
	client.CallService()
}