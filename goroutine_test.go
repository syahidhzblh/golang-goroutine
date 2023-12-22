package golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func DisplayTime(number int) {
	fmt.Println("Display", number)
}

func TestDisplayTime(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayTime(i)
	}

	time.Sleep(5 * time.Second)
}

func RunSayHello() {
	fmt.Println("Hello World")
}

func TestHelloWorld(t *testing.T) {
	go RunSayHello()
	fmt.Println("Hello After goroutine")

	time.Sleep(5 * time.Second)
}
