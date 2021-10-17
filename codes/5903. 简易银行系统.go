type Bank struct {
	balance []int64
}

func Constructor(balance []int64) Bank {
	return Bank{balance: balance}
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if account1 > len(this.balance) || account2 > len(this.balance) {
		return false
	} else if this.balance[account1-1] < money {
		return false
	} else {
		this.balance[account1-1] -= money
		this.balance[account2-1] += money
		return true
	}
}

func (this *Bank) Deposit(account int, money int64) bool {
	if account > len(this.balance) {
		return false
	} else {
		this.balance[account-1] += money
		return true
	}
}

func (this *Bank) Withdraw(account int, money int64) bool {
	if account > len(this.balance) || this.balance[account-1] < money {
		return false
	} else {
		this.balance[account-1] -= money
		return true
	}
}