package wallets

import "errors"

// Error Constants
var (
	ErrInsufficientFunds  = errors.New("cannot withdraw, insufficient funds")
	NegativeAmountDeposit = errors.New("wrong amount entered")
)
