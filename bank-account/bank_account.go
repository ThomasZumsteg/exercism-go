package account

import "sync"

//Account is a bank account
type Account struct {
	balance int64
	open    bool
	mu      *sync.Mutex
}

/*Open creates a bank account with a balance and doesn't allow
accounts to have a negative balance*/
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	account := Account{initialDeposit, true, &sync.Mutex{}}
	return &account
}

/*Close closes a bank account, doesn't work on already closed accounts*/
func (account *Account) Close() (int64, bool) {
	account.mu.Lock()
	defer account.mu.Unlock()
	if !account.open {
		return 0, false
	}
	balance := account.balance
	account.balance = 0
	account.open = false
	return balance, true
}

/*Balance shows the balance of a bank account, doesn't work on closed accounts*/
func (account *Account) Balance() (int64, bool) {
	account.mu.Lock()
	defer account.mu.Unlock()
	if !account.open {
		return 0, false
	}
	return account.balance, true
}

/*Deposit adds money to a bank account, doesn't allow negative balances
or closed accounts*/
func (account *Account) Deposit(amount int64) (int64, bool) {
	account.mu.Lock()
	defer account.mu.Unlock()
	if !account.open || account.balance < -amount {
		return 0, false
	}
	account.balance += amount
	return account.balance, true
}
