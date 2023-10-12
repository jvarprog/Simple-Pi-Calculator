package main

import (
	"fmt"
	"math"
	"testing"
)

func TestIntMax(t *testing.T) {
	i := math.MaxInt32
	top := float32(calcRatio(i))
	bottom := float32(i)
	result := 4 * top / bottom
	fmt.Printf("IntMax result: %v\n", result)
}

func TestInt8(t *testing.T) {
	i := math.MaxInt8
	top := float32(calcRatio(i))
	bottom := float32(i)
	result := 4 * top / bottom
	fmt.Printf("Int8 result: %v\n", result)
}
