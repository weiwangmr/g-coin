package cli

import (
	"fmt"
	"log"
	"coincore"
)

func (cli *CLI) listAddresses() {
	wallets, err := coincore.NewWallets()
	if err != nil {
		log.Panic(err)
	}
	addresses := wallets.GetAddress()

	for _, address := range addresses {
		fmt.Println(address)
	}
}