/*
name  amount limit req
myName 5000   1k   5k
yourName 6000 2k  2k
HowMany transactions needed for withdraw Money from account

*/

package main

import "fmt"

type Account struct {
	Name   string
	amount int
	limit  int
}

type BankTransaction interface {
	Deposit(ac Account)
	Withdraw(moneyReq int) int
	CheckBalance(name string)
}

var accounts = make([]Account, 0)

func (a *Account) Deposit(ac Account) {
	a.Name = ac.Name
	a.amount = ac.amount
	a.limit = ac.limit
	accounts = append(accounts, *a)
	fmt.Println(a)
}

func (a *Account) CheckBalance(name string) {

	fmt.Println(a.amount)
}
func (a *Account) Withdraw(moneyReq int) int {
	if a.amount < moneyReq {
		return 0
	}
	with := 0
	count := 0
	temp := moneyReq
	for a.amount != 0 {
		if with >= moneyReq {
			break
		}
		if temp < a.limit {
			a.amount = a.amount - temp

		} else {
			a.amount = a.amount - a.limit

		}
		temp = temp - a.limit
		with += a.limit
		count++

	}
	return count
}

func main() {
	var myTr BankTransaction
	myTr = &Account{}
	myTr.Deposit(Account{
		Name:   "narendra",
		amount: 1000,
		limit:  1000000,
	})
	fmt.Println(myTr.Withdraw(5500))
	myTr.CheckBalance("narendra")

}
