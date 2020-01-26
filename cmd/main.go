package main

import (
	"fmt"
	"github.com/gookit/color"
	"os"
	"time"
)

func main() {
	// initialise verable
	text := "Hello Gold!"
	filename := "hello.txt"
	color.Cyan.Printf("Simple to use %s\n", "color")
	ctime := time.RFC822
	fmt.Printf("current unix timestamp is :%v\n", ctime)
	// Testing
	//function return result
	createfile(0, 50)
	// Chek existing file, if existing function is close
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		color.BgGray.Println("Такой файл уже существует:", filename)
		os.Exit(0)
	} else {
		//fmt.Println("test1")
		color.Yellow.Println("test1")
	}

	file, err := os.Create(filename)

	if err != nil {
		color.Red.Println("Unable to create file:", err)
		return
	}
	defer file.Close()
	file.WriteString(text)
	fmt.Println("Test")
	fmt.Println("Done.")
	color.Red.Println("Done.")

}

func addfile(fileexist string) {

}

func createfile(a int, b int) {

	if a == 5 {
		z := a + b
		fmt.Println(z)
	} else {
		k := a - b
		fmt.Println(k)
	}

}
