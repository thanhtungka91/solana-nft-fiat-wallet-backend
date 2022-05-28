package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thanhtungka91/solana-nft-fiat-wallet-backend/models"
	"net/http"
)

func main() {
	router := gin.Default()

	db := models.ConnectDatabase()
	fmt.Println(db)
	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/balance/:wallet_address", func(c *gin.Context) {
		walletAddress := c.Param("wallet_address")

		result, _ := models.GetBalance(db, walletAddress)

		c.JSON(http.StatusOK, result)

	})

	router.POST("/user/:wallet_address", func(c *gin.Context) {
		b := c.FullPath() == "/user/:name/*action" // true
		c.String(http.StatusOK, "%t", b)
	})

	router.Run(":9000")
}
