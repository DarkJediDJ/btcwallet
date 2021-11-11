package btcwallet

import (
	"fmt"
	"sync"
)

// BtcWallet is a representation of your bitcoin balance
type BtcWallet struct {
	sync.Mutex
	btc float64
}

// PrettyPrint displays your wallet balance but with some user-friendly style
func (bw *BtcWallet) PrettyPrint() string {
	return fmt.Sprintf("%.8v ₿", bw.btc)
}

// GetBalance displays your wallet balance
func (bw *BtcWallet) GetBalance() float64 {
	return bw.btc
}
