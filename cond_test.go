package golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var cond = sync.NewCond(&sync.Mutex{})
var group = sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1)

	//Alur Lock -> Wait -> Execute -> Unlock
	cond.L.Lock()
	cond.Wait()
	fmt.Println("Done", value)
	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}

	//Give Signal after 1 second not need to wait (After 1 second, will execute other goroutine)
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		cond.Signal()
	}

	// for i := 0; i < 10; i++ {
	// 	time.Sleep(1 * time.Second)
	// 	cond.Broadcast()
	// }

	group.Wait()
}
