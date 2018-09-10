package cli

import (
	"fmt"
	"log"
	"coincore"
	"utils"
)

func (cli *CLI) getBalance(address string) {
	if !coincore.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := coincore.NewBlockchain()
	UTXOSet := coincore.UTXOSet{bc}
	defer bc.Db.Close()

	balance := 0
	pubKeyHash := utils.Base58Decode([]byte(address))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	UTXOs := UTXOSet.FindUTXO(pubKeyHash)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of '%s': %d\n", address, balance)
}