package web

import (
	"blockchain/common"

	"github.com/gin-gonic/gin"

	"github.com/rs/zerolog/log"

	"blockchain/core"
	"encoding/json"
	"net/http"
	"strconv"
)

var router *gin.Engine
var port int

func setup(params *common.Params) {

	gin.SetMode(gin.ReleaseMode)
	if params.Verbose {
		gin.SetMode(gin.DebugMode)
	}

	port = int(params.Port)
	router = gin.Default()
}

func Start(params *common.Params) {
	setup(params)

	root := router.Group("/")
	{
		root.GET("/blocks", getAllBlocks)
		root.GET("/blocks/:hash", getBlockByHash)
		root.GET("/tx/:from/:to/:amount", postTx) // TODO change to POST
	}

	log.Info().Msgf("webapp listening on port %d", port)
	router.Run(strconv.Itoa(port))
}

func getAllBlocks(ctx *gin.Context) {
	ok, blocks := core.GetBlocks()
	reply(ctx, ok, blocks)
}

func getBlockByHash(ctx *gin.Context) {
	ok, block := core.GetBlock(ctx.Param("hash"))
	reply(ctx, ok, block)
}

func postTx(ctx *gin.Context) {

	amountVal, err := strconv.ParseInt(ctx.Param("amount"), 10, 64)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	newTx := core.Tx{
		From:   ctx.Param("from"),
		To:     ctx.Param("to"),
		Amount: uint64(amountVal),
	}

	ok := core.AddTx(newTx)
	reply(ctx, ok, nil)
}

func reply(ctx *gin.Context, ok bool, content any) {
	if !ok {
		ctx.String(http.StatusInternalServerError, "{\"result\": \"nok\"}")

	} else if content == nil {
		ctx.String(http.StatusOK, "{\"result\": \"ok\"}")

	} else {

		out, err2 := json.Marshal(content)
		if err2 != nil {
			ctx.String(http.StatusInternalServerError, err2.Error())
		} else {
			ctx.String(http.StatusOK, string(out))
		}
	}
}
