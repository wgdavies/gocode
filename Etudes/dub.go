package main

import (
    "fmt"
    "runtime"
)

func main() {
	var (
		input float64=0
		output float64=0
	)

	fmt.Print("Enter a number: ")
//	var input float64
	fmt.Scanf("%f", &input)

//	output := input * 2
	output = input * 2

	fmt.Println(output)
//	fmt.Println(input * 2)
}

