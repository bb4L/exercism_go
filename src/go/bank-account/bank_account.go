package account

import "sync"

// Account bank account
type Account struct {
	sync.RWMutex
	closed  bool
	balance int64
}

// Open opens an account with the given balance
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{balance: initialDeposit}
}

// Close closes an account
func (a *Account) Close() (int64, bool) {
	a.Lock()
	defer a.Unlock()

	if a.closed {
		return 0, false
	}

	a.closed = true
	return a.balance, true
}

// Balance retrieves the balance of an account
func (a *Account) Balance() (int64, bool) {
	a.RLock()
	defer a.RUnlock()

	if a.closed {
		return 0, false
	}

	return a.balance, true
}

// Deposit deposits the given amount
func (a *Account) Deposit(amount int64) (int64, bool) {
	a.Lock()
	defer a.Unlock()

	if a.balance+amount < 0 || a.closed {
		return 0, false
	}

	a.balance += amount
	return a.balance, true
}
