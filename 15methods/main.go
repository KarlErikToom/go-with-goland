package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	//no inheritance in golang; no super or parent

	hitesh := User{"Hitesh", "hites@email.com", true, 16}

	fmt.Println(hitesh)
	fmt.Printf("hitesh details are : %+v \n ", hitesh)
	fmt.Printf("Name is %v and email is %v \n ", hitesh.Name, hitesh.Email)
	hitesh.GetStatus()
	hitesh.NewMail()
	fmt.Printf("Name is %v and email is %v \n ", hitesh.Name, hitesh.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)

}

func (u User) NewMail() {

	u.Email = "test@test.com"
	fmt.Println("Email of this user is: ", u.Email)
}
