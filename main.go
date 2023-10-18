package main

import (
	"fmt"
	"os"

	PiAlgo "github.com/jvarprog/Pi-Algo"
)

func optionsMenu() int {
	fmt.Print("\nOptions: \nMonte Carlo = 1\nLeibniz Sum = 2\nGauss-Legendre Method = 3\nExit = exit\n")
	fmt.Print("Option? : ")
	var input string
	fmt.Scanln(&input)

	switch input {
	case "1":
		return 1
	case "2":
		return 2
	case "3":
		return 3
	case "exit":
		fmt.Println("Exiting...")
		os.Exit(1)
	}
	return -1
}

func runOption(input int) {
	switch input {
	case 1:
		PiAlgo.MonteCarloSim()
	case 2:
		PiAlgo.LeibnizSim()
	case 3:
		PiAlgo.GaussLegendreMethod()
	case -1:
		fmt.Println("Input error")
	}

}

func main() {
	optionValue := optionsMenu()
	for optionValue != -1 {
    runOption(optionValue)
    optionValue = optionsMenu()
  }
}
