package web

import (
	"blockchain/common"

	"github.com/gin-gonic/gin"

	"github.com/rs/zerolog/log"

	"blockchain/core"
	"encoding/json"
	"net/http"
	"net/url"
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

	r := router.Group("/")
	{
		r.GET("/", redirect)
		r.GET("/blocks", getAllBlocks)
		r.GET("/blocks/:hash", getBlockByHash)
		r.GET("/tx/:from/:to/:amount", postTx) // TODO change to POST
	}

	log.Info().Msgf("webapp listening on port %d", port)
	router.Run(":" + strconv.Itoa(port))
}

func redirect(ctx *gin.Context) {
	location := url.URL{Path: "/blocks"}
	ctx.Redirect(http.StatusFound, location.RequestURI())
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
	if content == nil {
		ctx.JSON(http.StatusOK, gin.H{"result": ok})

	} else {

		out, err := json.Marshal(content)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			ctx.Writer.Write(out)
		}
	}
}
