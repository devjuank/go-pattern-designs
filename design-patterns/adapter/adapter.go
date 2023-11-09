package main

import "fmt"

type Payment interface {
	Pay()
}

type CashPayment struct{}
type BankPayment struct{}
type BankPaymentAdapter struct {
	BankPayment *BankPayment
	backAccount string
}

func (c *CashPayment) Pay() {
	fmt.Println("do cash payment")
	return
}

func ProcessPayment(p Payment) {
	p.Pay()
}

func (BankPayment) Pay(backAccount string) {
	fmt.Printf("do Bank payment account:%s\n", backAccount)
}

func (b BankPaymentAdapter) Pay() {
	b.BankPayment.Pay(b.backAccount)
}
func main() {
	cash := &CashPayment{}
	ProcessPayment(cash)

	bank := &BankPaymentAdapter{
		BankPayment: &BankPayment{},
		backAccount: "123",
	}
	ProcessPayment(bank)
}
