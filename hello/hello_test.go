package hello_test

import (
	"github.com/Mjkim-Programming/GoLearn/hello"

	"testing"
	"fmt"
)

func TestSayHello(t *testing.T) {
	// Given
	name := "KMJ"
	expected := fmt.Sprintf("Hello, %v!", name)

	// When
	res := hello.SayHello(name)

	// Then
	if res != expected {
		t.Error(fmt.Sprintf("Wrong Result : Expected %v, got %v", expected, res))
	}
}

func TestSayBye(t *testing.T) {
	// Given
	name := "KMJ"
	expected := fmt.Sprintf("Bye, %v!", name)

	// When
	res := hello.SayBye(name)

	// Then
	if res != expected {
		t.Error(fmt.Sprintf("Wrong Result : Expected %v, got %v", expected, res))
	}
}