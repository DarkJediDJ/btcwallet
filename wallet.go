package btcwallet

import (
	"strconv"
	"sync"
)

// BtcWallet is a representation of your bitcoin balance
type BtcWallet struct {
	sync.Mutex
	btc float64
}

// PrettyPrint displays your wallet balance but with some user-friendly style
func (bw *BtcWallet) PrettyPrint() string {
	return "Your balance is " + strconv.FormatFloat(bw.btc, 'E', -1, 64) + "₿"
}

// GetBalance displays your wallet balance
func (bw *BtcWallet) GetBalance() float64 {
	return bw.btc
}
