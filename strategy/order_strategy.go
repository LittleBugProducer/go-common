package main

import "context"

type OrderStrategy interface {
	Handle(ctx context.Context) error
}

type VegeOrderStrategy struct {
	City    string
	OrderID string
}

func (s *VegeOrderStrategy) Handle(ctx context.Context) error {
	// pre handle
	// order allocate
	// order distribute
	// after handle, msg notify
	return nil
}

type SpliceOrderStrategy struct {
	Sum     int
	OrderID string
}

func (s *SpliceOrderStrategy) Handle(ctx context.Context) error {
	// pre handle
	// wait notify
	// distribute
	// after handle, msg notify
	return nil
}

type NormalOrderStrategy struct {
	OrderID string
}

func (s *NormalOrderStrategy) Handle(ctx context.Context) error {
	// pre handle
	// order distribute
	// check confirm
	// after handle, msg notify
	return nil
}

type OrderContext struct {
	OrderStrategy OrderStrategy
}

func (c *OrderContext) Handle(ctx context.Context) error {
	return c.OrderStrategy.Handle(ctx)
}
