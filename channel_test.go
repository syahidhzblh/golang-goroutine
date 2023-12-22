package golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	//Membuat Channel
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		//Mengirim value ke channel
		channel <- "Syahid"
		fmt.Println("Selesai mengirim channel")
	}()
	// Menerima value dari channel, lalu disimpan ke variable data dan print
	data := <-channel
	fmt.Println(data)
	time.Sleep(10 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Syahid"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)
	data := <-channel
	fmt.Println(data)

	time.Sleep(10 * time.Second)
}
