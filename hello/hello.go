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

func IsAdult(age int) bool {
	if age > 18 {
		return true
	} else {
		return false
	}
}