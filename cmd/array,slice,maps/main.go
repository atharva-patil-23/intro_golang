package main

import "fmt"

func main() {
	var intArr [3]int32 //the array holds exatcly 3 elements and cannot change after initialised
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

}
