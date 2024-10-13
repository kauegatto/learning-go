package main

import "errors"

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(value int) {
	w.balance = w.balance + value
}

func (w Wallet) Balance() int {
	return w.balance
}

func (w *Wallet) Withdraw(value int) error {
	if value > w.Balance() {
		return errors.New("error")
	}
	
	w.balance -= value
	return nil
}

func main() {

}
