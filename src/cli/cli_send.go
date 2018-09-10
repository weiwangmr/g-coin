package cli

import (
	"fmt"
	"log"
	"coincore"
)

func (cli *CLI) send(from, to string, amount int) {
	if !coincore.ValidateAddress(from) {
		log.Panic("ERROR: Sender address is not valid")
	}
	if !coincore.ValidateAddress(to) {
		log.Panic("ERROR: Recipient address is not valid")
	}

	bc := coincore.NewBlockchain()
	UTXOSet := coincore.UTXOSet{bc}
	defer bc.Db.Close()

	tx := coincore.NewUTXOTransaction(from, to, amount, &UTXOSet)
	cbTx := coincore.NewCoinbaseTX(from, "")
	txs := []*coincore.Transaction{cbTx, tx}

	newBlock := bc.MineBlock(txs)
	UTXOSet.Update(newBlock)
	fmt.Println("Success!")

}