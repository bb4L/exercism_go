package account

import "sync"

// Account bank account
type Account struct {
	sync.Mutex
	closed  bool
	balance int64
}

// Open opens an account with the given balance
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{sync.Mutex{}, false, initialDeposit}
}

// Close closes an account
func (a *Account) Close() (payout int64, ok bool) {
	a.Lock()
	defer a.Unlock()

	if a.closed {
		payout, ok = 0, false
	} else {
		a.closed = true
		payout, ok = a.balance, true
	}
	return
}

// Balance retrieves the balance of an account
func (a *Account) Balance() (balance int64, ok bool) {
	a.Lock()
	defer a.Unlock()

	if a.closed {
		balance, ok = 0, false
	} else {
		balance, ok = a.balance, true
	}

	return
}

// Deposit deposits the given amount
func (a *Account) Deposit(amount int64) (balance int64, ok bool) {
	a.Lock()
	defer a.Unlock()

	if a.balance+amount < 0 || a.closed {
		balance, ok = 0, false
	} else {
		a.balance += amount
		balance, ok = a.balance, true
	}

	return
}
