package entity

import (
	"fmt"
	"runtime"
)

func init() {
	fmt.Println(runtime.NumCPU(), "--------------------------------")
	if 32 != runtime.NumCPU() { // runtime.NumCPU() != 32
		panic("this api is meant to be run on 32 core machines")
	}
}

type Coupon struct {
	ID             string
	Code           string
	Discount       int
	MinBasketValue int
}
