package coincore

import (
	"log"

	"github.com/boltdb/bolt"
)

// BlockchainIterator is used to  iterator over blockchain blocks
type BlockchainIterator struct {
	currentHash        []byte
	db                 *bolt.DB
}

// Next returns next block starting from the tip
func (i *BlockchainIterator) Next() *Block {
	var bc *Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BLOCKS_BUCKET))
		encodedBlock := b.Get(i.currentHash)
		bc = DeserializeBlock(encodedBlock)
		
		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	i.currentHash = bc.PrevBlockHash

	return bc
}