package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
)

func calcRatio(sampleSize int) int {
	returnValue := 0

	for i := 0; i <= sampleSize; i++ {
		x := rand.Float32()
		y := rand.Float32()

		if x*x+y*y <= 1 {
			returnValue++
		}
	}
	return returnValue
}

func determineSampleSize(s string) int {
	switch s {
	case "IntMax":
		return math.MaxInt32
	case "Int32":
		return math.MaxInt32
	case "Int16":
		return math.MaxInt16
	case "Int8":
		return math.MaxUint8
	}
	returnValue, err := strconv.Atoi(s)
	if err == nil {
		return returnValue
	}
	return 0
}

func monteCarloSim() {
	fmt.Print("Sample Size?: ")
	var i string
	fmt.Scanln(&i)
	sampleSize := determineSampleSize(i)
	sampleTop := calcRatio(sampleSize)
	piAprox := 4.0 * float32(sampleTop) / float32(sampleSize)
	fmt.Printf("Pi Aproximated with %v samples: %v\n", sampleSize, piAprox)
}

func leibnizSim() {
  fmt.Print("Sample Size?: ")
  var input string
  fmt.Scanln(&input)
  sampleSize := determineSampleSize(input)

  sum, term := 0.0, 0.0
  for i := 0; i <= sampleSize; i++ {
    term = math.Pow(-1.0, float64(i)) / (2.0 * float64(i) + 1.0)
    sum += term
  } 
  pi := sum * 4.0

  fmt.Printf("Pi aproximated with %v samples: %v\n", sampleSize, pi)
}

func optionsMenu() int {
	fmt.Print("Options: \nMonte Carlo = 1\nLeibniz Sum = 2\nExit = exit\n")
	fmt.Print("Option? : ")
	var input string
	fmt.Scanln(&input)

	switch input {
	case "1":
		return 1
	case "2":
		return 2
	case "exit":
		fmt.Println("Exiting...")
		os.Exit(1)
	}
	return -1
}

func runOption(input int) {
	switch input {
	case 1:
		monteCarloSim()
	case 2:
		leibnizSim()
	case -1:
		fmt.Println("Input error")
	}

}

func main() {
	optionValue := optionsMenu()
	runOption(optionValue)
}
