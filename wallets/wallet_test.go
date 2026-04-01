package wallets

import (
	"testing"
)

func TestWallet(t *testing.T) {
	w := Wallet{}

	assert := func(t testing.TB, wallet Wallet, want Coin) {
		t.Helper()

		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assetError := func(t testing.TB, err, want error) {
		t.Helper()

		if err == nil {
			t.Fatal("wanted an error but didn't get one")
		}

		if err != want {
			t.Errorf("got %q, want %q", err, want)
		}
	}

	t.Run("deposit 100 BTC into my wallet", func(t *testing.T) {
		w.Deposit(Coin(100))
		assert(t, w, Coin(100))
	})

	t.Run("withdraw 40 BTC from my wallet", func(t *testing.T) {
		w.Withdraw(Coin(40))
		assert(t, w, Coin(60))
	})

	t.Run("can not withdraw more than balance", func(t *testing.T) {
		err := w.Withdraw(Coin(100))
		assetError(t, err, ErrInsufficientFunds)

		currBal := Coin(60)
		assert(t, w, currBal)
	})

	t.Run("can not deposit negative amount", func(t *testing.T) {
		err := w.Deposit(Coin(-1))
		assetError(t, err, NegativeAmountDeposit)
	})
}
