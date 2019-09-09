package bank

import (
	"fmt"
	"testing"
	"time"
)

func TestDeposit(t *testing.T) {
	// Alice
	go func() {
		Deposit(200)
		time.Sleep(time.Millisecond)
		fmt.Println("balance = ", Balance())
	}()

	// Bob
	go Deposit(100)

	time.Sleep(time.Second)
}

func TestInit(t *testing.T) {
	Init()

	Deposit1(100)
	fmt.Printf("余额：%6d\n", Balance1())
}
