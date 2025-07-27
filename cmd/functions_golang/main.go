package main

import "fmt"

func main() {
	var printValue string = "Hello world"
	printMe(printValue)
}

func printMe(printValue string) { //the opening braces should be here only otherweise go throws error
	fmt.Println(printValue)
}
