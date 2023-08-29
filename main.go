package main

import "fmt"

func main() {
	/*
		fmt.Println("hello world")

		//scape sequence new line
		fmt.Println("1 \n2")

		//scape sequence tab
		fmt.Println("1 \t2")

		//scape sequence \\
		fmt.Println("1 \\ 2")

		//scape sequence " "
		fmt.Println("\"hello world\"")
	*/

	/*
		name := "Shourov"
		var country string = "Bangladesh"
		var age int = 27
		var cgpa float32 = 2.54

		fmt.Println("His name is", name, "from", country, "age", age, "with CGPA", cgpa)

	*/

	name := "Shourov"
	var country string = "Bangladesh"
	var age int = 27
	var cgpa float32 = 2.54

	// printF format
	fmt.Printf("%v is my name \n", name)
	fmt.Printf(". My age is %v \n", age)
	fmt.Printf(". %v is my country \n", country)
	fmt.Printf(". My CGPA is %v \n", cgpa)

}
