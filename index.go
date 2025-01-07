package main

import (
	"fmt"

	"github.com/Mjkim-Programming/GoLearn/hello"
)

func main() {
	message := hello.SayHello("GaegulDev")
	fmt.Println(message)
}