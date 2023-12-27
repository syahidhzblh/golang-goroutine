package golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
)

var counter int = 0

func OnlyOnce() {
	counter++
}

func TestOnlyOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			once.Do(OnlyOnce)
			group.Done()
		}()
	}

	fmt.Println(counter)
}
