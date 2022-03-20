package main

import (
	"fmt"
	"log"
)

type floats interface {
	float32 | float64
}
type uints interface {
	uint | uint8 | uint16 | uint32 | uint64
}
type ints interface {
	int | int8 | int16 | int32 | int64
}

type numbers interface {
	ints | uints | floats
}

func adder(a, b int) int {
	return a + b
}

func genericAdder[T numbers](a, b T) T {
	return a + b
}

func mapSum[K comparable, T numbers](m map[K]T) T {
	var sum T
	for _, v := range m {
		sum = sum + v
	}
	return sum
}
func realMain() error {
	stuff := map[string]float32{
		"foo": 5.0,
		"bar": 3.5,
	}
	fmt.Println("generic adder: ", genericAdder(3, 5))
	fmt.Println("adder: ", adder(3, 5))
	fmt.Println("map:", mapSum(stuff))
	return nil
}

func main() {
	err := realMain()
	if err != nil {
		log.Fatal(err)
	}
}
