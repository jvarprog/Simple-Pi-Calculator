package main

import (
	"math"
	"testing"
	"fmt"
)

func TestIntMax(t *testing.T) {
	i := math.MaxInt32
	top := float32(calcRatio(i))
	bottom := float32(i)
	fmt.Println(4 * top/bottom)
  }
