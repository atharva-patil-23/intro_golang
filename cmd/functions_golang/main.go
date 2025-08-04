package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "Hello world" //since parameter type is enforced so you can't put other data type
	printMe(printValue)

	var numerator int = 12
	var denominator int = 4
	var result, remainder, err = intDivision(numerator, denominator)
	// if err != nil {
	// 	fmt.Printf(err.Error())
	// } else if remainder == 0 {
	// 	fmt.Printf("the result of the integer division is %v", result)
	// } else {
	// 	fmt.Printf("The result of the integer division is %v and the remainder of the integer division is %v ", result, remainder)
	// }
	switch {
	case err != nil:
		fmt.Printf(err.Error())
	case remainder == 0:
		fmt.Printf("the result of the integer division is %v", result)
	default:
		fmt.Printf("The result of the integer division is %v and the remainder of the integer division is %v ", result, remainder)
	}
}

func printMe(printValue string) { //the opening braces should be here only otherweise go throws error
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) { // the int delclared here suggest a int value is gonna be returned

	var err error
	if denominator == 0 {
		err = errors.New("cannot divide By zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
