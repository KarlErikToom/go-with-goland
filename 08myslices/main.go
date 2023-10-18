package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices")

	var fruitList = []string{"Apple", "Peach", "Tomato"}

	fmt.Printf("type of fruitlist is %T \n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")

	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])

	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 345
	highScores[2] = 210
	highScores[3] = 450
	//highScores[4] = 777

	highScores = append(highScores, 555, 666, 635)

	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)

	fmt.Println(highScores)
}
