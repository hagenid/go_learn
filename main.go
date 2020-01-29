package main

import (
	"github.com/gookit/color"
	"main/cmd"
	"main/utils"
)

func main() {
	color.Red.Println("test")
	color.Red.Println(cmd.A)
	color.Green.Println(utils.R)
	color.Green.Println(utils.C)
}

// create new MyObject
