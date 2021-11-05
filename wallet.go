package main

import (
	"sync"
)

//BtcWallet is a representation of your bitcoin balance
type BtcWallet struct {
	sync.Mutex
	btc float64
}

func main() {
}

//GetBalance displays your wallet balance
func (bw *BtcWallet) GetBalance() float64 {
	return bw.btc
}
