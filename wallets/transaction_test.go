package wallets

import (
	"sync"
	"testing"
)

func TestTransaction(t *testing.T) {
	var logs []TransactionLog
	var wg sync.WaitGroup

	amosWallet := Wallet{
		bal: Coin(500),
	}

	johnWallet := Wallet{
		bal: Coin(100),
	}

	transactions := []Transaction{
		{
			to:     &amosWallet,
			from:   &johnWallet,
			amount: 100,
		},
		{
			from:   &amosWallet,
			to:     &johnWallet,
			amount: 100,
		},
		{
			from:   &amosWallet,
			to:     &johnWallet,
			amount: 300,
		},
	}

	logCh := make(chan TransactionLog, len(transactions))
	wg.Add(len(transactions))

	for _, transaction := range transactions {
		go func(tx Transaction) {
			defer wg.Done()
			l, err := tx.Send()

			if err != nil {
				t.Error(err)
			}

			logCh <- l
		}(transaction)
	}

	wg.Wait()
	close(logCh)

	for log := range logCh {
		logs = append(logs, log)
	}

	if len(logs) != len(transactions) {
		t.Errorf("got %d want %d", len(logs), len(transactions))
	}

	t.Logf("Amos wallet has %s", amosWallet.Balance())
	t.Logf("John wallet has %s", johnWallet.Balance())
}
