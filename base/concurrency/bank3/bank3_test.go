package bank3

import (
	"math/rand"
	"sync"
	"testing"
)

func TestBalance(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		Deposit(100)
		wg.Done()
	}()
	go func() {
		Deposit(1000)
		wg.Done()
	}()
	wg.Wait()
	if Balance() != 1100 {
		t.Fail()
	}
}

func TestWithDraw(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(40)

	go func() {
		for i := 0; i < 20; i++ {
			go func() {
				Deposit((rand.Intn(10) + 1) * 100)
				wg.Done()
			}()
		}
	}()

	go func() {
		for i := 0; i < 20; i++ {
			go func() {
				WithDraw((rand.Intn(10) + 1) * 100)
				wg.Done()
			}()
		}
	}()

	wg.Wait()
}
