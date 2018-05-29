package controller

import "fmt"

type controller struct{}

func New() *controller {
	return &controller{}
}

func (c *controller) Run() {
	fmt.Println("controller is running")
}
