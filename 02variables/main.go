package main

import "fmt"

const LoginToken string = "aaabbbbccc" //Public

func main() {
	var username string = "hitesh"
	fmt.Println(username)
	fmt.Printf("Variable is of type:  %T\n", username)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type:  %T\n", smallVal)

	var smallFloat float32 = 255.45541259
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type:  %T\n", smallFloat)

	//Deafault values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type:  %T\n", anotherVariable)

	//implicit type

	var website = "leanrcodeonline.in"
	fmt.Println(website)

	//no var style

	numberOfUsers := 3000
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type:  %T\n", LoginToken)

}
