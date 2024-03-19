package Bank

type Bank struct {
	count []int64
	max   uint
}

func Constructor(balance []int64) Bank {
	b := new(Bank)
	b.count = balance

	b.max = uint(len(balance))
	return *b
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if this.Withdraw(account1, money) {
		if this.Deposit(account2, money) {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func (this *Bank) Deposit(account int, money int64) bool {
	if account > 0 && uint(account) <= this.max {
		this.count[uint(account)-1] += money
		return true
	} else {
		return false
	}
}

func (this *Bank) Withdraw(account int, money int64) bool {
	if uint(account) <= this.max && this.userCount(uint(account)-1) >= money && account > 0 {
		this.count[uint(account)-1] -= money
		return true
	} else {
		return false
	}
}

func (this *Bank) userCount(u uint) int64 {
	return this.count[u]
}
