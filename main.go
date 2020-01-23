package main

import (
	"fmt"
	"os"
)

func main() {
	// initialise verable
	text := "Hello Gold!"
	filename := "hello.txt"

	// Testing function return result
	createfile(5, 5)
	// Chek existing file, if existing function is close
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		fmt.Println("Такой файл уже существует:", filename)
		os.Exit(0)
	} else {
		fmt.Println("test1")
	}

	file, err := os.Create(filename)

	if err != nil {
		fmt.Println("Unable to create file:", err)
		return
	}
	defer file.Close()
	file.WriteString(text)
	fmt.Println("Test")
	fmt.Println("Done.")

}

func createfile(a int, b int) {
	z := a + b
	fmt.Println(z)

}
