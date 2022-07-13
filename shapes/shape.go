package main

import (
	"fmt"
	"math"
	"reflect"
)

type shape interface {
	getArea() float64
}
type square struct {
	sideLength float64
}

type triangle struct {
	heightLength float64
	baseLength   float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.baseLength * t.heightLength
}

func (s square) getArea() float64 {
	return math.Pow(s.sideLength, 2)
}

func printArea(sh shape) {
	fmt.Printf("area of %s is : %f \n", reflect.TypeOf(sh).Name(), sh.getArea())
}
