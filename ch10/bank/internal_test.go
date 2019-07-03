package bank

import (
	"fmt"
	"testing"
	"time"
)

func TestBankDeposit(t *testing.T) {
	go Deposit(100)
}

func TestBankAsyncDepositAndPrint(t *testing.T) {
	go func() {
		Deposit(200)
		fmt.Println("=", Balance())
	}()
}

// This test case is supposed to produce a race condition
func TestRace(t *testing.T) {
	go func() {
		d := 200
		Deposit(d)
		time.Sleep(100 * time.Millisecond)
		b := Balance()
		if d != b {
			t.Error("Expected,", d, "got", b)
		}
		fmt.Println("=", b)
	}()
	go Deposit(100)
	// give some time to completet async processing
	time.Sleep(200 * time.Millisecond)
}
