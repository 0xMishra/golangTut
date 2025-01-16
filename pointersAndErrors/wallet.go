package pointersanderrors

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Stringer interface {
	String() string
}

type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(m Bitcoin) {
	w.balance += m
}

var ErrInsufficientFunds = errors.New("insufficient funds")

func (w *Wallet) Withdraw(m Bitcoin) error {
	if w.balance < m {
		return ErrInsufficientFunds
	}
	w.balance -= m
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
