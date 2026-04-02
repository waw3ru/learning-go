package wallets

import (
	"time"
)

type Transaction struct {
	to     *Wallet
	from   *Wallet
	amount Coin
}

type TransactionLog struct {
	amount Coin
	ts     time.Time
}

func (t *Transaction) Send() (log TransactionLog, e error) {
	if err := t.from.Withdraw(t.amount); err != nil {
		return TransactionLog{}, err
	}

	if err := t.to.Deposit(t.amount); err != nil {
		// rollback — re-deposit the amount back to sender
		t.from.Deposit(t.amount)
		return TransactionLog{}, err
	}

	log.amount = t.amount
	log.ts = time.Now()

	return log, nil
}
