package common

import (
	"flag"
	"fmt"
)

const (
	DefaultTxPerBlock      int  = 10    // number of txs that fill up a block
	DefaultBlockTimeMillis int  = 60000 // 60 seconds for open block
	DefaultVerbose         bool = false
	DefaultPort            int  = 8080
)

type Params struct {
	TxPerBlock      int
	BlockTimeMillis int
	Verbose         bool
	Port            int
}

func init() {
	flag.Usage = func() {
		printFlagUsage(Version)
	}
}

func ParseParams() Params {

	var params Params

	flag.BoolVar(&params.Verbose, "verbose", DefaultVerbose, "True to output all request, logic and balance information .")
	flag.IntVar(&params.TxPerBlock, "txPerBlock", DefaultTxPerBlock, "Specify max. amount of tx that fill up a block.")
	flag.IntVar(&params.BlockTimeMillis, "blockTimeMillis", DefaultBlockTimeMillis, "Specify max. amount for an open block in millis.")
	flag.IntVar(&params.Port, "port", DefaultPort, "Webapp port")

	flag.Parse()

	return params
}

func printFlagUsage(version string) {
	fmt.Printf("blockchain-exercise v%v\n", version)
	fmt.Printf("Usage: \n\t./blockchain -txPerBlock <uint> -blockTimeMillis <uint> -port <uint> [-verbose]\n")
	fmt.Printf("Example: \n")
	fmt.Printf("\t./blockchain -txPerBlock 127 -blockTimeMillis 60000 -port 8080 -verbose\n")
}
