package golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()

	for value := range ticker.C {
		fmt.Println(value)
	}

}

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)
	go func() {
		time.Sleep(5 * time.Second)
	}()

	for value := range channel {
		fmt.Println(value)
	}

}
