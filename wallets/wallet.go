package wallets

import (
	"errors"
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

func (w *Wallet) Deposit(amount Coin) {
	w.bal += amount
}

func (w *Wallet) Balance() Coin {
	return w.bal
}

func (w *Wallet) Withdraw(amount Coin) error {
	if amount > w.bal {
		return errors.New("cannot withdraw more than balance")
	}
	w.bal -= amount
	return nil
}
