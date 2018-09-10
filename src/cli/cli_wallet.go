package cli 

import (
	"fmt"
	"coincore"
)

func (cli *CLI) createWallet() {
	wallets, _ := coincore.NewWallets()
	address := wallets.CreateWallet()
	wallets.SaveToFile() 

	fmt.Printf("Your new address: %s\n", address)
}