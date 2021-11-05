package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

var wg sync.WaitGroup

func TestDeposit(t *testing.T) {
	btcwallet := BtcWallet{btc: 56.4}
	wg.Add(3)
	go func(value float64, wg *sync.WaitGroup) {
		defer wg.Done()
		btcwallet.Deposit(value, wg)
	}(10.1, &wg)
	go func(value float64, wg *sync.WaitGroup) {
		defer wg.Done()
		btcwallet.Deposit(value, wg)
	}(12.1, &wg)
	go func(value float64, wg *sync.WaitGroup) {
		defer wg.Done()
		btcwallet.Deposit(value, wg)
	}(15.0, &wg)
	wg.Wait()
	assert.Equal(t, 93.6, btcwallet.GetBalance(), "Deposit safety for concurrency useage")
	wg.Add(1)
	go func(value float64, wg *sync.WaitGroup) {
		defer wg.Done()
		btcwallet.Deposit(value, wg)
	}(-10.2, &wg)
	wg.Wait()
	assert.Equal(t, 93.6, btcwallet.GetBalance(), "Validation negative numbers")
	wg.Add(1)
	go func(value float64, wg *sync.WaitGroup) {
		defer wg.Done()
		btcwallet.Deposit(value, wg)
	}(0, &wg)
	wg.Wait()
	assert.Equal(t, 93.6, btcwallet.GetBalance(), "Validation zero number")
	assert.Nil(t, btcwallet.Deposit(1.1, &wg), "Deposit() should be nil")
}

func TestWithdraw(t *testing.T) {
	btcwallet := BtcWallet{btc: 76.5}
	wg.Add(3)
	go func(value float64, wg *sync.WaitGroup) {
		defer wg.Done()
		btcwallet.Withdraw(value, wg)
	}(10.1, &wg)
	go func(value float64, wg *sync.WaitGroup) {
		defer wg.Done()
		btcwallet.Withdraw(value, wg)
	}(12.1, &wg)
	go func(value float64, wg *sync.WaitGroup) {
		defer wg.Done()
		btcwallet.Withdraw(value, wg)
	}(15.2, &wg)
	wg.Wait()
	assert.Equal(t, 39.4, btcwallet.GetBalance(), "Withdraw safety for concurrency useage")
	wg.Add(1)
	go func(value float64, wg *sync.WaitGroup) {
		defer wg.Done()
		btcwallet.Withdraw(value, wg)
	}(-10.2, &wg)
	wg.Wait()
	assert.Equal(t, 39.4, btcwallet.GetBalance(), "Validation negative numbers")
	wg.Add(1)
	go func(value float64, wg *sync.WaitGroup) {
		defer wg.Done()
		btcwallet.Withdraw(value, wg)
	}(0, &wg)
	wg.Wait()
	assert.Equal(t, 39.4, btcwallet.GetBalance(), "Validation zero number")
	assert.NotNil(t, btcwallet.Withdraw(100.1, &wg), "Withdraw more bitcoins then user have")
	assert.Nil(t, btcwallet.Withdraw(1.1, &wg), "Withdraw() should be nil")
}
