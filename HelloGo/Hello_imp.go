package main

import (
	"context"
)

// HelloImp servant implementation
type HelloImp struct {
}

// Init servant init
func (imp *HelloImp) Init() error {
	//initialize servant here:
	//...
	return nil
}

// Destroy servant destroy
func (imp *HelloImp) Destroy() {
	//destroy servant here:
	//...
}

func (imp *HelloImp) Add(ctx context.Context, a int32, b int32, c *int32) (int32, error) {
	//Doing something in your function
	//...
	return 0, nil
}
func (imp *HelloImp) Sub(ctx context.Context, a int32, b int32, c *int32) (int32, error) {
	//Doing something in your function
	//...
	return 0, nil
}
