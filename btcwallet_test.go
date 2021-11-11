package btcwallet

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

var wg sync.WaitGroup

func TestDeposit(t *testing.T) {
	btcwallet := BtcWallet{btc: 56.4}
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(value float64, wg *sync.WaitGroup) {
			defer wg.Done()
			err := btcwallet.Deposit(value, wg)
			if err != nil {
				fmt.Println(err)
			}
		}(10.0, &wg)
	}
	wg.Wait()
	assert.Equal(t, 86.4, btcwallet.GetBalance(), "Deposit safety for concurrency useage")
	wg.Add(1)
	go func(value float64, wg *sync.WaitGroup) {
		defer wg.Done()
		err := btcwallet.Deposit(value, wg)
		if err != nil {
			fmt.Println(err)
		}
	}(-10.2, &wg)
	wg.Wait()
	assert.Equal(t, 86.4, btcwallet.GetBalance(), "Validation negative numbers")
	wg.Add(1)
	go func(value float64, wg *sync.WaitGroup) {
		defer wg.Done()
		err := btcwallet.Deposit(value, wg)
		if err != nil {
			fmt.Println(err)
		}
	}(0, &wg)
	wg.Wait()
	assert.Equal(t, 86.4, btcwallet.GetBalance(), "Validation zero number")
	assert.Nil(t, btcwallet.Deposit(1.1, &wg), "Deposit() should be nil")
}

func TestWithdraw(t *testing.T) {
	btcwallet := BtcWallet{btc: 76.5}
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(value float64, wg *sync.WaitGroup) {
			defer wg.Done()
			err := btcwallet.Withdraw(value, wg)
			if err != nil {
				fmt.Println(err)
			}
		}(10.0, &wg)
	}
	wg.Wait()
	assert.Equal(t, 46.5, btcwallet.GetBalance(), "Withdraw safety for concurrency useage")
	wg.Add(1)
	go func(value float64, wg *sync.WaitGroup) {
		defer wg.Done()
		err := btcwallet.Withdraw(value, wg)
		if err != nil {
			fmt.Println(err)
		}
	}(-10.2, &wg)
	wg.Wait()
	assert.Equal(t, 46.5, btcwallet.GetBalance(), "Validation negative numbers")
	wg.Add(1)
	go func(value float64, wg *sync.WaitGroup) {
		defer wg.Done()
		err := btcwallet.Withdraw(value, wg)
		if err != nil {
			fmt.Println(err)
		}
	}(0, &wg)
	wg.Wait()
	assert.Equal(t, 46.5, btcwallet.GetBalance(), "Validation zero number")
	assert.NotNil(t, btcwallet.Withdraw(100.1, &wg), "Withdraw more bitcoins then user have")
	assert.Nil(t, btcwallet.Withdraw(1.1, &wg), "Withdraw() should be nil")
}

func TestPrettyPrint(t *testing.T) {
	btcwallet := BtcWallet{btc: 76.5}
	assert.Equal(t, "76.5 ₿", btcwallet.PrettyPrint(), "Validation zero number")
}
