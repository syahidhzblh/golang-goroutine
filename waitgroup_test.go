package golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronus(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)

	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	x := 0
	for i := 0; i < 100; i++ {
		go RunAsynchronus(group)
		x++
	}

	group.Wait()
	fmt.Println("Complete")
	fmt.Println("Total Hello", x)
}
