package main

import (
	"blockchain/common"
	"blockchain/core"
	"blockchain/web"

	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	p := common.ParseParams()

	setup(p)
	start(p)
}

func setup(params common.Params) {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		NoColor:    false,
		TimeFormat: common.DateFormat,
	}) // pretty logs, albeit less performant
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if params.Verbose {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}

func start(p common.Params) {
	log.Info().Msg("starting core...")
	core.Start(p.TxPerBlock, p.BlockTimeMillis)

	log.Info().Msg("starting web...")
	web.Start(p.Port, p.Verbose)
}
