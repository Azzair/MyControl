package main

import (
	"fmt"
	logger "gtk/MyControl/Lib/applogger"
	"gtk/MyControl/Lib/testLib2"
	"time"
)

func main() {

	fmt.Println("main")
	//testLib1.PrintHello("hello from main")
	testLib2.PrintHello("hello from main")

	logger.Info("hello from logger")
	logger.Allert("logger allert!!")
	logger.Critical("logger critical!!")
	logger.Fatal("logger fattal!!")
	time.Sleep(3 * time.Second)

}

func init() {
	fmt.Println("init")
}
