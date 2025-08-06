package main

import "fmt"

func main() {
	//different ways to declare an array
	var intArr [3]int32                     //the array holds exatcly 3 elements and cannot change after initialised
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

	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Println(intSlice)
	fmt.Println("the length is %v with capacity %v", len(intSlice), cap(intSlice))

	intSlice = append(intSlice, 7) //unlike array we can append an element to the slice
	// when we append a new element it checks if there is any space for the new element to be appended to the array , if not a new array is created with the double the original capacity
	// but still we cannot access the new empty indexes
	fmt.Println(intSlice)
	fmt.Println("the length is %v with capacity %v", len(intSlice), cap(intSlice)) //here the cap shows the capacity of the array
	//[4,5,6] => [4,5,6,7,*,*]
	//we can't access the * we get index out of range error

	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...) // we append multiple values using spread operator
	fmt.Println(intSlice)
}
