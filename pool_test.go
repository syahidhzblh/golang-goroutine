package golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Pool biasanya digunakan untuk case koneksi ke database
func TestPool(t *testing.T) {
	pool := sync.Pool{
		//Default Value
		New: func() interface{} {
			return "New"
		},
	}

	// value := []struct {
	// 	Name string
	// }{
	// 	{Name: "Syahid"},
	// 	{Name: "Hisbul"},
	// 	{Name: "Ganteng"},
	// }

	// for _, values := range value {
	// 	pool.Put(&values.Name)
	// }

	pool.Put("Syahid")
	pool.Put("Hisbul")
	pool.Put("Ganteng")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get() //Get Data di Pool
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data) //Balikan data ke Pool after Get
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Selesai")
}
