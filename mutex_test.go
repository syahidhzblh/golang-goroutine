package golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Jumlah Counter = ", x)
}

type BankAccount struct {
	mutex   sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.mutex.Lock()
	account.Balance = account.Balance + amount
	account.mutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.mutex.RLock()
	balance := account.Balance
	account.mutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Balance = ", account.GetBalance())
}
