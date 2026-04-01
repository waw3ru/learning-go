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

	assetError := func(t testing.TB, err error) {
		t.Helper()
		if err == nil {
			t.Error("wanted an error but didn't get one")
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
		assetError(t, err)

		currBal := Coin(60)
		assert(t, w, currBal)
	})
}
