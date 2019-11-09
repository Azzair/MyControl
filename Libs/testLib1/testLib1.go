package testLib1

import "fmt"

func PrintHello(str string) {
	fmt.Printf("[TEST LIB 1] - %s\n", str)
}

func init() {
	PrintHello("init")
}
