package main

import (
	"fmt"
	"log"
)

type Account struct {
	balance float32
}

func (a *Account) Withdraw(amount float32) float32 {
	log.Printf("Withdraw %f", amount)

	if (a.balance - amount) < 0 {
		log.Printf("Not enough money")

		return a.balance
	}

	a.balance -= amount

	return a.balance
}

func (a *Account) Deposit(amount float32) float32 {
	log.Printf("Deposit %f", amount)

	a.balance += amount

	return a.balance
}

func (a *Account) Balance() float32 {
	return a.balance
}

func main() {
	acc := Account{}

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				if j%2 == 1 {
					acc.Withdraw(100)
					continue
				}

				acc.Deposit(100)
			}
		}()
	}

	fmt.Scanln()

	println(acc.Balance())

	closure()

}

func closure() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("got", i)
		}()
	}

	fmt.Scanln()
}
