package main

import (
	"errors"
	"fmt"
)

type AccountSync struct {
	balance      float64
	deltaChain   chan float64
	balanceChain chan float64
	errorChain   chan error
}

func NewAccountSync(balance float64) (a *AccountSync) {
	a = &AccountSync{
		balance:      balance,
		deltaChain:   make(chan float64),
		balanceChain: make(chan float64),
		errorChain:   make(chan error, 1),
	}

	go a.run()
	return
}

func (a *AccountSync) Balance() float64 {
	return <-a.balanceChain
}

func (a *AccountSync) Deposit(amount float64) error {
	a.deltaChain <- amount
	return <-a.errorChain
}

func (a *AccountSync) Withdraw(amount float64) error {
	a.deltaChain <- -amount
	return <-a.errorChain
}

func (a *AccountSync) applyDelta(delta float64) error {
	stateStr := "Deposit on"
	if delta < 0 {
		stateStr = "Withdraw on"
	}

	fmt.Println(stateStr, delta)

	newBalance := a.balance + delta

	if newBalance < 0 {
		return errors.New("Not enough money")
	}

	a.balance = newBalance
	return nil
}

func (a *AccountSync) run() {
	var delta float64
	for {
		select {
		case delta = <-a.deltaChain:
			a.errorChain <- a.applyDelta(delta)

		case a.balanceChain <- a.balance:

		}
	}
}

func main() {
	acc := NewAccountSync(20)

	for i := 0; i < 100; i++ {

		go func() {
			for j := 0; j < 10; j++ {
				if j%2 == 1 {
					acc.Withdraw(50)
					//err := acc.Withdraw(50)
					//if err != nil {
					//	return
					//}
					continue
				}

				acc.Deposit(50)
				//err := acc.Deposit(50)
				//if err != nil {
				//	return
				//}
			}
		}()
	}

	_, err := fmt.Scanln()
	if err != nil {
		return
	}
	fmt.Println("Balance=", acc.Balance())
}
