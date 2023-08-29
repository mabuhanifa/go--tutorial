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

	var name, country string
	var age int
	var cgpa float32

	name = "Shourov"
	country = "Bangladesh"
	age = 27
	cgpa = 2.54
	fmt.Println("His name is", name, "from", country, "age", age, "with CGPA", cgpa)
}
