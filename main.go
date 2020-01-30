package main

import (
	"github.com/gookit/color"
	"main/cmd"
	"main/utils"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	isAdmin bool   `json:"isadmin"`
}

func main() {


	
	color.Red.Println("test")
	color.Red.Println(cmd.A)
	color.Green.Println(utils.R)
	color.Green.Println(utils.C)

	p := Person{
		Name:    "Mike",
		Age:     16,
		isAdmin: false,
	}

	color.Red.Println(p, p.Name, p.Age, p.isAdmin)

}

// create new MyObject
