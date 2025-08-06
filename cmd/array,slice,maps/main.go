package main

import "fmt"

func main() {
	var intArr [3]int32 //the array holds exatcly 3 elements and cannot change after initialised

	var Array2 [3]int32 = [3]int32{1, 2, 3} //imidiatly initialise the array

	//or

	Array3 := [3]int32{1, 2, 3}

	//or

	Array4 := [...]int32{1, 2, 3, 4}

	fmt.Println(Array3)
	fmt.Println(Array2)
	fmt.Println(Array4)
	intArr[1] = 123
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

}
