package web

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/rs/zerolog/log"

	"blockchain/core"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

var router *gin.Engine

func setup(verbose bool) {

	gin.SetMode(gin.ReleaseMode)
	if verbose {
		gin.SetMode(gin.DebugMode)
	}

	router = gin.Default()
}

func Start(port int, verbose bool) {
	setup(verbose)

	router.Group("/")
	{
		router.GET("/", redirect)
		router.GET("/blocks", getAllBlocks)
		router.GET("/blocks/:hash", getBlockByHash)
		router.GET("/tx/:from/:to/:amount", postTx) // TODO change to POST
	}

	log.Info().Msgf("starting http server, listening on port %d...", port)
	router.Run(fmt.Sprintf(":%d", port))
	log.Info().Msg("http server started")
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

	amount, err := strconv.ParseInt(ctx.Param("amount"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newTx := core.Tx{
		From:   ctx.Param("from"),
		To:     ctx.Param("to"),
		Amount: uint64(amount),
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
