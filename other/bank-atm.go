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

type accounts map[string]*Account

func (a *accounts) Deposit(ac Account) {
	if *a == nil {
		*a = make(map[string]*Account)
	}
	(*a)[ac.Name] = &ac
	fmt.Println(a)
}

func (a *accounts) CheckBalance(name string) {
	val, ok := (*a)[name]
	if ok {
		fmt.Println(name+"'s current Balance is: ", val)
		return
	}
	fmt.Println("No User Found With Name " + name)
}
func (a *accounts) Withdraw(moneyReq int, name string) int {
	value, ok := (*a)[name]
	if !ok {
		fmt.Println("No User Found With Name " + name)
		return 0
	}
	if value.amount < moneyReq {
		return 0
	}
	with := 0
	count := 0
	temp := moneyReq
	for value.amount != 0 {
		if with >= moneyReq {
			break
		}
		if temp < value.limit {
			value.amount = value.amount - temp

		} else {
			value.amount = value.amount - value.limit

		}
		temp = temp - value.limit
		with += value.limit
		count++

	}
	return count
}

func main() {
	var myTr accounts
	myTr = make(accounts, 0)
	myTr.Deposit(Account{
		Name:   "narendra",
		amount: 1000000,
		limit:  1000,
	})
	fmt.Println("User has to do ", myTr.Withdraw(5500, "narendra"), " transactions to withdraw the money")
	myTr.CheckBalance("narendra")

}
