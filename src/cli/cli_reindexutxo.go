package cli 

import (
	"fmt"
	"coincore"
)

func (cli *CLI) reindexUTXO() {
	bc := coincore.NewBlockchain()
	UTXOSet := coincore.UTXOSet{bc}
	UTXOSet.Reindex()

	count := UTXOSet.CountTransactions()
	fmt.Printf("Done! There are %d transactions in the UTXO set.\n", count)
}