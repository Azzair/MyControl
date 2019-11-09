package testLib2

import "gtk/MyControl/Libs/testLib1"

func PrintHello(str string) {
	testLib1.PrintHello(string("[TEST LIB 2] - " + str))
}

func init() {
	PrintHello("init")
}
