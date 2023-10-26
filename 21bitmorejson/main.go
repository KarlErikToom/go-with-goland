package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON video")
	EncodeJson()

}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJs Bootcamp", 299, "Learncodeonline.in", "password123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "Learncodeonline.in", "abc123", []string{"full-stack", "js"}},
		{"ANGULAR Bootcamp", 299, "Learncodeonline.in", "lampalamp123", nil},
	}

	//package this data as JSON data
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}
