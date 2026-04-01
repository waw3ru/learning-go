package wallets

import (
	"fmt"

	"github.com/waw3ru/learning-go/utils"
)

type Coin float64

func (b Coin) String() string {
	return fmt.Sprintf("%.2f BTC", b)
}

/* *
 * Always round off -> .2
 * */
func (b Coin) Float64() float64 {
	return utils.RoundTo(float64(b), 2)
}

type Wallet struct {
	bal Coin
}

func (w *Wallet) Deposit(amount Coin) *utils.CustomError {
	if amount < 1 {
		return utils.ThrowErr(NegativeAmountDeposit.Error(), "wrong amount entered", nil)
	}
	w.bal += amount
	return nil
}

func (w *Wallet) Balance() Coin {
	return w.bal
}

func (w *Wallet) Withdraw(amount Coin) *utils.CustomError {
	if amount > w.bal {
		return utils.ThrowErr(ErrInsufficientFunds.Error(), "cannot withdraw, insufficient funds", nil)
	}
	w.bal -= amount
	return nil
}
