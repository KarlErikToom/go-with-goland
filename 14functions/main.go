package main

import "fmt"

func main() {

	fmt.Println("Welcome to functions in golang")
	greeter()
	result := adder(3, 5)

	fmt.Println("Result is ", result)

	proRes := proAdder(2, 5, 7, 3)
	fmt.Println("Pro results is:", proRes)
}

func adder(num1 int, num2 int) int {
	return num1 + num2
}
func proAdder(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func greeter() {

	fmt.Println("Namastey from golang")
}
