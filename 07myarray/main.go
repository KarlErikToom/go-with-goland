package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Banana"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list is: ", len(fruitList))

	var vegList = [3]string{"Potato", "beans", "mushroom"}

	fmt.Println("Vegetable list is: ", len(vegList))
}
