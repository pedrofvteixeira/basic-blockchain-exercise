package blockchain

import (
	"blockchain/common"
	"blockchain/core"
	"blockchain/web"
)

func main() {
	p := common.ParseParams()
	core.Start(p)
	web.Start(p)
}
