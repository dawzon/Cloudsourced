package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("+-------------------------+")
	fmt.Println("| Welcome to Cloudsourced |")
	fmt.Println("+-------------------------+")
	fmt.Println()

	fmt.Println("What would you like to do?")
	fmt.Println("1. Worker mode")
	fmt.Println("2. Server mode")
	fmt.Println("...")
	fmt.Println("0. Exit")

	var choice = -1
	for {
		fmt.Scan(&choice)
		switch choice {
		case 0:
			os.Exit(0)
		case 1: //Worker mode
			fmt.Println("Entering worker mode...")
		case 2: //Server mode
			fmt.Println("Entering server mode...")
		default:
			fmt.Println("Invalid choice")
			//Ask for another number
		}
	}
}
