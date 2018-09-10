package coincore

import "bytes"

// TXInput represents a transaction input
type TXInput struct {
	Txid             []byte // input transaction id
	Vout 	         int    // input transaction TXOut[] index
	Signature        []byte
	PubKey           []byte
}

// UsesKey checks whether the address initiated the transaction 
func (in *TXInput) UsesKey(pubKeyHash []byte) bool {
	lockingHash := HashPubKey(in.PubKey)

	return bytes.Compare(lockingHash, pubKeyHash) == 0
}