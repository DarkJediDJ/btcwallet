package btcwallet

import (
	"errors"
	"sync"
)

// Deposit provides way to add value to your btcwallet
func (bw *BtcWallet) Deposit(value float64, wg *sync.WaitGroup) error {
	bw.Lock()
	defer bw.Unlock()
	if value <= 0 {
		return errors.New("Youre trying to deposit negative or zero value")
	}
	bw.btc += value
	return nil
}
