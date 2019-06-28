package bank

// Package bank implements a bank with only one account

// This is supposed to produce race condition when called
// fromindependent threads.

// the account balance
var balance int

// Deposit puts the  given amount on the account
func Deposit(amount int) { balance = balance + amount }

// Balance retrieves the balance from account
func Balance() int { return balance }
