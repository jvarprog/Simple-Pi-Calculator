package main

import (
	"math/rand"
  "fmt"
)

func calcRatio(sampleSize int) int {
  returnValue := 0

  for i := 0; i <= sampleSize; i++ {
    x := rand.Float32()
    y := rand.Float32()

    if (x*x + y*y <= 1) {
      returnValue++
    }
  } 
  return returnValue
}

func main() { 
  fmt.Println("Sample Size: ")
}
