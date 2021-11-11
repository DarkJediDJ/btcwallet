package btcwallet

import (
	"errors"
	"sync"
)

// Withdraw provides way to subtract value to your btcwallet
func (bw *BtcWallet) Withdraw(value float64, wg *sync.WaitGroup) error {
	bw.Lock()
	defer bw.Unlock()
	if value <= 0 {
		return errors.New("Youre trying to withdraw negative or zero value")
	}
	if bw.btc < value {
		return errors.New("You cant withdraw more mouney then you have")
	}
	bw.btc -= value
	return nil
}
