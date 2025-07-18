package wallet

import (
	"errors"
	"fmt"
)

var ErrInsufficientBalance = errors.New("cannot withdraw, insufficient balance")

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return ErrInsufficientBalance
	}
	w.balance -= amount
	return nil
}
