
package account

import (
	"sync"
)

// Account represent a bank account
type Account struct {
	lock sync.Mutex
	money int64
	open bool
}

// Open creates a account
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}

	return &Account{money:initialDeposit, open:true}
}

// Close closes a account and get its money
func (A *Account) Close() (payout int64, ok bool) {

	A.lock.Lock()
	defer A.lock.Unlock()

	if !A.open {
		return 0, false
	}

	A.open = false
	return A.money, true
}

// Balance gets money in the account
func (A *Account) Balance() (balance int64, ok bool) {
	
	A.lock.Lock()
	defer A.lock.Unlock()

	if !A.open {
		return 0, false
	}

	return A.money, true
}

// Deposit makes deposits and withdrawals
func (A *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	
	A.lock.Lock()
	defer A.lock.Unlock()

	if !A.open {
		return A.money, false
	}

	if amount < 0 && A.money < -amount {
		return A.money, false
	}

	A.money += amount
	return A.money, true
}