package bank

import (
	"train/monitor/economy/shop"
	"train/monitor/monitorfile"
)

type Bank struct {
	Rating           int           // 评级
	LoanInterestRate int           // 贷款利率
	InterestRate     int           // 存款利率
	Saving           int           // 存款
	Loans            map[int]*Loan // 贷款实例
	Trust            int           // 信用,当-2时，rating--,+2时Rating++
}

type Loan struct {
	Money        int           // 贷款金额
	Days         int           // 天数
	NeedPay      int           // 需要还的钱
	Rating       int           // 借贷所花的
	MortgageList []*shop.Cards // 抵押的英雄们
}

func BankInit() *Bank {
	return &Bank{
		Rating:           monitorfile.EconomyMapToint("初始信用"),
		LoanInterestRate: monitorfile.EconomyMapToint("初始贷款利率"),
		InterestRate:     monitorfile.EconomyMapToint("初始利率"),
		Saving:           0,
		Loans:            map[int]*Loan{},
		Trust:            0,
	}
}

func StrToRating(str string) int {
	switch str {
	case "D":
		return 1
	case "C":
		return 2
	case "B":
		return 3
	case "A":
		return 4
	case "AA":
		return 5
	case "AAA":
		return 6
	default:
		return 0
	}
}
func intToRating(u int) string {
	switch u {
	case 1:
		return "D"
	case 2:
		return "C"
	case 3:
		return "B"
	case 4:
		return "A"
	case 5:
		return "AA"
	case 6:
		return "AAA"
	default:
		return "E"
	}
}
func StrToLoanRate(str string) int {
	switch str {
	case "D":
		return 50
	case "C":
		return 30
	case "B":
		return 25
	case "A":
		return 20
	case "AA":
		return 15
	case "AAA":
		return 15
	default:
		return 0
	}
}
func StrToInterest(str string) int {
	switch str {
	case "D":
		return 0
	case "C":
		return 0
	case "B":
		return 5
	case "A":
		return 5
	case "AA":
		return 10
	case "AAA":
		return 15
	default:
		return 0
	}
}

// 存钱
func (b *Bank) Save(money int) {
	// 如果有钱
	b.Saving += money
}

// 关于取钱
func (b *Bank) Take(money int) bool {
	if b.Rating <= StrToRating("D") {
		return false
	} else {
		if money > b.Saving {
			return false
		} else {
			b.Saving -= money
			return true
		}
	}
}

// 关于还贷
func (b *Bank) Pay(loan int) bool {
	if b.Loans[loan].NeedPay > b.Saving {
		return false
	} else {
		b.Saving -= b.Loans[loan].NeedPay
		b.Trust++
		b.CheckTrust()
		b.Loans[loan] = nil
		return true
	}
}

// 关于借贷
// CreateLoan 当Rating>=AA 时，可以总共有两笔贷款 当Rating==AAA时，可以总共有三笔
func (b *Bank) MakeLoan(Money int, cards []*shop.Cards) bool {
	// first loan
	ninetine := (Money * 9) / 10
	if len(b.Loans) == 0 && b.Rating > 0 {
		loan := MakeLoan(Money, b.Rating)
		b.Saving += ninetine
		b.Loans[int(len(b.Loans))] = loan
		loan.MortgageList = cards
		return true
	}
	// second loan
	if len(b.Loans) >= 1 && b.Rating > StrToRating("AA") {
		loan := MakeLoan(Money, b.Rating)
		b.Saving += ninetine
		b.Loans[int(len(b.Loans))] = loan
		loan.MortgageList = cards
		return true
	}
	if len(b.Loans) == 2 && b.Rating == StrToRating("AAA") {
		loan := MakeLoan(Money, b.Rating)
		b.Saving += ninetine
		b.Loans[int(len(b.Loans))] = loan
		loan.MortgageList = cards
		return true
	}
	return false
}

func MakeLoan(money int, rating int) *Loan {

	return &(Loan{
		Money:   money,
		Days:    0,
		NeedPay: money,
		Rating:  rating,
	})
}

func (b *Bank) DayPast() {
	// SAVING++
	if b.Saving != 0 {
		interest := (b.Saving * b.InterestRate) / 100
		b.Saving += interest
	}
	//fmt.Println(interest)
	// loan++
	for _, loan := range b.Loans {
		// loan money ++
		// loan
		loan.NeedPay += loan.Money * StrToLoanRate(intToRating(loan.Rating)) / 100
		if loan.NeedPay >= loan.Money*2 {
			b.Trust--
			b.CheckTrust()
		}
	}
}

func (b *Bank) CheckTrust() {
	if b.Trust <= -2 {
		if b.Rating != 0 {
			b.Rating--
		}
		b.Trust = 0
	}
	if b.Trust >= 2 {
		if b.Rating != StrToRating("AAA") {
			b.Rating++
		}
		b.Trust = 0
	}
}
