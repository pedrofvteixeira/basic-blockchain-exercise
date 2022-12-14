package core

import (
	"strings"

	"github.com/rs/zerolog/log"
)

var ledger []Block = make([]Block, 0) // TODO actual persistency

func init() {
	// ledger must never be empty
	if len(ledger) == 0 {

		genesis := Block{
			make([]Tx, 0),
			asHexString(nil, nil),
			"",
		}

		log.Debug().Msgf("adding genesis block %s", genesis)
		putBlockInLedger(genesis)
	}
}

func putBlockInLedger(block Block) error {
	ledger = append(ledger, block)
	return nil
}

func getLedgerBlocks() ([]Block, error) {
	return ledger, nil
}

func getLedgerBlock(hash string) (bool, Block, error) {
	for _, block := range ledger {
		if strings.EqualFold(hash, block.Hash) {
			return true /*ok*/, block, nil
		}
	}
	return false /*nok*/, Block{}, nil
}

func getLedgerLastBlock() Block {
	return ledger[len(ledger)-1]
}
