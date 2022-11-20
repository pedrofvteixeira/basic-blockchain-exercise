package core

import (
	"blockchain/common"
	"encoding/json"
	"fmt"

	"github.com/robfig/cron"
	"github.com/rs/zerolog/log"
)

type Tx struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount uint64 `default:"0" json:"amount"`
}

type Block struct {
	Txs      []Tx   `json:"txs"`
	Hash     string `json:"hash"`
	Previous string `json:"previous"`
}

var blockTimeMillis int
var txPerBlock int
var currentBlock Block

func setup(params *common.Params) {
	blockTimeMillis = params.BlockTimeMillis
	txPerBlock = params.TxPerBlock

	currentBlock = Block{}
}

func Start(params *common.Params) {
	setup(params)

	blockCron := cron.New()
	cronTimer := fmt.Sprintf("@every %dms", blockTimeMillis)
	log.Debug().Msgf("block close set at ", cronTimer)
	blockCron.AddFunc(cronTimer, func() { closeBlock() })
	blockCron.Start()
}

func GetBlock(hash string) (bool, Block) {
	ok, block, err := getLedgerBlock(hash)
	if err != nil {
		log.Error().Msgf("error getting block by hash %s", hash, err)
		return false, block
	}

	return ok, block
}

func GetBlocks() (bool, []Block) {
	blocks, err := getLedgerBlocks()
	if err != nil {
		log.Error().Msgf("error getting blocks", err)
		return false, blocks
	}
	return true, blocks
}

func AddTx(tx Tx) bool {
	if len(currentBlock.Txs) >= txPerBlock {
		err := closeBlock()
		if err != nil {
			log.Error().Msgf("error putting tx %s in currentBlock %s", tx, currentBlock, err)
			return false
		}
	}

	currentBlock.Txs = append(currentBlock.Txs, tx)
	return true
}

func getBlockTxsAsByteArray() []byte {
	out, err := json.Marshal(currentBlock.Txs)
	if err != nil {
		log.Error().Msgf("error getting open block Txs", err)
	}
	return []byte(string(out))
}

func closeBlock() error {
	hash := common.AsHash(currentBlock.Previous).Sum(getBlockTxsAsByteArray())
	currentBlock.Hash = string(hash)

	log.Info().Msgf("closing block %s", currentBlock.Hash)
	err := putBlockInLedger(currentBlock)
	if err != nil {
		log.Error().Msgf("error getting open block Txs", err)
		return err
	}

	// open up a new block
	currentBlock = Block{
		make([]Tx, txPerBlock),
		"",
		getLedgerLastBlock().Hash,
	}

	return nil
}
