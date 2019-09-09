package bank3

import (
	"fmt"
	"sync"
)

var (
	mutex   sync.Mutex
	balance int
)

func Deposit(amount int) {
	mutex.Lock()
	defer mutex.Unlock()
	deposit(amount)
	fmt.Printf("存款：%d, 账户余额%d\n", amount, Balance())
}

func Balance() int {
	return balance
}

func WithDraw(amount int) bool {
	mutex.Lock()
	defer mutex.Unlock()
	currentBalance := Balance()
	if Balance() < amount {
		fmt.Printf("余额%d不足，取款%d失败\n", currentBalance, amount)
		return false
	}
	deposit(-amount)
	fmt.Printf("当前余额%d， 取款%d，剩余余额%d\n", currentBalance, amount, Balance())
	return true
}

func deposit(amount int)  {
	balance += amount
}
