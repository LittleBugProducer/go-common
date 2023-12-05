package main

import "context"

func main() {
	VegeOrderStrategy := &VegeOrderStrategy{
		City:    "Beijing",
		OrderID: "123456",
	}
	handler := &OrderContext{
		OrderStrategy: VegeOrderStrategy,
	}
	handler.Handle(context.Background())
}
