package cli

import (
	"fmt"
	"log"
	"coincore"
)

func (cli *CLI) createBlockchain(address string) {
	if !coincore.ValidateAddress(address) {
		log.Panic("Error: address is not valid")
	}
	bc := coincore.CreateBlockchain(address)
	bc.Db.Close()
	fmt.Println("Done!")
}