package main

import (
	"fmt"
	"gtk/MyControl/Libs/testLib2"
)

func main() {
	fmt.Println("main")
	//testLib1.PrintHello("hello from main")
	testLib2.PrintHello("hello from main")
}

func init() {
	fmt.Println("init")
}
