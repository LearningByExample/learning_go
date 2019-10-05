package wallet

import (
	"errors"
	"fmt"
)

//Bitcoin is an int
type Bitcoin int

//Wallet that hold bitcoins
type Wallet struct {
	balance Bitcoin
}

//Stringer something into an string
type Stringer interface {
	String() string
}

//String representation of bitcin
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

//Deposit some amount in our wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

//Balance of our wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

//ErrInsufficientFunds nto having enough founds
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

//Withdraw an amount from our walllet
func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
