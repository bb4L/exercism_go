package account

import "sync"

// Account bank account
type Account struct {
	open    bool
	balance int64
	mutex   *sync.Mutex
}

// Open opens an account with the given balance
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{true, initialDeposit, &sync.Mutex{}}
}

// Close closes an account
func (a *Account) Close() (payout int64, ok bool) {
	a.mutex.Lock()
	if !a.open {
		payout, ok = 0, false
	} else {
		a.open = false
		payout, ok = a.balance, true
	}

	a.mutex.Unlock()
	return
}

// Balance retrieves the balance of an account
func (a *Account) Balance() (balance int64, ok bool) {
	a.mutex.Lock()
	if !a.open {
		balance, ok = 0, false
	} else {
		balance, ok = a.balance, true
	}

	a.mutex.Unlock()
	return
}

// Deposit deposits the given amount
func (a *Account) Deposit(amount int64) (balance int64, ok bool) {
	a.mutex.Lock()
	if a.balance+amount < 0 || !a.open {
		balance, ok = 0, false
	} else {
		a.balance += amount
		balance, ok = a.balance, true
	}

	a.mutex.Unlock()
	return
}
