package hello

import (
	"fmt"
)

func SayHello(name string) string {
	return fmt.Sprintf("Hello, %v!", name)
}

func SayBye(name string) string {
	return fmt.Sprintf("Bye, %v!", name)
}