package main

import "fmt"

func main() {
	var printValue string = "Hello world" //since parameter type is enforced so you can't put other data type
	printMe(printValue)

	var numerator int = 12
	var denominator int = 4
	var result int = intDivision(numerator, denominator)
	fmt.Println(result)
}

func printMe(printValue string) { //the opening braces should be here only otherweise go throws error
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) int { // the int delclared here suggest a int value is gonna be returned
	var result int = numerator / denominator
	return result
}
