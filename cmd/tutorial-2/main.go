package main

import "fmt"

func main() {
	var intNum int = 3210
	fmt.Println(intNum)

	var floatNum float64 = 123456.9
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 1
	//var result float32 = floatNum32 + intNum32  will throw error
	var result float32 = floatNum32 + float32(intNum32) //will work becuz type casting is done.
	fmt.Println(result)

	//integer divide by integer is rounded down , if you want remainder use %

	var myString string = "hello world"
	fmt.Println(myString)

	//len function to calculate length of string but, it returns bytes

	var myBoolean bool = false
	fmt.Println(myBoolean)

	var intNumb int //here go sets the default value of the int to 0
	fmt.Println((intNumb))

	//default values for empty string is ''
	//default values for boolean is false

	var myVar = "this_is_string"
	fmt.Println(myVar) //here the data type is infered to be string

	myVar2 := "this_is_also_string" //we can also drop the "var" keyword
	fmt.Println(myVar2)

	var1, var2 := 1, 2 //we can initialise multiple variables at once
	fmt.Println(var1)
	fmt.Println(var2)
	fmt.Println(var1, var2)

	//but adding a type is good practice and makes code readable

	const myConst string = "Const value"
	fmt.Println(myConst)
	// const is same as var, but we can't change its value once its created
	//value of constants need to be declared
}
